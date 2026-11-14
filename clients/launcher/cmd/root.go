package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/ui"
	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:   "decepticon",
	Args:  cobra.ArbitraryArgs,
	Short: "Decepticon —— 紅隊自主駭客代理",
	Long: ui.RenderBanner() + "\n" + ui.Dim.Render("紅隊自主駭客代理") + `

啟動授權紅隊工作階段：
  decepticon                        啟動服務並開啟對話
                                    （首次啟動會執行 onboard 初始化精靈）
  decepticon --continue             接續最近一次工作階段
  decepticon --resume               開啟互動式工作階段選擇器
  decepticon --resume=<thread-id>   接續指定的對話串

管理服務堆疊與安裝：
  decepticon status                 顯示執行中的服務（--json 供腳本使用）
  decepticon logs langgraph         追蹤代理服務日誌
  decepticon doctor                 診斷安裝狀態（--json 供 CI 使用）
  decepticon update                 檢查並套用更新
  decepticon stop                   停止所有服務（離開 CLI 後服務仍會繼續執行，
                                    讓下次啟動更快）

執行 'decepticon <command> --help' 可查看單一指令的詳細說明。`,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	decorateLauncherHelp()
	defaultHelp := rootCmd.HelpFunc()
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		out := cmd.OutOrStdout()
		var text bytes.Buffer
		cmd.SetOut(&text)
		defer cmd.SetOut(out)
		defaultHelp(cmd, args)
		_, _ = (ui.NeutralWriter{Writer: out}).Write(text.Bytes())
	})
	command, err := rootCmd.ExecuteC()
	if err != nil {
		if errors.Is(err, errDoctorFailed) {
			os.Exit(1)
		}
		if command == mcpServeCmd {
			fmt.Fprintln(os.Stderr, err)
		} else {
			ui.Error(err.Error())
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Bool("no-alt-screen", false, "使用終端機原生捲動，而不使用全螢幕對話介面")
	rootCmd.PersistentFlags().BoolVar(&ui.DisableMouse, "no-mouse", false, "停用終端機滑鼠擷取")
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(fmt.Sprintf("Decepticon %s\n", version))
}
