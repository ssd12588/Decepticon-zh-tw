# 10 分鐘上手（QUICKSTART）

從零到第一次執行，全部用繁體中文介面。遇到錯誤請對照 [TROUBLESHOOTING.md](TROUBLESHOOTING.md)。

---

## 0. 你需要什麼

| 需求 | 說明 |
|---|---|
| Docker ＋ Docker Compose v2 | 其餘全部都在容器裡跑，不必本機裝 Python／Go／Node |
| 一個 LLM 憑證 | API 金鑰（Anthropic／OpenAI／Gemini／DeepSeek／OpenRouter…）、訂閱制 OAuth，或本機 Ollama |
| 授權 | 只對**你自己擁有、或已取得明確書面授權**的目標執行 |

---

## 1. 安裝

**macOS / Linux / WSL2**

```bash
curl -fsSL https://decepticon.red/install | bash
```

**Windows（原生 PowerShell）**

```powershell
irm https://decepticon.red/install.ps1 | iex
```

安裝完成後，`decepticon` 指令就會在 PATH 上。先做一次唯讀檢查：

```bash
decepticon doctor      # 看 Docker／Compose／設定／驗證是否就緒
```

`doctor` 的輸出會長這樣（中文狀態：`正常`／`失敗`／`選用`）：

```text
docker: 正常 —— Docker 執行檔是否存在
compose: 正常 —— Docker Compose V2 回應（3 秒逾時）
config: 失敗 —— 設定檔可讀且存在 Compose 檔案
  復原方式：全新安裝請執行 `decepticon onboard`；既有安裝請先用 `decepticon status` 檢視再變更設定
```

---

## 2. 設定（onboard 精靈）

```bash
decepticon onboard
```

精靈會依序問你五件事：

1. **驗證方式** —— API 金鑰、訂閱制 OAuth（Claude／ChatGPT／Gemini／Copilot／SuperGrok／Perplexity），或本機 Ollama
2. **供應商** —— 上面選定的那家，或自訂的 OpenAI 相容端點
3. **憑證** —— API 金鑰、OAuth token 或端點 URL
4. **模型設定檔** —— `eco`（平衡，預設）／`max`（效能）／`test`（開發 CI）
5. **LangSmith** —— 選用的追蹤，不需要就跳過

設定會寫進 `~/.decepticon/.env`。想重設：`decepticon onboard --reset`。

### 設定選項對照表

| 變數／選項 | 用途 | 常見值 | 備註 |
|---|---|---|---|
| `DECEPTICON_LANGUAGE` | 代理「對人說話」的語言 | `tw`、`en`、`ja` | **本漢化版建議設 `tw`**；工具名與指令仍為英文 |
| `DECEPTICON_CHANNEL` | 更新通道 | `latest`、`stable` | `stable` 是浸泡後的正式版 |
| `AUTO_UPDATE` | 啟動時自動檢查更新 | `true`／`false` | 也可用 `--no-update` 單次停用 |
| `DECEPTICON_HOME` | 安裝與設定目錄 | `~/.decepticon` | 換目錄等於多份獨立安裝 |
| `DECEPTICON_MODEL` | 指定模型（覆寫設定檔） | `anthropic/claude-...`、`openai/...` | 留空則依設定檔自動選 |
| `DECEPTICON_LITELLM_MODELS` | 臨時註冊額外模型 | JSON／逗號清單 | 給 Groq、Azure、Bedrock 等用 |
| `LANGGRAPH_PORT` | LangGraph 對外連接埠 | `2024` | 被占用時改這裡 |
| `NEO4J_PASSWORD` | 知識圖譜密碼 | 任意字串 | 首次 onboard 會產生 |
| `OLLAMA_API_BASE`／`OLLAMA_MODEL` | 本機模型 | `http://host.docker.internal:11434` | 不需 API 金鑰 |
| `ANTHROPIC_API_KEY` 等 | 各供應商金鑰 | `sk-...` | 預留字串（`your-...-key-here`）會被忽略 |
| `BLUE_SENSOR_URL`／`BLUE_MONITOR_URL` | Blue Cell 感測器端點 | `http://127.0.0.1:18081` | 只有藍隊監控要用 |
| `DECEPTICON_LLM__PROXY_URL`／`SANDBOX_URL` | 以 SDK 模式執行時的服務位址 | `http://localhost:4000` | 只有 `pip install decepticon` 才需要 |

### 模型設定檔怎麼選

| 設定檔 | 各代理等級 | 適合誰 |
|---|---|---|
| `eco`（預設） | 協同／利用／修補／分析為 HIGH，執行為 MID，偵察為 LOW | 一般使用、成本可控 |
| `max` | 全部 HIGH | 高價值目標、要最好結果 |
| `test` | 全部 LOW | 開發與 CI，壓低成本 |

---

## 3. 啟動

```bash
decepticon
```

會拉起核心管理平面（LiteLLM、PostgreSQL、Neo4j、Skillogy、LangGraph、sandbox），
然後進入終端機對話介面。第一次啟動要下載映像，會比較久。

> 提示：加 `--no-alt-screen` 可改用終端機原生捲動；加 `--no-mouse` 可停用滑鼠。

---

## 4. 第一次授權測試任務

1. 在 CLI 選 **New engagement**（新授權測試任務），輸入名稱。
2. 填**範圍內目標**與**排除項目**（網域、URL、IP／CIDR、支援的 glob）。
3. 確認正規化後的範圍清單 —— 被丟棄或展開的項目會標示出來。
4. 啟動器會寫出機器可讀的 RoE 並進入紅隊模式。

想在動手前先討論策略，可手動打 `/plan` 進入計畫模式（非必要）。

> ⚠️ 只對你擁有、或已取得**明確書面授權**的系統執行。詳見 README 的免責聲明。

---

## 5. 看結果

```bash
decepticon logs --engagement <任務名稱>    # 重播對話紀錄
decepticon status                          # 服務狀態
decepticon kg-health                       # 知識圖譜（Neo4j）健康檢查
```

對話中常用的斜線指令：

| 指令 | 用途 |
|---|---|
| `/help`、`/shortcuts` | 指令清單、鍵盤快速鍵 |
| `/model`、`/status` | 換模型、看工作階段狀態 |
| `/pwd` | 看目前工作區路徑 |
| `/sessions`、`/resume` | 瀏覽已存對話、接續暫停的工作 |
| `/queue` | 看排隊中的訊息 |
| `/new` | 在同一任務另開新對話（會保留舊對話） |
| `/quit`、`/exit` | 離開對話（服務繼續跑） |

按鍵：Enter 送出、Shift+Enter／Alt+Enter／Ctrl+J 換行、Ctrl+R 搜尋歷史、
Ctrl+A 移到行首、Esc 先關面板再暫停工作。

---

## 6. 停止與清理

```bash
decepticon stop     # 停服務，保留資料
make clean          # 停服務 + 刪 volume（全部重來）
```

---

## 7. 一頁速查

```bash
decepticon                  # 啟動並進入對話
decepticon onboard          # 重新設定金鑰與模型
decepticon status           # 服務狀態（--json 給腳本用）
decepticon doctor           # 唯讀診斷（--deep 加容器內檢查）
decepticon logs langgraph   # 追日誌
decepticon update           # 更新
decepticon stop             # 停止
```

下一步：[TROUBLESHOOTING.md](TROUBLESHOOTING.md) · 完整英文文件 [`docs/`](../docs/)
