# 漢化術語對照表（Glossary）

> 用途：未來任何人（或 AI）要新增／修改漢化內容時，請**照這張表翻**，避免同一個詞在不同地方譯法不一致。
>
> 原則：
> 1. **技術專有名詞保留英文**（Docker、API、SARIF、JWT、XSS…），因為使用者會在網路上搜尋原文。
> 2. **產品名不翻**（Decepticon、LiteLLM、LangGraph、Neo4j、Docker、Playwright、Caddy、Ghidra）。
> 3. **指令、參數、路徑、環境變數一律不翻**（`decepticon onboard`、`--json`、`DECEPTICON_LANGUAGE`）。
> 4. **狀態值不翻**（`running`、`completed`、`failed`），只翻「顯示用」標籤。
> 5. 中文用**繁體**，標點用全角（，。（）；：），破折號用 `——`。

---

## 核心概念

| 英文 | 統一譯法 | 備註 |
|---|---|---|
| Engagement | 授權測試任務（engagement） | 本專案最核心的名詞，首次出現夾註英文 |
| Rules of Engagement (RoE) | 授權規則（RoE） | 授權書、時間窗、排除範圍 |
| Operator | 操作員 | 不用「操作者」 |
| Target | 目標 | |
| Scope | 範圍 | |
| Out of Scope | 不在範圍內 | 依語境可作「排除範圍」 |
| Finding | 發現（finding） | 名詞；不譯為「漏洞」以免與 vulnerability 混淆 |
| Vulnerability | 漏洞 | |
| Exploit | 利用（exploit） | 動詞／名詞皆可；不譯「利用程式」以外的變化 |
| Exploitation | 利用 | |
| Post-exploitation | 後滲透 | |
| Exploit chain | 攻擊鏈 | |
| Payload | payload | **不翻**（社群慣用原文） |
| Proof of Concept (PoC) | 概念驗證（PoC） | 首次出現夾註英文 |
| Severity | 嚴重度 | 不用「危險等級」 |
| Critical / High / Medium / Low / Info | 重大／高／中／低／資訊 | 值本身（`critical`）不翻 |
| Remediation | 修補建議 | |
| Verification | 驗證 | |
| False Positive | 誤報 | |
| Attack Surface | 攻擊面 | |
| Reconnaissance | 偵察 | |
| Bypass | 繞過 | |
| Evidence | 證據 | |
| Privilege Escalation | 提權 | |
| Lateral Movement | 橫向移動 | |
| Persistence | 持續潛伏 | 依語境可作「持久化」 |
| Objective | 目標項目 | 作戰計畫中的一條目標 |
| Opplan（operation plan） | 作戰計畫（OPPLAN） | 專有名詞，不拆譯 |

## 代理與流程

| 英文 | 統一譯法 | 備註 |
|---|---|---|
| Agent | 代理 | 不用「智能體」 |
| Sub-agent | 子代理 | |
| Agent roster | 代理名冊 | |
| Orchestration | 協同運作 | |
| Turn | 回合 | |
| Session | 工作階段 | |
| Thread | 對話串（thread） | LangGraph 的 thread，首次夾註 |
| Transcript | 對話紀錄 | |
| Workspace | 工作區 | |
| Run | 執行 | 名詞；例：本次執行 |
| Scan | 掃描 | |
| Headless | 無介面 | 不用「無頭」 |
| Non-interactive | 非互動 | |
| TUI | TUI | 不翻 |
| Launcher | 啟動器（launcher） | 指令名不動 |
| Onboard | 初始化設定（onboard） | 指令名不動 |
| Doctor | 環境診斷（doctor） | 指令名不動 |
| Health check | 健康檢查 | |
| Panel | 面板 | |
| Status line | 狀態列 | |
| Empty state | 空狀態 | |
| Keybinding / Shortcut | 快速鍵 | |
| Plan mode | 計畫模式 | 指令 `/plan` 不動 |
| Todo | 待辦 | |

## 能力與知識

| 英文 | 統一譯法 | 備註 |
|---|---|---|
| Skill | 技能（skill） | 目錄名 `skills/` 與 `SKILL.md` 不動 |
| Skillogy | Skillogy | 專有名詞，不翻 |
| Knowledge graph | 知識圖譜 | |
| Model profile | 模型設定檔 | |
| Provider | 供應商 | |
| Model catalog | 模型目錄 | |
| Fallback | 備援 | 不用「降級」 |
| Blue Cell | Blue Cell | 藍隊監控子系統，**不翻** |
| Offensive Vaccine | 攻擊疫苗循環 | 專案自有概念，首次夾註 |

## 設定與環境

| 英文 | 統一譯法 | 備註 |
|---|---|---|
| Environment Variable | 環境變數 | 變數名本身不翻 |
| API Key | API 金鑰 | |
| Subscription OAuth | 訂閱制 OAuth | |
| Tier | 等級 | 例：生產等級 |
| Stack | 服務堆疊 | |
| Service | 服務 | |
| Container | 容器 | |
| Sandbox | 沙箱 | |
| Configuration | 設定 | |
| Config file | 設定檔 | |
| Default | 預設 | |
| Deprecated | 已棄用 | |
| Optional | 選用 | |
| Required | 必要 | |
| Telemetry | 遙測 | |
| Log | 日誌 | |
| Exit code | 結束碼 | |

## 安全分類（保留英文縮寫）

| 英文 | 處理方式 |
|---|---|
| XSS / CSRF / SSRF / XXE / RCE / SSTI / IDOR | 保留縮寫，首次出現可加中文說明 |
| JWT / SARIF / CVSS / CVE / PoC | 保留原文 |
| SQL Injection | SQL 注入 |
| Command Injection | 指令注入 |
| Path Traversal | 路徑穿越 |
| Insecure Deserialization | 不安全的反序列化 |
| Prototype Pollution | 原型污染 |
| Open Redirect | 開放重新導向 |
| Broken Access Control | 存取控制失效 |
| Session Fixation | 工作階段固定 |
| Subdomain Takeover | 子網域接管 |
| Information Disclosure | 資訊外洩 |
| Supply chain | 供應鏈 |

## 介面與檔案

| 英文 | 統一譯法 |
|---|---|
| Endpoint | 端點 |
| Header | 標頭 |
| Request / Response | 請求／回應 |
| Proxy | 代理伺服器 |
| Intercept | 攔截 |
| Replay | 重播 |
| Cookie | Cookie（不翻） |
| Traceback | 錯誤追蹤（traceback） |
| Pull Request / PR | PR |
| Issue | issue（小寫，不翻） |
| Repository | 儲存庫 |
| Branch | 分支 |
| Commit | commit（不翻） |
| Fork | fork（不翻） |
| Upstream | 上游 |
| Release | 發行版 |
| Changelog | 更新日誌 |

## 標點與排版

| 規則 | 正確 | 錯誤 |
|---|---|---|
| 破折號 | `——`（兩個全角） | `-` 或 `--` |
| 頓號 | `、` | `,` |
| 括號 | `（）` 全角（程式碼內除外） | `()` |
| 引號 | `「」` 或 `**粗體**` | `""` |
| 中英之間 | 不加空格（`使用 Docker`） | `使用 Docker` 之間加半形空格 |
| 數字與單位 | `5 分鐘`、`10 個` | `5分鐘` |
| 省略號 | `…` | `...`（程式碼內除外） |

---

## 維護方式

1. 新增譯法前，先在本表搜尋是否已有對應詞。
2. 若必須新增，請同步更新本表（一個 PR 只改一處，避免漂移）。
3. 漢化範圍與排除清單見 [README「關於本漢化版」](../README.md#關於本漢化版)。
