package cmd

import (
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/compose"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/ui"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:     "kg-health",
	Aliases: []string{"graph-health"},
	Short:   "執行 LangGraph 健康診斷",
	Long: `在執行中的 langgraph 服務內執行 LangGraph 健康診斷
（python -m decepticon.tools.research.health）並印出報告。
健康時結束碼為 0，否則為 1。

服務堆疊必須正在執行：請先以 'decepticon start' 啟動；
檢查失敗時可用 'decepticon logs langgraph' 查看服務日誌。`,
	Example: `  decepticon kg-health                  執行診斷（別名：decepticon graph-health）
  decepticon kg-health 2>/dev/null      抑制容器日誌雜訊`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.Info("正在執行 LangGraph 健康診斷…")
		return compose.New().Exec("langgraph", "python", "-m", "decepticon.tools.research.health")
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
