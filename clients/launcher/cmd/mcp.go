package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/compose"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/config"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "讓程式開發代理連接執行中的 Decepticon",
	Long: `透過 Model Context Protocol（MCP）把 Decepticon 的 MCP 工具提供給
程式開發代理使用。

請先啟動 Decepticon（'decepticon start'），再讓你的程式開發代理
指向這個 MCP 伺服器，即可操作執行中的 Decepticon：查看狀態與健康
檢查、讀取日誌，並依授權測試任務流程作業。`,
	Example: `  decepticon mcp serve             以 stdio 提供 Decepticon MCP 工具

  # 註冊到 Claude Code（請先執行 'decepticon start'）：
  claude mcp add decepticon -- decepticon mcp serve`,
}

var mcpServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "以 stdio 提供 Decepticon MCP 工具",
	Long:  "以 stdio 對 Claude Code、Codex 或其他 MCP 用戶端提供 Decepticon MCP 工具。請先啟動 Decepticon。",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		home := config.DecepticonHome()
		for _, name := range []string{"docker-compose.yml", ".env"} {
			if _, err := os.Stat(filepath.Join(home, name)); err != nil {
				return fmt.Errorf("Decepticon 尚未安裝於 %s；請先執行 decepticon onboard：%w", home, err)
			}
		}
		return compose.New().RunMCPStdio()
	},
}

func init() {
	mcpCmd.AddCommand(mcpServeCmd)
	rootCmd.AddCommand(mcpCmd)
}
