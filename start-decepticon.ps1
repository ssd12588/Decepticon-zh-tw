<#
.SYNOPSIS
  Decepticon 繁體中文漢化版 · Windows 一鍵啟動

.DESCRIPTION
  幫你完成：檢查環境 → 建置漢化版 → 初始化設定 → 啟動 Decepticon。
  設定會記住（%USERPROFILE%\.decepticon\zh-launcher.json），下次不必重輸。

  官方安裝腳本抓的是上游英文版；本腳本從原始碼建置「漢化版」，
  所以畫面上的 CLI、TUI、診斷訊息都會是繁體中文。

.PARAMETER Check
  只檢查環境與建置狀態，不下載、不啟動（不耗任何 API 額度）。

.PARAMETER ResetConfig
  清除先前記住的啟動設定（語言、模型選擇）。

.PARAMETER NoCli
  只啟動服務堆疊，不進入對話介面（適合先確認服務都起來）。

.PARAMETER DecepticonArgs
  直接透傳給 decepticon 的參數，例如 --continue、--resume、--no-mouse。

.EXAMPLE
  .\start-decepticon.ps1 -Check

.EXAMPLE
  .\start-decepticon.ps1

.EXAMPLE
  .\start-decepticon.ps1 -NoCli

.EXAMPLE
  .\start-decepticon.ps1 -DecepticonArgs '--continue'

.NOTES
  建議使用 Windows Terminal + PowerShell 7（pwsh），中文顯示較穩定。
  若無法執行腳本，先執行：
    Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
#>

[CmdletBinding()]
param(
    [switch]$Check,
    [switch]$ResetConfig,
    [switch]$NoCli,
    [Parameter(ValueFromRemainingArguments = $true)]
    [string[]]$DecepticonArgs
)

$ErrorActionPreference = 'Stop'

# ---------- 0. 終端編碼（中文顯示） ----------
try { chcp 65001 | Out-Null } catch { }
try {
    $OutputEncoding = [System.Text.Encoding]::UTF8
    [Console]::OutputEncoding = [System.Text.Encoding]::UTF8
} catch { }

function Write-Step($msg) { Write-Host "`n▶ $msg" -ForegroundColor Cyan }
function Write-Ok($msg)   { Write-Host "  [OK] $msg" -ForegroundColor Green }
function Write-Warn2($msg){ Write-Host "  [!]  $msg" -ForegroundColor Yellow }
function Write-Err($msg)  { Write-Host "  [X]  $msg" -ForegroundColor Red }
function Write-Hint($msg) { Write-Host "       $msg" -ForegroundColor Gray }

$repoRoot = Split-Path -Parent $PSCommandPath
$launcherDir = Join-Path $repoRoot 'clients\launcher'
$cliDir = Join-Path $repoRoot 'clients\cli'
$launcherExe = Join-Path $launcherDir 'bin\decepticon.exe'
$configDir = Join-Path $env:USERPROFILE '.decepticon'
$configPath = Join-Path $configDir 'zh-launcher.json'

if (-not (Test-Path (Join-Path $repoRoot 'docker-compose.yml'))) {
    Write-Err '這不是 Decepticon 專案目錄'
    Write-Hint "請把腳本放在 repo 根目錄（目前是：$repoRoot）"
    exit 1
}

# ---------- 1. Docker ----------
Write-Step '檢查 Docker'
$dockerOk = $false
if (Get-Command docker -ErrorAction SilentlyContinue) {
    $null = docker info 2>$null
    if ($LASTEXITCODE -eq 0) {
        $dockerOk = $true
        Write-Ok 'Docker 正在執行'
    } else {
        Write-Err 'Docker 已安裝，但沒有在執行'
        Write-Hint '請開啟 Docker Desktop，等狀態變成 Engine running 再試。'
    }
} else {
    Write-Err '找不到 docker 指令'
    Write-Hint '請安裝 Docker Desktop：https://www.docker.com/products/docker-desktop/'
    Write-Hint '安裝後請「完全關閉終端機再重開」，PATH 才會生效。'
}
$composeOk = $false
if ($dockerOk) {
    $null = docker compose version 2>$null
    $composeOk = ($LASTEXITCODE -eq 0)
    if ($composeOk) { Write-Ok 'Docker Compose v2 可用' }
    else {
        Write-Err '缺少 Docker Compose v2'
        Write-Hint 'Docker Desktop 內建；若使用 Docker Engine，請安裝 compose 外掛。'
    }
}

# ---------- 2. 建置工具鏈 ----------
Write-Step '檢查建置工具'
$goOk = [bool](Get-Command go -ErrorAction SilentlyContinue)
$nodeOk = [bool](Get-Command npm -ErrorAction SilentlyContinue)
if ($goOk) { Write-Ok ("Go 可用（" + (go version) + '）') }
else {
    Write-Warn2 '找不到 Go（用來編譯啟動器）'
    Write-Hint '安裝：winget install GoLang.Go'
}
if ($nodeOk) { Write-Ok ("Node 可用（npm " + (npm --version) + '）') }
else {
    Write-Warn2 '找不到 npm（用來建置對話介面）'
    Write-Hint '安裝：winget install OpenJS.NodeJS.LTS'
}

# ---------- 3. 建置啟動器 ----------
Write-Step '建置 Decepticon 啟動器（漢化版）'
$needLauncher = $true
if (Test-Path $launcherExe) {
    $exeTime = (Get-Item $launcherExe).LastWriteTimeUtc
    $newer = Get-ChildItem $launcherDir -Recurse -File -Include *.go | Where-Object { $_.LastWriteTimeUtc -gt $exeTime }
    if (-not $newer) {
        $needLauncher = $false
        Write-Ok '已有最新建置（原始碼沒有更新）'
    }
}
if ($needLauncher) {
    if (-not $goOk) {
        Write-Err '需要 Go 才能建置啟動器'
        Write-Hint '執行 winget install GoLang.Go，重開終端機後再跑本腳本。'
        exit 1
    }
    Push-Location $launcherDir
    try {
        New-Item -ItemType Directory -Force -Path (Join-Path $launcherDir 'bin') | Out-Null
        Write-Hint '第一次建置會下載 Go 模組，請稍等…'
        go build -ldflags '-X github.com/BitterSecurity/Decepticon/clients/launcher/cmd.version=dev' -o bin\decepticon.exe .
        if ($LASTEXITCODE -ne 0) { throw 'go build 失敗' }
        Write-Ok '啟動器建置完成（bin\decepticon.exe）'
    } catch {
        Write-Err "建置失敗：$($_.Exception.Message)"
        Write-Hint '請把上面的錯誤訊息完整貼出來回報。'
        Pop-Location; exit 1
    } finally { Pop-Location }
}

# ---------- 4. 建置對話介面（Ink） ----------
Write-Step '建置對話介面（Ink CLI）'
$cliEntry = Join-Path $cliDir 'dist\index.js'
$needCli = $false
if (-not $nodeOk) {
    Write-Warn2 '沒有 npm，略過對話介面建置'
    Write-Hint '啟動器仍可使用（--help、doctor、status），但無法進入 AI 對話。'
} elseif (-not (Test-Path $cliEntry)) {
    $needCli = $true
} else {
    $srcNewer = Get-ChildItem $cliDir -Recurse -File -Include *.ts,*.tsx |
        Where-Object { $_.FullName -notlike '*node_modules*' -and $_.LastWriteTimeUtc -gt (Get-Item $cliEntry).LastWriteTimeUtc }
    if ($srcNewer) {
        $needCli = $true
        Write-Hint '偵測到介面原始碼較新，重新建置…'
    } else {
        Write-Ok '對話介面已是最新建置'
    }
}
if ($nodeOk -and $needCli) {
    Push-Location $repoRoot
    try {
        if (-not (Test-Path (Join-Path $repoRoot 'node_modules'))) {
            Write-Hint '安裝 Node 相依套件（npm install）…'
            npm install --no-audit --no-fund
            if ($LASTEXITCODE -ne 0) { throw 'npm install 失敗' }
        }
        Write-Hint '建置共用串流套件…'
        npm run build --workspace=@decepticon/streaming | Out-Null
        Push-Location $cliDir
        try {
            npm run build | Out-Null
            if ($LASTEXITCODE -ne 0) { throw '對話介面建置失敗' }
            Write-Ok '對話介面建置完成（clients\cli\dist）'
        } finally { Pop-Location }
    } catch {
        Write-Err "建置失敗：$($_.Exception.Message)"
        Write-Hint '可先只驗證啟動器：.\start-decepticon.ps1 -NoCli'
        Pop-Location; exit 1
    } finally { Pop-Location }
}

# ---------- 5. 記住的設定 ----------
Write-Step '確認設定'
if ($ResetConfig -and (Test-Path $configPath)) {
    Remove-Item $configPath -Force
    Write-Warn2 '已清除先前的啟動設定'
}
$saved = $null
if (Test-Path $configPath) {
    try { $saved = Get-Content $configPath -Raw -Encoding UTF8 | ConvertFrom-Json } catch { }
}
$language = $env:DECEPTICON_LANGUAGE
if (-not $language -and $saved) { $language = $saved.DECEPTICON_LANGUAGE }
if (-not $language) { $language = 'tw' }

try {
    if (-not (Test-Path $configDir)) { New-Item -ItemType Directory -Path $configDir -Force | Out-Null }
    ([ordered]@{ DECEPTICON_LANGUAGE = $language }) | ConvertTo-Json | Set-Content -Path $configPath -Encoding UTF8
} catch {
    Write-Warn2 "無法儲存設定（不影響本次執行）：$($_.Exception.Message)"
}
$env:DECEPTICON_LANGUAGE = $language
Write-Ok "代理輸出語言：$language（tw = 繁體中文）"

$envPath = Join-Path $configDir '.env'
if (-not (Test-Path $envPath)) {
    Write-Warn2 '尚未完成初始化設定（找不到 ~\.decepticon\.env）'
    Write-Hint '稍後會自動執行 decepticon onboard，請依畫面輸入 API 金鑰與模型設定檔。'
} else {
    Write-Ok '已找到既有設定（~\.decepticon\.env）'
}

# ---------- 6. 只有檢查模式（純唯讀，不初始化、不啟動） ----------
if ($Check) {
    Write-Step '執行唯讀診斷'
    & $launcherExe doctor
    Write-Host ''
    if (-not ($dockerOk -and $composeOk)) {
        Write-Err '環境尚未就緒，請先處理上面標示 [X] 的項目'
        exit 1
    }
    if (-not (Test-Path $envPath)) {
        Write-Warn2 'Docker 已就緒，但還沒完成初始化設定'
        Write-Hint '下一步：直接執行 .\start-decepticon.ps1（會引導你輸入 API 金鑰）'
    } else {
        Write-Ok '環境檢查完成，可以正式啟動'
        Write-Hint '下一步：.\start-decepticon.ps1'
    }
    exit 0
}

if (-not ($dockerOk -and $composeOk)) {
    Write-Err 'Docker 沒有就緒，無法啟動'
    Write-Hint '開啟 Docker Desktop，等 Engine running 後再跑一次。'
    exit 1
}

# ---------- 6.5 首次初始化設定 ----------
if (-not (Test-Path $envPath)) {
    Write-Step '初始化設定（decepticon onboard）'
    Write-Hint '請依畫面選擇驗證方式、供應商、API 金鑰與模型設定檔。'
    Write-Host ''
    & $launcherExe onboard
    if ($LASTEXITCODE -ne 0) {
        Write-Err '初始化沒有完成'
        Write-Hint '可重跑：.\clients\launcher\bin\decepticon.exe onboard --reset'
        exit 1
    }
    Write-Ok '初始化完成（~\.decepticon\.env）'
}

# ---------- 7. 啟動 ----------
Write-Step '啟動 Decepticon'
Write-Host '  首次啟動會下載數 GB 映像，請耐心等待…' -ForegroundColor Gray
Write-Host '  按 Ctrl+C 可中止；服務會繼續在背景執行，下次啟動更快。' -ForegroundColor Gray
Write-Host ''

$allArgs = @()
if ($NoCli) {
    # --background：只啟動服務並備妥授權測試任務，不進入互動式 CLI
    $allArgs += @('start', '--background')
}
if ($DecepticonArgs) { $allArgs += $DecepticonArgs }

& $launcherExe @allArgs
$code = $LASTEXITCODE

Write-Host ''
switch ($code) {
    0 { Write-Ok '結束（正常）' }
    1 { Write-Err '發生錯誤｜請看 docs-zh\TROUBLESHOOTING.md' }
    default { Write-Hint "結束碼：$code" }
}
Write-Host '  常用指令：' -ForegroundColor Gray
Write-Host '    .\start-decepticon.ps1 -Check        環境檢查（不耗額度）' -ForegroundColor White
Write-Host '    .\clients\launcher\bin\decepticon.exe status' -ForegroundColor White
Write-Host '    .\clients\launcher\bin\decepticon.exe stop' -ForegroundColor White
exit $code
