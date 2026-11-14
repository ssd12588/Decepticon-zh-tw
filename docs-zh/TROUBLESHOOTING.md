# 排錯手冊（TROUBLESHOOTING）

把**中文錯誤訊息**對照回原因與解法。第一步永遠是：

```bash
decepticon doctor
```

它會列出 `檢查項目: 狀態 —— 說明`，狀態只有三種：`正常`、`失敗`、`選用`。
每一項失敗下面都會有一行 `復原方式：`，可以直接照著做。

---

## 一、CLI 錯誤訊息對照

| 中文訊息 | 原因 | 解法 |
|---|---|---|
| `未知的指令 "xxx"；你是不是想用 "yyy"？` | 指令打錯字（含拼字建議） | 照建議執行；或跑 `decepticon help` |
| `無法辨識的參數；請改用 --resume=THREAD_ID…` | 直接在 `decepticon` 後面接了非指令的參數 | 接續對話用 `--resume`、`--resume=<id>` 或 `--continue` |
| `Decepticon 尚未安裝於 <路徑>；請先執行 decepticon onboard：<原因>` | `mcp serve` 找不到 `docker-compose.yml` 或 `.env` | 先跑 `decepticon onboard`；或用 `DECEPTICON_HOME` 指向正確安裝目錄 |
| `必要檢查未通過；請參考 doctor 報告中的復原指令` | 必要檢查（docker／compose／config／auth／langgraph）有失敗 | 依 `doctor` 輸出逐項修 |
| `授權測試任務名稱格式不正確` | `--engagement` 只接受小寫英數與 `-`，長度 3～64，前後不可為 `-` | 例：`acme-external-2026` |
| `--engagement 需要搭配 --deep` | 用了 `doctor --engagement` 卻沒加 `--deep` | 加上 `--deep`，或去掉 `--engagement` |
| `服務參數與 --engagement 不能同時使用` | `decepticon logs svc --engagement x` | 兩種模式擇一 |
| `未知的服務 "xxx"；已記載的服務：…` | 服務名不在 Compose 設定內 | 用 `decepticon status` 看實際服務名 |
| `驗證服務 "xxx" 失敗：<原因>` | 讀取 Compose 服務清單失敗 | 確認在安裝目錄、`.env` 可讀 |
| `檢查更新失敗：<原因>` | 連不到 GitHub API（網路、代理、額度） | 檢查網路／代理；稍後重試 `decepticon update` |
| `設定檔同步失敗：<原因>` | 下載或寫入 `docker-compose.yml` 等設定失敗 | 檢查 `~/.decepticon` 寫入權限 |
| `映像拉取失敗：<原因>` | Docker 映像拉取失敗（磁碟、登入、網路） | `docker login`、清空間後重試 |
| `infracontrol 停止失敗：<原因>` | opscontrol 常駐程式沒在跑或沒權限 | 多為無害訊息；必要時 `decepticon infracontrol status` |
| `取消回合 <id>：<原因>` | 取消時 LangGraph 回錯 | 確認服務在跑、`thread-id` 正確 |
| `授權測試任務 <名稱> 目前沒有執行中的回合 —— 無需取消。` | 該任務本來就沒在跑 | 不是錯誤，忽略即可 |
| `已安裝 Decepticon Skill 於 <路徑>` | 正常輸出 | — |

---

## 二、`decepticon doctor` 檢查項目對照

| 項目 | 失敗代表什麼 | 怎麼修 |
|---|---|---|
| `docker` | PATH 上找不到 `docker` | 安裝 Docker Desktop／Docker Engine，重開終端機 |
| `docker_daemon` | Docker 裝了但沒在跑 | 開 Docker Desktop；Linux 用 `sudo systemctl start docker` |
| `compose` | 缺 Compose v2 外掛 | 安裝 compose plugin，`docker compose version` 應有輸出 |
| `config` | 缺 `.env` 或 `docker-compose.yml` | 跑 `decepticon onboard` |
| `auth` | 沒設定任何驗證方式 | 跑 `decepticon onboard` 選一種憑證 |
| `langgraph` | 主機連不到 LangGraph，或 assistant 未註冊 | 看 `decepticon logs langgraph`；確認 `LANGGRAPH_PORT` 沒被占用 |
| `opscontrol_socket`（選用） | opscontrol Unix socket 不存在 | `decepticon infracontrol status`；用 `decepticon start --no-update` 啟動 |
| `opscontrol`（選用） | 常駐程式無回應 | 同上；或看 `decepticon logs` |
| `blue_sensor`／`blue_monitor`（選用） | 沒設端點，或端點不通 | 不需要藍隊監控就忽略；需要時設 `BLUE_SENSOR_URL`／`BLUE_MONITOR_URL` |

---

## 三、平台與環境問題

### 中文顯示成亂碼或方框（Windows）

本漢化版啟動時會把主控台切到 UTF-8（CP65001）。若仍異常：

```powershell
chcp                    # 應該要顯示 65001
chcp 65001              # 不是的話手動切
```

- 建議用 **Windows Terminal** 搭配 **PowerShell 7（pwsh）**，不要用舊的 `cmd.exe` 主控台主機。
- 若字型缺中文，於 Windows Terminal 設定換成 `Cascadia Mono`／`Microsoft JhengHei`。

### `decepticon` 指令找不到（Windows）

安裝後請**完全關閉終端機再重開**，PATH 才會生效。或用完整路徑呼叫。

### 執行 `start-decepticon.ps1` 時中文變亂碼或直接報語法錯誤

原因是 `.ps1` 被存成「UTF-8 無 BOM」時，Windows PowerShell 5.1 會用 ANSI（CP950）解讀。
本倉庫的腳本已存成 **UTF-8 with BOM**，若你自行編輯後失效，請重新存檔：

```powershell
$p = ".\start-decepticon.ps1"
$c = Get-Content -Raw -Encoding UTF8 $p
[IO.File]::WriteAllText($p, $c, (New-Object System.Text.UTF8Encoding($true)))
```

或者直接用 PowerShell 7（`pwsh`）執行，它預設以 UTF-8 讀檔。

### 埠被占用

`langgraph: 失敗` 且訊息提到 `connectex` 或 `connection refused`：

```bash
# 換一個對外埠
LANGGRAPH_PORT=2025 decepticon
```

同時清掉舊堆疊：`decepticon stop`。

### Docker 資源不足（映像下載卡住、容器一直重啟）

- 第一次啟動要抓數 GB 映像，請預留 20 GB 以上磁碟。
- Docker Desktop 建議給 4 核／8 GB 記憶體以上。
- 卡住時：`decepticon stop` → 開 Docker Desktop 的「Troubleshoot → Restart」。

### 模型回 401／403／`model not found`

- `auth: 正常` 只代表「有設定」，**不驗證金鑰是否有效**。
- 常見原因：`.env` 還是預留值（`your-anthropic-key-here`）、金鑰過期、供應商沒開該模型權限。
- 解法：`decepticon onboard --reset` 重設，或直接改 `~/.decepticon/.env`。

### 想讓代理講繁體中文

```bash
DECEPTICON_LANGUAGE=tw decepticon
```

或寫進 `~/.decepticon/.env`：`DECEPTICON_LANGUAGE=tw`。
若設了卻還是英文，請確認變數有傳進容器：`decepticon doctor` 的 `config` 項目要是 `正常`。

### WSL2 相關

- 確認 Docker Desktop 的「Resources → WSL Integration」有開啟你的發行版。
- 專案請放在 WSL 的檔案系統（`~/...`），不要放在 `/mnt/c/...`（速度與權限問題多）。

### `make` 相關目標失敗（開發者）

`make dogfood`／`make dev` 需要 `make`、`docker compose` 與網路。
Windows 上建議在 WSL2 內執行，或使用 Git Bash。

---

## 四、回報問題時請附上

1. `decepticon doctor` 的完整輸出
2. 你下的指令與完整錯誤訊息
3. 作業系統與版本（`winver`／`uname -a`）、終端機種類
4. 若與翻譯用詞有關，附上 `docs-zh/glossary.md` 中你認為該調整的詞

請開在[本漢化版倉庫的 issue](https://github.com/ssd12588/Decepticon-zh-tw/issues)。
