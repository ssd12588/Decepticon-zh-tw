package cmd

import (
	"github.com/BitterSecurity/Decepticon/clients/launcher/cmd/infracontrol"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/compose"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/ui"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止所有 Decepticon 服務",
	Long: `停止所有 Decepticon 服務，包含 opscontrol 常駐程式。

會停止並移除所有工作負載設定檔中的 Compose 容器，然後關閉
opscontrol 常駐程式。對話與工作區資料都會保留，只停止正在
執行的服務。`,
	Example: `  decepticon stop    停止所有 Decepticon 服務與常駐程式`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := compose.New()
		ui.Info("正在停止 Decepticon 服務…")
		c.RemoveOrphanedCLI()
		// Clear legacy root-level scratch/session buffers before tearing the
		// stack down. Current runs keep these under engagement workspaces.
		c.CleanScratch()
		if err := c.Down(); err != nil {
			return err
		}
		// Tear down the infracontrol daemon after compose so any
		// workload it owned has already been removed by `compose down`.
		// Non-fatal: a never-started daemon is a no-op here.
		if err := infracontrol.Stop(); err != nil {
			ui.Warning("infracontrol 停止失敗：" + err.Error())
		}
		ui.Success("所有服務已停止")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
