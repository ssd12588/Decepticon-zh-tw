package cmd

import (
	"encoding/json"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/compose"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "顯示 Decepticon 服務狀態",
	Long:    "顯示目前 Decepticon Compose 堆疊中的服務。加上 --json 會輸出適合腳本處理的 JSON 陣列，沒有服務在執行時輸出 []。",
	Example: "  decepticon status\n  decepticon status --json",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		asJSON, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}
		c := compose.New()
		if !asJSON {
			return c.Ps()
		}
		rows, err := c.PsJSON()
		if err != nil {
			return err
		}
		return json.NewEncoder(cmd.OutOrStdout()).Encode(rows)
	},
}

func init() {
	statusCmd.Flags().Bool("json", false, "以 JSON 陣列輸出服務列表")
	rootCmd.AddCommand(statusCmd)
}
