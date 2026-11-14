# 專案狀態（交接用）

> 給「未來的自己」或「新的 AI 對話」快速掌握現況。
> 最後更新：完成 README 主頁漢化、CLI 第一批與 doctor 之後。

---

## 一句話現況

`ssd12588/Decepticon-zh-tw` 是 [BitterSecurity/Decepticon](https://github.com/BitterSecurity/Decepticon)
的**繁體中文漢化版**。啟動器 CLI 與 README 已完成並通過建置／測試；TUI、Python CLI 與
`docs/` 全量文件仍在進行中。

- 預設分支：`main`（已含目前所有漢化，GitHub 首頁顯示中文 README）
- 同步分支：`zh-TW`（與 `main` 同步，日常開發推這裡）
- 版本基準：上游 `main` @ `4cecf55`（PR #1196 合併後）
- 授權：Apache-2.0（`LICENSE` 原文保留、未修改）

---

## 已完成

| 範圍 | 位置 | 狀態 |
|---|---|---|
| 啟動器 CLI 主畫面與第一批指令 | `clients/launcher/cmd/{root,status,logs,health,stop,cancel,update,skill,mcp,help_ux}.go` | ✅ |
| `doctor` 診斷輸出與復原指令 | `clients/launcher/cmd/doctor.go` | ✅ |
| Windows UTF-8 主控台（修 CP950 亂碼） | `clients/launcher/internal/platform/console_*.go`、`main.go` | ✅ |
| README 主頁（含「關於本漢化版」） | `README.md` | ✅ |
| 代理開發規範 | `AGENTS.md` | ✅ |
| 術語表 | `docs-zh/glossary.md` | ✅ |
| 中文文件索引 | `docs-zh/README.md` | ✅ |

**刻意保持英文**（改了會影響 AI 判斷或直接弄壞程式）：

- `packages/decepticon/decepticon/agents/prompts/**`（33 檔）
- `packages/decepticon/decepticon/skills/**`、`skills/**`（349 個 `SKILL.md`）
- 程式碼識別字、JSON 鍵、協定欄位、狀態值、CLI 參數、環境變數、payload

---

## 驗證狀態

| 檢查 | 指令 | 結果 |
|---|---|---|
| Go 編譯 | `cd clients/launcher && go build ./...` | ✅ |
| Go 測試（CLI 套件） | `go test ./cmd/...` | ✅ `ok` |
| Go 測試（平台／UI） | `go test ./internal/platform/... ./internal/ui/...` | ✅ `ok` |
| 實際輸出 | `decepticon --help` | ✅ 繁體中文正常 |
| 實際輸出 | `decepticon doctor` | ✅ 繁體中文（含復原方式） |
| 全量 Go 測試 | `go test ./...` | ⚠️ `internal/infracontrol` 1 項失敗，見下方已知限制 |
| TUI／Python CLI／`docs/` | — | ⬜ 尚未開始 |

---

## 已知限制

1. **`internal/infracontrol` symlink 測試**：`TestEngagementWorkspaceRejectsEscapeAndMissingDirectory`
   在 Windows 需 Developer Mode 或系統管理員權限才能建立 symlink，本機回報
   `A required privilege is not held by the client`。**與漢化無關**，上游行為一致。
2. **本機 clone 是 shallow**：為了快速勘察用 `--depth 1` 取得。要在本機做
   `git merge upstream/main` 之前，請先 `git fetch --unshallow upstream`。
3. **`docs/` 179 篇尚未漢化**：目前只需漢化的部分已完成，全量文件翻譯是下一階段。

---

## 尚未完成（依優先序）

| 優先 | 項目 | 說明 |
|---|---|---|
| ★★★ | `clients/launcher/cmd/onboard.go` | 初始化精靈，約 1,300 行、字串最多的一支 |
| ★★★ | `clients/launcher/cmd/start.go` | 啟動流程、就緒提示、錯誤訊息 |
| ★★★ | `internal/health`、`internal/compose`、`internal/engagement` | 錯誤訊息（`doctor` 輸出裡仍可見英文） |
| ★★ | Ink TUI（`clients/cli/src/**`） | 24 個 component、26 個 slash command 模組、`shortcuts.ts` |
| ★★ | Python 無介面 CLI（`packages/decepticon/decepticon/cli/*.py`） | argparse 說明與錯誤訊息 |
| ★★ | `CONTRIBUTING.md`、`docs/*.md` | 179 篇，建議分批（先 getting-started、setup-guide、cli-reference） |
| ★ | `start-decepticon.ps1` | Windows 一鍵啟動（CRLF 換行） |
| ★ | 中文 `TROUBLESHOOTING.md` | 依實際錯誤訊息持續補齊 |

❌ **不要做**：漢化 `prompts/**`、`skills/**`、`SKILL.md`、狀態值、指令參數。

---

## 常用指令

```powershell
# 建置與測試（Go 啟動器）
cd clients/launcher
go build ./...
go test ./cmd/...

# 產生可執行檔並看中文說明
go build -ldflags "-X github.com/BitterSecurity/Decepticon/clients/launcher/cmd.version=dev" -o bin/decepticon.exe .
.\bin\decepticon.exe --help
.\bin\decepticon.exe doctor

# 同步上游
git fetch upstream
git merge upstream/main

# 推送漢化
git push origin zh-TW
git checkout main; git merge --ff-only zh-TW; git push origin main; git checkout zh-TW
```

---

## 分支策略

- `main`：對外顯示（GitHub 首頁讀這個），**只接受從 `zh-TW` fast-forward 的合併**
- `zh-TW`：日常開發分支；每批完成後 `git push origin zh-TW`

---

## 給新對話的開場提示

> 這個倉庫是 `ssd12588/Decepticon-zh-tw`，上游 `BitterSecurity/Decepticon` 的繁體中文漢化版。
> 請先讀 `docs-zh/PROJECT-STATUS.md` 與 `docs-zh/glossary.md` 了解現況與術語，再依
> 「尚未完成」清單繼續。漢化時：
> 1. 只改「給人看」的字面量；不要動 `prompts/**`、`skills/**`、狀態值、指令參數與識別字。
> 2. 每批完成都要跑 `go build ./...` 與對應測試，並同步更新被打斷的英文測試斷言。
> 3. 提交訊息用 UTF-8 檔案 + `git commit -F`；每批推 `zh-TW`，必要時再 fast-forward 到 `main`。
> 4. 中文一律繁體、全角標點、破折號 `——`、中英之間不加空格。