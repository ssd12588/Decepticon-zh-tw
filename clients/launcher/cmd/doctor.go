package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/config"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/health"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/infracontrol"
	"github.com/spf13/cobra"
)

var errDoctorFailed = errors.New("必要檢查未通過；請參考 doctor 報告中的復原指令")

type doctorCheck struct {
	Name     string `json:"name"`
	Required bool   `json:"required"`
	OK       bool   `json:"ok"`
	Detail   string `json:"detail"`
	Recovery string `json:"recovery,omitempty"`
}

type doctorReport struct {
	OK     bool          `json:"ok"`
	Checks []doctorCheck `json:"checks"`
}

func init() {
	command := &cobra.Command{
		Use:   "doctor",
		Short: "唯讀的安裝與服務診斷",
		Long: `對 Decepticon 安裝執行唯讀診斷。

依序檢查：Docker 是否存在與常駐程式是否回應、Docker Compose V2、
設定檔（.env 與 docker-compose.yml）與驗證方式、LangGraph 就緒狀態、
opscontrol Unix socket 與常駐程式健康狀態，以及選用的藍隊監控端點。

加上 --deep 時，會在 langgraph 容器內額外執行唯讀的執行期診斷
（python -m decepticon.diagnostics）：KG 儲存區可達性、Skillogy 健康
狀態、HITL／模型備援插槽解析，以及 MCP 驗證模式。

本指令不會修改任何東西。必要檢查失敗時結束碼非 0；選用檢查只會
回報，不影響結束碼。`,
		Example: `  decepticon doctor           執行所有診斷並印出報告
  decepticon doctor --json    輸出適合 CI 使用的機器可讀報告
  decepticon doctor --deep    同時執行容器內的執行期診斷`,
		Args: cobra.NoArgs,
	}
	command.Flags().Bool("json", false, "輸出機器可讀的報告")
	command.Flags().Bool("deep", false, "在 langgraph 容器內額外執行唯讀的執行期診斷（需要服務堆疊）")
	command.Flags().String("engagement", "", "把深度 KG 修訂版與節點統計限定在此授權測試任務")
	command.RunE = func(cmd *cobra.Command, _ []string) error {
		engagement, _ := cmd.Flags().GetString("engagement")
		if engagement != "" {
			valid, _ := regexp.MatchString(`^[a-z0-9][a-z0-9-]{1,62}[a-z0-9]$`, engagement)
			if !valid {
				return fmt.Errorf("授權測試任務名稱格式不正確")
			}
			deep, _ := cmd.Flags().GetBool("deep")
			if !deep {
				return fmt.Errorf("--engagement 需要搭配 --deep")
			}
		}
		report := diagnose(cmd.Context())
		if deep, err := cmd.Flags().GetBool("deep"); err != nil {
			return err
		} else if deep {
			addDeepRuntimeChecksForEngagement(cmd.Context(), &report, engagement)
		}
		jsonOutput, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}
		if jsonOutput {
			if err := json.NewEncoder(cmd.OutOrStdout()).Encode(report); err != nil {
				return fmt.Errorf("寫入 doctor 報告失敗：%w", err)
			}
		} else {
			for _, check := range report.Checks {
				status := "正常"
				if !check.OK {
					status = "失敗"
					if !check.Required {
						status = "選用"
					}
				}
				cmd.Printf("%s: %s —— %s\n", check.Name, status, check.Detail)
				if check.Recovery != "" {
					cmd.Printf("  復原方式：%s\n", check.Recovery)
				}
			}
		}
		if !report.OK {
			return errDoctorFailed
		}
		return nil
	}
	rootCmd.AddCommand(command)
}

func diagnose(ctx context.Context) doctorReport {
	report := doctorReport{OK: true}
	add := func(name string, required, ok bool, detail, recovery string) {
		if ok {
			recovery = ""
		}
		report.Checks = append(report.Checks, doctorCheck{name, required, ok, detail, recovery})
		if required && !ok {
			report.OK = false
		}
	}
	_, err := exec.LookPath("docker")
	add("docker", true, err == nil, "Docker 執行檔是否存在", "開啟 https://docs.docker.com/get-docker/ 安裝 Docker，然後執行 `docker version`")
	add("docker_daemon", true, doctorExec(ctx, "docker", "info", "--format", "{{.ServerVersion}}") == nil, "Docker 常駐程式回應（3 秒逾時）", "開啟 Docker Desktop，或執行 `sudo systemctl start docker`，再重試 `decepticon doctor`")
	add("compose", true, doctorExec(ctx, "docker", "compose", "version") == nil, "Docker Compose V2 回應（3 秒逾時）", "安裝 Docker Compose 外掛，並用 `docker compose version` 確認")
	env, configErr := config.LoadEnv(config.EnvPath())
	_, composeErr := os.Stat(filepath.Join(config.DecepticonHome(), "docker-compose.yml"))
	add("config", true, configErr == nil && composeErr == nil, "設定檔可讀且存在 Compose 檔案", "全新安裝請執行 `decepticon onboard`；既有安裝請先用 `decepticon status` 檢視再變更設定")
	authOK := configErr == nil && config.ValidateAuth(env) == nil
	add("auth", true, authOK, "已設定的驗證方式（僅檢查是否設定，不驗證供應商登入是否有效）", "執行 `decepticon onboard` 設定驗證方式")
	client := &http.Client{Timeout: 3 * time.Second, CheckRedirect: func(_ *http.Request, _ []*http.Request) error { return http.ErrUseLastResponse }}
	graphCtx, cancelGraph := context.WithTimeout(ctx, 3*time.Second)
	graphErr := health.ProbeLangGraph(graphCtx, env)
	cancelGraph()
	graphDetail := "主機端 LangGraph HTTP 與 decepticon assistant 已就緒"
	if graphErr != nil {
		graphDetail = graphErr.Error()
	}
	add("langgraph", true, graphErr == nil, graphDetail, "檢查 LANGGRAPH_PORT 與 Docker 對外連接埠；用 `decepticon logs langgraph` 對照容器健康狀態；再重試 `decepticon start --no-update`")
	socket := infracontrol.HostSocketPath()
	info, socketErr := os.Stat(socket)
	add("opscontrol_socket", false, socketErr == nil && info.Mode()&os.ModeSocket != 0, "Opscontrol Unix socket 是否存在", "執行 `decepticon infracontrol status`；用 `decepticon start --no-update` 啟動")
	transport := &http.Transport{DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
		return (&net.Dialer{Timeout: 3 * time.Second}).DialContext(ctx, "unix", socket)
	}}
	defer transport.CloseIdleConnections()
	opsClient := &http.Client{Transport: transport, Timeout: 3 * time.Second}
	add("opscontrol", false, doctorHTTP(ctx, opsClient, "http://opscontrol/v1/health"), "Opscontrol 常駐程式健康回應", "執行 `decepticon infracontrol status`；並查看 `decepticon logs`")
	for _, component := range []struct{ name, key, fallback string }{{"blue_sensor", "BLUE_SENSOR_URL", "http://127.0.0.1:18081"}, {"blue_monitor", "BLUE_MONITOR_URL", "http://127.0.0.1:18085"}} {
		endpoint := config.Get(env, component.key, component.fallback)
		detail := "未設定主機端端點；監控狀態未知"
		ok := false
		if endpoint != "" {
			ok = doctorHTTP(ctx, client, endpoint+"/metrics")
			detail = "已設定主機端端點的健康回應"
		}
		add(component.name, false, ok, detail, "執行 `decepticon logs`；設定一個主機可連的 "+component.key+" 後重試 `decepticon doctor`")
	}
	return report
}

func doctorExec(ctx context.Context, name string, args ...string) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, name, args...)
	command.WaitDelay = time.Second
	return command.Run()
}

func doctorHTTP(ctx context.Context, client *http.Client, endpoint string) bool {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return false
	}
	response, err := client.Do(request)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode >= 200 && response.StatusCode < 300
}
