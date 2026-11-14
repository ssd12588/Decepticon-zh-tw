[![繁體中文](https://img.shields.io/badge/Language-繁體中文-green?style=for-the-badge)](README.md)
[![English](https://img.shields.io/badge/Language-English-blue?style=for-the-badge)](https://github.com/BitterSecurity/Decepticon/blob/main/README.md)
[![한국어](https://img.shields.io/badge/Language-한국어-red?style=for-the-badge)](README_KO.md)

<div align="center">
  <img src="assets/logo_banner.png" alt="Decepticon Logo">
</div>

<h1 align="center">Decepticon — 自主紅隊代理</h1>

<p align="center"><i>「又一個 AI 駭客？我猜它會跑 nmap，然後寫一份報告。」</i></p>

<div align="center">

<a href="https://trendshift.io/repositories/15987?utm_source=repository-badge&amp;utm_medium=badge&amp;utm_campaign=badge-repository-15987" target="_blank" rel="noopener noreferrer"><img src="https://trendshift.io/api/badge/repositories/15987" alt="BitterSecurity%2FDecepticon | Trendshift" width="250" height="55"/></a>
<a href="https://trendshift.io/repositories/15987?utm_source=trendshift-badge&amp;utm_medium=badge&amp;utm_campaign=badge-trendshift-15987" target="_blank" rel="noopener noreferrer"><img src="https://trendshift.io/api/badge/trendshift/repositories/15987/daily?language=Python" alt="BitterSecurity%2FDecepticon | Trendshift" width="250" height="55"/></a>

<br/>

<a href="https://github.com/BitterSecurity/Decepticon/blob/main/LICENSE">
  <img src="https://img.shields.io/github/license/BitterSecurity/Decepticon?style=for-the-badge&color=blue" alt="License: Apache 2.0">
</a>
<a href="https://github.com/BitterSecurity/Decepticon/stargazers">
  <img src="https://img.shields.io/github/stars/BitterSecurity/Decepticon?style=for-the-badge&color=yellow" alt="Stargazers">
</a>
<a href="https://github.com/BitterSecurity/Decepticon/graphs/contributors">
  <img src="https://img.shields.io/github/contributors/BitterSecurity/Decepticon?style=for-the-badge&color=orange" alt="Contributors">
</a>

<br/>

<a href="https://discord.gg/TZUYsZgrRG">
  <img src="https://img.shields.io/badge/Discord-Join%20Us-7289DA?logo=discord&logoColor=white&style=for-the-badge" alt="加入我們的 Discord">
</a>
<a href="https://decepticon.red">
  <img src="https://img.shields.io/badge/Website-decepticon.red-brightgreen?logo=vercel&logoColor=white&style=for-the-badge" alt="官方網站">
</a>
<a href="https://docs.decepticon.red">
  <img src="https://img.shields.io/badge/Docs-docs.decepticon.red-8B5CF6?logo=bookstack&logoColor=white&style=for-the-badge" alt="文件">
</a>
<a href="https://app.decepticon.red">
  <img src="https://img.shields.io/badge/Live%20App-app.decepticon.red-FF2D55?logo=rocket&logoColor=white&style=for-the-badge" alt="線上託管版">
</a>

</div>

<br/>

<div align="center">
  <p><strong>商業版產品示範</strong></p>
  <video src="https://github.com/user-attachments/assets/7a9b233c-fd72-4633-a408-03953c345a15" width="800" controls></video>
</div>


<div align="center">

### ☁️ 不想自己架？**Decepticon 雲端版已經上線。**

跳過 Docker 設定 —— 直接在瀏覽器裡執行自主紅隊授權測試任務。

<a href="https://app.decepticon.red">
  <img src="https://img.shields.io/badge/Launch%20the%20Live%20App-app.decepticon.red-FF2D55?logo=rocket&logoColor=white&style=for-the-badge" alt="開啟 app.decepticon.red 線上版">
</a>

</div>

<div align="center">
  <p><strong>開源 CLI 示範</strong></p>
  <video src="https://github.com/user-attachments/assets/b3fd40d8-e859-4a39-97f4-bd825694ad96" width="800" controls></video>
</div>

---

## 關於本漢化版

**本漢化版由 [阿樂](https://github.com/ssd12588) 負責繁體中文漢化與維護。**

- 本倉庫是 [BitterSecurity/Decepticon](https://github.com/BitterSecurity/Decepticon) 的**繁體中文漢化版**，依 Apache-2.0 授權散布，原始版權聲明與 [`LICENSE`](LICENSE) 均完整保留。
- **已漢化**：終端機（CLI）主畫面與第一批指令的說明與錯誤訊息（`--help`、`status`、`logs`、`doctor`、`update`…）、`doctor` 診斷輸出與復原指令、`AGENTS.md`、本 README，以及 `docs-zh/` 完整中文入門包。
- **進行中**：`decepticon onboard` 初始化精靈、TUI 面板與工作階段選擇器、Python 無介面 CLI、`CONTRIBUTING.md`、`docs/` 全量文件（進度見 [PROJECT-STATUS](docs-zh/PROJECT-STATUS.md)）。
- **中文入門文件**：[10 分鐘上手](docs-zh/QUICKSTART.md) · [常見錯誤排錯](docs-zh/TROUBLESHOOTING.md) · [術語對照表](docs-zh/glossary.md) · [專案狀態（交接用）](docs-zh/PROJECT-STATUS.md)
- **Windows 一鍵啟動**：[`start-decepticon.ps1`](start-decepticon.ps1)：檢查 Docker、記住模型與金鑰、給出中文的友善錯誤提示。
- **繁體中文輸出**：上游已內建語言機制，設定 `DECEPTICON_LANGUAGE=tw` 之後，代理所有「給人看」的說明（進度、總結、狀態、錯誤訊息）都會改用繁體中文，而工具名稱、CLI 指令、檔案路徑、程式碼與技術縮寫仍保留英文。
- **刻意不改**：代理提示詞（`packages/decepticon/decepticon/agents/prompts/**`）、技能知識包（`skills/**`、`SKILL.md`）、程式碼識別字、協定欄位、JSON 鍵、狀態值、CLI 參數與環境變數、payload、HTTP 請求與回應 —— 這些改了會影響 AI 判斷品質或直接弄壞程式。
- 與上游同步：

  ```bash
  git remote add upstream https://github.com/BitterSecurity/Decepticon.git
  git fetch upstream
  git merge upstream/main
  ```

- 漢化相關問題請在本倉庫開 issue，無需打擾上游維護者。

---

## 安裝

**先決條件**：[Docker](https://docs.docker.com/get-docker/) 與 Docker Compose v2。
支援 macOS（Apple Silicon + Intel）、Linux（amd64 + arm64）、Windows（amd64 + arm64）—— Windows 可用原生 PowerShell，或透過 WSL2（Ubuntu / Kali）。

**macOS / Linux / WSL2**
```bash
curl -fsSL https://decepticon.red/install | bash
decepticon onboard   # 互動式設定精靈（供應商、API 金鑰、模型設定檔）
decepticon           # 啟動核心服務堆疊並進入終端機 CLI
```

預設啟動會拉起核心管理平面（LiteLLM、PostgreSQL、Neo4j、Skillogy、LangGraph、sandbox），並啟動終端機 CLI。專用工作負載（BloodHound CE、Sliver C2、REA/Ghidra……）採隨選啟動 —— 由協同器透過 `ops_start("ad")` 之類的呼叫產生。終端機 CLI 是開源版的操作介面；託管的網頁主控台屬於 [Decepticon Cloud](https://decepticon.red)。

**Windows（原生 PowerShell）**
```powershell
irm https://decepticon.red/install.ps1 | iex
decepticon onboard
decepticon
```

→ **[快速上手](docs/getting-started.md)** · **[完整安裝導引](docs/setup-guide.md)**

### 互動式 CLI 內

輸入 `/` 可瀏覽指令，用上下鍵選擇後按 Enter 執行。Tab 會補完選中的指令；
別名也會出現在說明與自動補完中。

| 指令 | 用途 |
|------|------|
| `/help`、`/shortcuts` | 瀏覽指令或鍵盤快速鍵 |
| `/model`、`/status` | 選擇已設定的模型，或檢視本次工作階段的連線 |
| `/plan` | 進入選用的計畫模式，接著送出你的規劃訊息 |
| `/mode red`、`/mode blue` | 回到紅隊，或切換到 Blue 調查 |
| `/new` | 閒置時在同一授權測試任務與模式下開始全新對話 |
| `/sessions`、`/resume` | 瀏覽已儲存的對話，或接續暫停中的工作 |
| `/version`（`/about`）、`/pwd`（`/cwd`） | 查看執行中的版本與工作區路徑 |
| `/quit`、`/exit` | 離開對話；服務與伺服器端工作仍會繼續執行 |

Enter 送出；Shift+Enter、Alt+Enter 或 Ctrl+J 換行。在第一／最後一行輸入時，
上下鍵會叫回歷史輸入；Ctrl+R 可搜尋歷史。Esc 會先關閉取得焦點的介面，再暫停
進行中的工作；`/resume` 可接續。Ctrl+A 移到目前行首。`/new` 之前可用 `/queue`
檢視等待中的訊息。

`/new` 會保留已儲存的對話與終端機捲動內容；它不會建立另一個授權測試任務，
也不會改變目標授權。`/plan` 與 `/new` 都不接受參數。

上述快速鍵與斜線指令描述的是目前的 `main`，包含尚未發行的變更。合併原始碼
不會更新已安裝的映像。請見[版本紀錄](CHAGELOG.md)、[完整 CLI 參考](docs/cli-reference.md#slash-commands)
與[升級行為](docs/cli-experience.md#upgrade-behavior)。

### 從 Claude Code 或 Codex 使用

啟動 Decepticon 之後，在你使用的程式開發代理註冊它的 MCP 伺服器：

```bash
claude mcp add --scope user decepticon -- decepticon mcp serve
codex mcp add decepticon -- decepticon mcp serve
```

用你已安裝的那個代理對應的指令，然後執行 `decepticon skill install`，把操作
指引加到兩個代理。Agent Skill、工具工作流程與原始碼安裝方式的說明，請見
[外部代理指南](docs/integrations/external-agents.md)。

### 以程式庫使用（pip）

想在代理之上再往上建 —— 產品、研究整合或自製協同器？從 PyPI 安裝 SDK：

```bash
pip install decepticon              # 核心 SDK
pip install "decepticon[neo4j]"     # 加上知識圖譜的攻擊鏈工具
```

`decepticon` 是**用戶端 SDK**：它提供代理工廠、middleware、工具與技能，並把 LLM
呼叫與沙箱執行透過 HTTP 導向執行期服務（`DECEPTICON_LLM__PROXY_URL`、`SANDBOX_URL`）。
實際執行代理仍需要這些服務 —— 請用上面的 Docker 堆疊，或把 URL 指向你自己的等效服務。
工廠覆寫介面、宣告式 `PluginBundle` 外掛與安全閘門，請見 **[以程式庫使用 Decepticon](docs/library-usage.md)**。

---

## 💖 支持 Decepticon

[![Sponsor](https://img.shields.io/badge/Sponsor-Decepticon-red?style=for-the-badge&logo=github)](https://github.com/sponsors/PurpleCHOIms)

我們正把 Decepticon 打造成 AI 驅動威脅環境的**攻擊疫苗**。如果你也相信自主紅隊
是強化防禦的一條路，歡迎支持這個專案。

---

## 基準測試

<div align="center">
  <img src="assets/benchmark/decepticon_donut.png" alt="Decepticon —— XBOW 通過率 102/104（98.08%）" width="560">
</div>

| 基準測試 | 難度 | 通過率 |
|----------|------|--------|
| [XBOW validation-benchmarks](https://github.com/BitterSecurity/xbow-validation-benchmarks) | 簡單（Level 1） | **45 / 45**（100 %） |
| [XBOW validation-benchmarks](https://github.com/BitterSecurity/xbow-validation-benchmarks) | 中等（Level 2） | **50 / 51**（98.0 %） |
| [XBOW validation-benchmarks](https://github.com/BitterSecurity/xbow-validation-benchmarks) | 困難（Level 3） | **7 / 8**（87.5 %） |
| [XBOW validation-benchmarks](https://github.com/BitterSecurity/xbow-validation-benchmarks) | **全部等級** | **102 / 104**（98.08 %） |

- **[各題目完整索引、攻擊類型矩陣與 LangSmith 追蹤紀錄](benchmark/results/README.md)**
- **[與其他 AI 滲透測試代理的比較（Strix、PentestGPT、MAPTA、Cyber-AutoAgent、XBOW 商業版……）](docs/benchmark-comparison.md)**

---

## Decepticon 是什麼？

「AI + 駭客」這個領域充滿了跑跑 nmap、印一份報告的展示品。這裡講的不是那種東西。

**Decepticon 是專業的自主紅隊代理。** 它會像真正的對手那樣執行貼近現實的攻擊鏈
—— 偵察、利用、提權、橫向移動、C2 —— 而不是掃描器的做法。

每一場紅隊授權測試任務都從操作員核准的目標與授權規則（RoE）開始。選用的計畫模式
會額外加入 Soundwave 的討論與七份規劃文件，並與唯讀的授權規則並存。執行協同器
擁有 OPPLAN；紅隊流程不強制要求使用計畫模式。

→ **[授權測試任務流程深入解析](docs/engagement-workflow.md)**

---

## 為什麼選 Decepticon？

**真實攻擊鏈，不是勾表式掃描。** Decepticon 會讀 OPPLAN，沿著任何可行的路徑推進
目標 —— 轉向、調整、串接各種手法。

**真的能操作互動式 shell。** 真正的攻擊工具都是互動式的（`msfconsole`、
`sliver-client`、`evil-winrm`）。Decepticon 把每道指令都跑在常駐的 tmux 工作階段裡，
並自動偵測提示字元 —— 當工具掉進互動式提示時，代理可以直接送後續指令，不必繞路。

**強化的沙箱隔離。** 所有指令都在 Kali Linux 沙箱內、於專用作業網路（`sandbox-net`）
執行，與管理平面（`decepticon-net`）分離。LangGraph 透過 Docker socket 驅動沙箱。
→ **[架構](docs/architecture.md)**

**攻擊為防禦服務。** 規劃中的[攻擊疫苗](docs/offensive-vaccine.md)循環，會透過
「攻擊 → 防禦 → 驗證」的週期，把發現轉成防禦面的改進。

---

## 架構

<div align="center">
  <img src="assets/decepticon_infra.svg" alt="Decepticon 基礎架構" width="680">
</div>

雙網路設計。**常駐**的管理平面（LiteLLM、PostgreSQL、Skillogy、LangGraph）與常駐的
沙箱平面在整場授權測試任務期間都會保持運行；其餘都是**動態產生** —— 專用工作負載
（BloodHound CE、Sliver C2、REA/Ghidra……）只有在協同器呼叫 `ops_start(...)` 時才會啟動
（見 [ADR-0006](docs/adr/0006-agent-driven-container-lifecycle.md)）。網路配置：管理平面在
`decepticon-net`；沙箱、C2 伺服器與目標在 `sandbox-net`。Neo4j 同時接兩張網，讓位於
管理平面的代理能保存從沙箱內寫入的發現。

→ **[架構深入解析](docs/architecture.md)** · **[知識圖譜](docs/knowledge-graph.md)**

---

## 代理

16 個專責代理，依攻擊鏈階段分組，每個目標項目都使用全新的上下文視窗 —— 不累積雜訊。

協同運作 · 偵察 · 利用 · 後滲透 · 漏洞研究 · 領域專家（AD、Cloud、智慧合約、逆向、分析）。

→ **[完整代理名冊與 middleware 堆疊](docs/agents.md)**

---

## 模型與供應商

以等級為基礎、會感知憑證的備援鏈。你依優先順序宣告自己有哪些憑證；Decepticon 就
據此在每個等級建出「主要 → 備援」的鏈。

| 設定檔 | 各代理的等級 | 適用情境 |
|--------|--------------|----------|
| **eco**（預設） | 依代理而定（協同器／利用／修補／分析為 HIGH，執行為 MID，偵察／Soundwave 為 LOW） | 生產環境 |
| **max** | 所有代理皆 HIGH | 高價值目標 |
| **test** | 所有代理皆 LOW | 開發／CI |

**對應等級的供應商**：Anthropic、OpenAI、Google Gemini、MiniMax、DeepSeek、xAI、
Mistral、OpenRouter、Nvidia NIM、Ollama（本機）。
**訂閱制 OAuth**：Claude Max/Pro/Team、ChatGPT Pro/Plus/Team、Gemini Advanced、
Copilot Pro、SuperGrok、Perplexity Pro。

用 `decepticon onboard` 設定。→ **[完整模型參考與備援範例](docs/models.md)**

---

## 文件

| 主題 | 文件 |
|------|------|
| 安裝與第一次授權測試任務 | [快速上手](docs/getting-started.md) |
| 完整設定、OAuth 與供應商 | [安裝導引](docs/setup-guide.md) |
| 程式開發代理的 CLI 包裝 | [Decepticon skill](skills/decepticon/SKILL.md) |
| 文件索引與儲存庫導覽 | [文件索引](docs/README.md) |
| 所有 CLI 指令與鍵盤快速鍵 | [CLI 參考](docs/cli-reference.md) |
| 所有 `make` 目標 | [Makefile 參考](docs/makefile-reference.md) |
| 代理名冊與 middleware | [代理](docs/agents.md) |
| REA/Ghidra 逆向工程 | [REA 逆向](docs/features/rea-reversing.md) |
| 本機防禦監控 | [Blue Cell](docs/features/blue-cell.md) |
| 模型設定檔與備援鏈 | [模型](docs/models.md) |
| 技能系統與格式規範 | [技能](docs/skills.md) |
| 擴充能力、技能、工作流程與用戶端 | [擴充契約](docs/design/extension-contract.md) |
| 系統架構與網路隔離 | [架構](docs/architecture.md) |
| Neo4j 知識圖譜 | [知識圖譜](docs/knowledge-graph.md) |
| 端到端授權測試任務流程 | [授權測試任務流程](docs/engagement-workflow.md) |
| 攻擊疫苗循環 | [攻擊疫苗](docs/offensive-vaccine.md) |
| 版本紀錄 | [更新日誌](CHAGELOG.md) |
| 參與 Decepticon 開發 | [參與貢獻](docs/contributing.md) |

---

## 參與貢獻

```bash
git clone https://github.com/BitterSecurity/Decepticon.git
cd Decepticon
make dogfood  # 用本地程式碼走完整開源體驗（啟動器 → onboard → CLI）
make dev      # 後端熱重載（compose watch）—— 日常開發循環
```

→ **[參與貢獻指南](docs/contributing.md)**

---

## 社群

加入 [Discord](https://discord.gg/TZUYsZgrRG) —— 提問、分享授權測試紀錄、討論手法。

---

## 免責聲明

未取得系統擁有者的明確書面授權，請勿在任何系統或網路上使用本專案。未經授權存取
電腦系統屬於違法行為。你必須為自己的行為負全部責任。作者與貢獻者對不當使用不負
任何責任。

---

## 授權

[Apache-2.0](LICENSE)

---

## 致謝

感謝 Decepticon 所建構、整合與學習的這些專案，其維護者與貢獻者。

| 專案 | 在 Decepticon 中的角色 |
|------|------------------------|
| [REA](https://github.com/morluto/rea) 與 [Ghidra](https://github.com/NationalSecurityAgency/ghidra) | 選用的逆向工程工作負載與有證據支撐的二進位分析。[整合說明](docs/features/rea-reversing.md)。 |
| [Deep Agents](https://github.com/langchain-ai/deepagents) 與 [LangGraph](https://github.com/langchain-ai/langgraph) | 代理執行框架、協同運作、串流與持續性執行狀態。 |
| [LiteLLM](https://github.com/BerriAI/litellm) | 模型閘道與供應商路由。 |
| [insane-search](https://github.com/fivetaku/insane-search) | 為沙箱網頁存取改作的上游引擎。[改作聲明](packages/decepticon/decepticon_sandbox/web/NOTICE.md)。 |
| [Playwright](https://github.com/microsoft/playwright) | 網頁調查用的瀏覽器自動化。 |
| [Caddy](https://github.com/caddyserver/caddy) 與 [Fluent Bit](https://github.com/fluent/fluent-bit) | 本機 Blue 感測器入口與日誌收集。[Blue Cell 架構](docs/features/blue-cell.md)。 |
| [Ink](https://github.com/vadimdemedes/ink)、[Bubble Tea](https://github.com/charmbracelet/bubbletea) 與 [Huh](https://github.com/charmbracelet/huh) | 互動式 CLI 呈現與啟動器表單。 |
| [PayloadsAllTheThings](https://github.com/swisskyrepo/PayloadsAllTheThings)、[h4cker](https://github.com/The-Art-of-Hacking/h4cker) 與 [Awesome-Hacking](https://github.com/Hack-with-Github/Awesome-Hacking) | 安全知識與學習參考。[姓名標示細節](THIRD_PARTY_LICENSES.md)。 |
| [XBOW validation benchmarks](https://github.com/XBOW-engineering/validation-benchmarks) | 基準測試整合所使用的評測任務。[結果與方法](benchmark/results/README.md)。 |

在 CLI 互動設計上，我們也研究了
[Claude Code](https://code.claude.com/docs/en/interactive-mode)、
[Codex CLI](https://github.com/openai/codex) 與
[Gemini CLI](https://github.com/google-gemini/gemini-cli)。
由此產生的取捨記錄在 [CLI 參考檢視](docs/cli-experience.md)。

上游授權與姓名標示聲明仍然適用；請見
[第三方授權](THIRD_PARTY_LICENSES.md) 與改作程式碼旁的聲明。
