package cmd

import (
	"fmt"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/config"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/engagement"
	"github.com/spf13/cobra"
)

var cancelCmd = &cobra.Command{
	Use:   "cancel <engagement-name|thread-id>",
	Short: "取消授權測試任務正在執行的回合",
	Long: `取消某個授權測試任務（可能是背景執行）正在跑的回合。服務不會被關閉 ——
這裡只停止該任務的代理回合，不會停掉整個服務堆疊；要停整個堆疊請用
'decepticon stop'，之後想重新啟動的回合再執行一次 'decepticon cancel'。

除了取消動作本身之外皆為唯讀。若目前沒有執行中的回合，只會回報，不算錯誤。`,
	Example: `  decepticon cancel my-engagement
  decepticon cancel 01a16679-e6af-4d09-a238-57ae6af4e79c`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		env, envErr := config.LoadEnv(config.EnvPath())
		if envErr != nil {
			env = map[string]string{}
		}
		baseURL := langgraphBaseURL(env)
		client := engagement.DefaultLogsClient()
		ctx := cmd.Context()

		threadID, err := engagement.ResolveThreadID(ctx, client, baseURL, args[0])
		if err != nil {
			return err
		}
		run, err := engagement.FindActiveRun(ctx, client, baseURL, threadID)
		if err != nil {
			return err
		}
		if run == nil {
			cmd.Printf("授權測試任務 %s 目前沒有執行中的回合 —— 無需取消。\n", args[0])
			return nil
		}
		if err := engagement.CancelRun(ctx, client, baseURL, threadID, run.RunID); err != nil {
			return fmt.Errorf("取消回合 %s：%w", run.RunID, err)
		}
		cmd.Printf("正在取消授權測試任務 %s（對話串 %s，回合 %s，原狀態 %s）。\n",
			args[0], threadID, run.RunID, run.Status)
		cmd.Println("服務仍在執行 —— 可用 'decepticon attach' 檢視，用 'decepticon stop' 關閉。")
		return nil
	},
}

func init() { rootCmd.AddCommand(cancelCmd) }
