package cmd

import (
	"fmt"
	"strings"

	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/compose"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/config"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/ui"
	"github.com/BitterSecurity/Decepticon/clients/launcher/internal/updater"
	"github.com/spf13/cobra"
)

var (
	forceUpdate   bool
	updateChannel string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "檢查並套用更新",
	Long: `到 GitHub 檢查是否有較新的 Decepticon 發行版並套用。

會更新啟動器執行檔、重新同步設定檔（docker-compose.yml、
litellm 設定等）並拉取更新後的 Docker 映像。加上 --force 時，
即使版本沒有變更也會重新整理設定檔與 Docker 映像。

更新通道預設讀取 .env 的 DECEPTICON_CHANNEL（找不到時用
latest）；--channel 只覆寫本次執行。啟動時的更新檢查可用
~/.decepticon/.env 的 AUTO_UPDATE 控制，或加上 --no-update 停用。`,
	Example: `  decepticon update                檢查並套用更新
  decepticon update --force        重新同步設定檔與 Docker 映像
  decepticon update --channel stable
                                   本次執行改追 stable（已浸泡）通道`,
	RunE: runUpdate,
}

func init() {
	updateCmd.Flags().BoolVarP(&forceUpdate, "force", "f", false, "即使版本未變更，也重新整理設定檔與 Docker 映像")
	updateCmd.Flags().StringVar(&updateChannel, "channel", "", "本次執行的更新通道：stable（已浸泡的正式版）或 latest（最新的正式版）。預設讀取 .env 的 DECEPTICON_CHANNEL，否則為 latest")
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cmd *cobra.Command, args []string) error {
	// Load env first so the channel (and branch override) are known
	// before the release fetch.
	env := make(map[string]string)
	if config.EnvExists() {
		env, _ = config.LoadEnv(config.EnvPath())
	}
	// --channel overrides .env for this invocation only.
	rawChannel := updateChannel
	if rawChannel == "" {
		rawChannel = config.Get(env, "DECEPTICON_CHANNEL", "")
	}
	ch := updater.ResolveChannel(rawChannel)

	ui.Info(fmt.Sprintf("正在檢查更新（%s 通道）…", ch))

	release, err := updater.FetchRelease(ch)
	if err != nil {
		return fmt.Errorf("檢查更新失敗：%w", err)
	}

	hasUpdate := updater.CompareVersions(version, release.TagName)
	if !hasUpdate && !forceUpdate {
		ui.Success(fmt.Sprintf("已是最新版本（%s）", version))
		return nil
	}

	if hasUpdate {
		ui.Info(fmt.Sprintf("有新版本可用：%s → %s", version, release.TagName))
	} else {
		ui.Info("正在重新整理設定檔與 Docker 映像…")
	}

	ref := release.TagName
	if branch := strings.TrimSpace(env["DECEPTICON_BRANCH"]); branch != "" {
		ref = branch
	}

	if hasUpdate {
		// Full upgrade flow — shared with the launch-time interactive
		// prompt so behavior stays consistent between the two paths.
		if err := updater.ApplyUpdate(release, ref); err != nil {
			ui.Warning(err.Error())
		}
	} else {
		// --force: re-sync config + re-pull images without bumping the
		// binary (already on release.TagName).
		ui.Info("正在同步設定檔…")
		// Mirror ApplyUpdate's release/branch split: pass the Release
		// through when ref tracks the tag so the manifest verifies, nil
		// when DECEPTICON_BRANCH is overriding the ref (branch mode).
		syncRelease := release
		if ref != release.TagName && ref != strings.TrimPrefix(release.TagName, "v") {
			syncRelease = nil
		}
		if err := updater.SyncConfigFiles(ref, syncRelease); err != nil {
			ui.Warning("設定檔同步失敗：" + err.Error())
		}
		c := compose.New()
		targetVersion := strings.TrimPrefix(release.TagName, "v")
		ui.Info("正在拉取 Docker 映像（" + targetVersion + "）…")
		if err := c.Pull(targetVersion); err != nil {
			ui.Warning("映像拉取失敗：" + err.Error())
		}
	}

	ui.Success("更新完成")
	ui.DimText("請重新啟動 Decepticon 以使用新版本")
	return nil
}
