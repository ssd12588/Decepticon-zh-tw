# 中文文件（docs-zh）

這裡是 **Decepticon 繁體中文漢化版**的中文入門文件。原廠英文文件仍在 [`docs/`](../docs/)，兩者並存。

## 從哪裡開始

| 文件 | 內容 | 什麼時候看 |
|------|------|-----------|
| [QUICKSTART.md](QUICKSTART.md) | 從零到第一次執行，含「設定選項對照表」 | 第一次安裝 |
| [TROUBLESHOOTING.md](TROUBLESHOOTING.md) | 把中文錯誤訊息對照回原因與解法 | 卡住的時候 |
| [glossary.md](glossary.md) | 漢化術語對照表（一詞一譯） | 想改翻譯或回報用詞 |
| [PROJECT-STATUS.md](PROJECT-STATUS.md) | 交接文件：進度、驗證狀態、未完成事項 | 隔一段時間回來、或換 AI 對話接手 |

## 漢化範圍（一眼看懂）

- **已漢化**：CLI 的 `--help` 與錯誤訊息（主畫面、`status`、`logs`、`doctor`、`update` 等）、`doctor` 診斷輸出與復原指令、README、`AGENTS.md`、本目錄四份文件。
- **進行中**：`onboard` 精靈、TUI 面板、Python 無介面 CLI、`CONTRIBUTING.md`、`docs/` 全量文件。
- **刻意不改**：代理提示詞（`packages/decepticon/decepticon/agents/prompts/**`）、技能知識包（`skills/**`、`SKILL.md`）、程式碼識別字、協定欄位、JSON 鍵、狀態值（`running`／`completed`／`failed`）、CLI 參數、環境變數、payload、HTTP 請求與回應、`LICENSE`。

## 讓 AI 用繁體中文回答你

Decepticon 上游已內建語言機制，不必改提示詞：

```bash
# 方式一：寫進 ~/.decepticon/.env
DECEPTICON_LANGUAGE=tw

# 方式二：只對這次執行生效
DECEPTICON_LANGUAGE=tw decepticon
```

`tw` 會被解析成 `zh-tw`（Traditional Chinese）。代理的對人說明會變成繁體中文，
而工具名稱、指令、檔案路徑與技術縮寫仍保留英文 —— 這樣它才不會把 `running` 之類的
狀態值也翻掉。

## 最常見的三個坑

1. **Windows 中文亂碼**：本漢化版已在啟動器啟動時把主控台切到 UTF-8（CP65001）。若你看到方框或問號，請用 Windows Terminal 或 PowerShell 7，並執行 `chcp 65001`。
2. **Docker 沒有在跑**：`decepticon doctor` 會顯示 `docker_daemon: 失敗`，並在下方給你可直接複製的復原指令。
3. **模型金鑰沒設定**：`auth: 失敗` 代表還沒跑過 `decepticon onboard`，或 `.env` 裡的憑證是預留字串（`your-...-key-here`）。

## 回報翻譯問題

請在 [本倉庫的 issue](https://github.com/ssd12588/Decepticon-zh-tw/issues) 回報，附上：
指令、完整輸出、你的作業系統與終端機。**不要**打到上游 `BitterSecurity/Decepticon`。