package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/compose"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/config"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/engagement"
	"github.com/spf13/cobra"
)

var documentedLogServices = []string{"langgraph", "litellm", "postgres", "neo4j", "sandbox"}

var logsCmd = &cobra.Command{
	Use:   "logs [service]",
	Short: "追蹤服務日誌，或重播授權測試任務的對話紀錄",
	Long: `兩種模式：

服務模式（預設）：追蹤目前 Compose 堆疊中某個服務的日誌。
預設服務為 langgraph。已記載的服務：langgraph、litellm、postgres、
neo4j、sandbox。其他堆疊服務（例如 cli 或 c2-sliver）只要存在於
Compose 設定中也可接受。

任務模式（--engagement）：從 LangGraph 對話串狀態重播某個授權測試
任務（可能為背景執行）的對話紀錄，也就是 MCP decepticon_transcript
工具讀取的同一份內容。唯讀；該任務不需要正在執行。`,
	Example: `  decepticon logs
  decepticon logs litellm
  decepticon logs cli
  decepticon logs --engagement my-engagement
  decepticon logs --engagement my-engagement --limit 20
  decepticon logs --engagement 01a16679-e6af-4d09-a238-57ae6af4e79c`,
	Args: cobra.MaximumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		names, err := compose.New().ServiceNames()
		if err != nil {
			names = documentedLogServices
		}
		return names, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		engagementName, err := cmd.Flags().GetString("engagement")
		if err != nil {
			return err
		}
		if engagementName != "" {
			if len(args) > 0 {
				return fmt.Errorf("服務參數與 --engagement 不能同時使用")
			}
			return runEngagementTranscript(cmd, engagementName)
		}

		service := "langgraph"
		if len(args) > 0 {
			service = args[0]
		}
		if !isDocumentedLogService(service) {
			names, err := compose.New().ServiceNames()
			if err != nil {
				return fmt.Errorf("驗證服務 %q 失敗：%w", service, err)
			}
			found := false
			for _, name := range names {
				if name == service {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("未知的服務 %q；已記載的服務：%s", service, strings.Join(documentedLogServices, "、"))
			}
		}
		return compose.New().Logs(service)
	},
}

// runEngagementTranscript resolves the engagement (or explicit thread id) to
// its LangGraph thread and prints the trailing transcript window. The
// endpoint resolution matches the readiness probe: process env overrides
// the .env file, matching Compose interpolation.
func runEngagementTranscript(cmd *cobra.Command, target string) error {
	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		return err
	}
	env, envErr := config.LoadEnv(config.EnvPath())
	if envErr != nil {
		env = map[string]string{}
	}
	baseURL := langgraphBaseURL(env)
	client := engagement.DefaultLogsClient()
	ctx := cmd.Context()

	threadID, err := engagement.ResolveThreadID(ctx, client, baseURL, target)
	if err != nil {
		return err
	}
	state, err := engagement.FetchThreadState(ctx, client, baseURL, threadID)
	if err != nil {
		return err
	}
	cmd.Printf("授權測試任務 %s（對話串 %s）—— %d 則訊息，計畫修訂版 %d\n\n",
		state.EngagementName, state.ThreadID, len(state.Messages), state.PlanRevision)
	for _, message := range engagement.ReplayLimit(state.Messages, limit) {
		cmd.Printf("[%s] %s\n", message.Role, strings.ReplaceAll(message.Content, "\n", "\n  "))
	}
	return nil
}

// langgraphBaseURL resolves the stack's host LangGraph endpoint the same way
// the readiness probe does.
func langgraphBaseURL(env map[string]string) string {
	port := config.Get(env, "LANGGRAPH_PORT", "2024")
	if processPort, set := os.LookupEnv("LANGGRAPH_PORT"); set {
		if processPort == "" {
			port = "2024"
		} else {
			port = processPort
		}
	}
	return "http://127.0.0.1:" + port
}

func isDocumentedLogService(name string) bool {
	for _, known := range documentedLogServices {
		if name == known {
			return true
		}
	}
	return false
}

func init() {
	logsCmd.Flags().String("engagement", "", "重播授權測試任務的對話紀錄（名稱或對話串 ID），而不追蹤服務日誌")
	logsCmd.Flags().Int("limit", 50, "輸出最後 N 則訊息（任務模式）")
	_ = logsCmd.Flags().SetAnnotation("engagement", cobra.BashCompFilenameExt, []string{""})
	rootCmd.AddCommand(logsCmd)
}
