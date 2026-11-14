package cmd

import (
	"fmt"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/agentskill"
	"github.com/spf13/cobra"
)

var skillOptions agentskill.Options

var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "為程式開發代理安裝 Decepticon 操作指引",
	Long: `為 Claude Code 或 Codex 等程式開發代理安裝 Decepticon Agent Skill，
讓它們知道如何透過已安裝的 CLI 操作執行中的 Decepticon
（status、health、logs、互動式工作階段）。`,
}

var skillInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "為 Claude Code 或 Codex 安裝 Decepticon Agent Skill",
	Long: `為 Claude Code、Codex 或兩者同時安裝 Decepticon Agent Skill。

這個技能會教程式開發代理如何連上執行中的 Decepticon：查看狀態與
健康檢查、讀取日誌、操作互動式授權測試任務。MCP 屬於選用整合。`,
	Example: `  decepticon skill install                  同時安裝給 Claude Code 與 Codex
  decepticon skill install --client claude  只安裝給 Claude Code
  decepticon skill install --client codex   只安裝給 Codex
  decepticon skill install --force          備份後取代已修改的技能`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		skillOptions.Version = version
		paths, err := agentskill.Install(cmd.Context(), skillOptions)
		for _, path := range paths {
			fmt.Fprintln(cmd.OutOrStdout(), "已安裝 Decepticon Skill 於", path)
		}
		return err
	},
}

func init() {
	skillInstallCmd.Flags().StringVar(&skillOptions.Client, "client", "both", "程式開發代理：codex、claude 或 both")
	skillInstallCmd.Flags().StringVar(&skillOptions.Source, "from", "", "從本機 Decepticon Skill 目錄安裝")
	skillInstallCmd.Flags().BoolVar(&skillOptions.Force, "force", false, "備份後取代已修改的 Decepticon Skill")
	skillCmd.AddCommand(skillInstallCmd)
	rootCmd.AddCommand(skillCmd)
}
