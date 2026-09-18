# claude-code-agentrouter-setup.ps1
#
# Windows PowerShell port of the original bash setup script.
# Installs NVM for Windows, Node.js 22, Claude Code, the VS Code
# extension, and configures AgentRouter.
#
# IMPORTANT: Run this as Administrator. Unlike Unix nvm (a per-session
# shell function), NVM for Windows switches Node versions by physically
# repointing a system-wide symlink at C:\Program Files\nodejs — that
# requires elevated privileges every time it happens.

$ErrorActionPreference = "Stop"

function Info    { param($msg) Write-Host "[INFO] $msg" }
function Success { param($msg) Write-Host "[OK]   $msg" -ForegroundColor Green }
function Warn    { param($msg) Write-Host "[WARN] $msg" -ForegroundColor Yellow }
function ErrorMsg { param($msg) Write-Host "[ERROR] $msg" -ForegroundColor Red }

Write-Host ""
Write-Host "============================================================"
Write-Host "       CLAUDE CODE COMPLETE AUTOMATED SETUP (Windows)"
Write-Host "============================================================"
Write-Host ""

# ============================================================
# 1. Confirm Administrator privileges
# ============================================================
# Replaces the original's OS-detection step — on a Windows-only script,
# that check becomes unnecessary. What actually matters is confirming
# we have the privileges NVM for Windows needs later.

$currentPrincipal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
$isAdmin = $currentPrincipal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

if (-not $isAdmin) {
    ErrorMsg "This script must be run as Administrator."
    ErrorMsg "Right-click PowerShell and choose 'Run as Administrator', then run this script again."
    exit 1
}

Success "Running with Administrator privileges."
Success "Operating system: Windows"

# ============================================================
# 2. Persistent environment variables
# ============================================================
# Bash persists variables by appending 'export' lines to ~/.bashrc or
# ~/.zshrc, then re-sourcing that file in every new shell. Windows has
# no equivalent "shell config file" — instead, environment variables
# set at the User scope are stored directly by Windows itself and are
# automatically available in every future PowerShell or CMD session,
# with no separate config file to maintain or re-source.

function Set-PersistentEnvVar {
    param(
        [string]$Name,
        [string]$Value
    )
    [System.Environment]::SetEnvironmentVariable($Name, $Value, "User")
    # Also set for the rest of *this* script's session, since a variable
    # set at User scope doesn't apply to the already-running process.
    Set-Item -Path "Env:$Name" -Value $Value
}

Success "Environment variables will persist via Windows User settings (no config file needed)."

# ============================================================
# 3. Check curl
# ============================================================
# Windows 10 1809+ and Windows 11 both include curl.exe natively —
# the same minimum OS version Claude Code itself requires.

if (-not (Get-Command curl.exe -ErrorAction SilentlyContinue)) {
    ErrorMsg "curl was not found. It ships with Windows 10 1809+ and Windows 11."
    ErrorMsg "If it's missing, your Windows version may be too old for Claude Code."
    exit 1
}

Success "curl is available."

# ============================================================
# 4. Install / load NVM for Windows
# ============================================================
# A genuinely different tool from Unix nvm, not a renamed port of it.
# Installed via winget — built into Windows 10 1809+ and Windows 11 —
# rather than a curl-piped install script, since winget is the current,
# maintained distribution channel for it.

if (Get-Command nvm -ErrorAction SilentlyContinue) {
    Info "NVM for Windows is already installed."
} else {
    Info "NVM for Windows was not found."
    Info "Installing NVM for Windows via winget..."

    winget install --id CoreyButler.NVMforWindows --exact --silent

    # winget installs to a fresh PATH entry; refresh this session's PATH
    # from the registry so 'nvm' is recognized without reopening the terminal.
    $env:Path = [System.Environment]::GetEnvironmentVariable("Path", "Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path", "User")

    if (-not (Get-Command nvm -ErrorAction SilentlyContinue)) {
        ErrorMsg "NVM for Windows installation failed, or requires a new terminal window to be detected."
        ErrorMsg "Close this window, reopen PowerShell as Administrator, and re-run this script."
        exit 1
    }
}

Success "NVM for Windows is ready."

# ============================================================
# 5. Install Node.js 22
# ============================================================
# 'nvm install <version>' and 'nvm use <version>' share the same syntax
# as Unix nvm. There's no 'nvm alias default' step here, though — NVM
# for Windows has no separate "default" concept distinct from "currently
# active," since 'nvm use' already changes the system-wide symlink
# directly. Whatever you last activated stays active for every future
# session on its own, with nothing else to set.

$nodeRequiredMajor = 22

$nodeCmd = Get-Command node -ErrorAction SilentlyContinue

if ($nodeCmd) {
    $nodeVersion = node --version
    $nodeMajor = [int]($nodeVersion -replace '^v(\d+)\..*', '$1')

    Info "Node.js detected: $nodeVersion"

    if ($nodeMajor -ge $nodeRequiredMajor) {
        Success "Node.js version is compatible."
    } else {
        Warn "Current Node.js version is older than 22."
        Info "Installing Node.js 22..."
        nvm install 22
        nvm use 22
    }
} else {
    Info "Node.js was not found."
    Info "Installing Node.js 22..."
    nvm install 22
    nvm use 22
}

# Refresh PATH again so 'node'/'npm' resolve to the version nvm just activated.
$env:Path = [System.Environment]::GetEnvironmentVariable("Path", "Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path", "User")

Success "Node.js: $(node --version)"
Success "npm: $(npm --version)"

# ============================================================
# 6. Install Claude Code CLI
# ============================================================
# Identical command to the original — npm's global install works the
# same way on Windows.

if (Get-Command claude -ErrorAction SilentlyContinue) {
    Success "Claude Code is already installed."
    Info "Claude location: $((Get-Command claude).Source)"
} else {
    Info "Claude Code is not installed."
    Info "Installing Claude Code..."

    npm install -g "@anthropic-ai/claude-code"

    if (-not (Get-Command claude -ErrorAction SilentlyContinue)) {
        ErrorMsg "Claude Code installation failed."
        exit 1
    }

    Success "Claude Code installed."
}

# ============================================================
# 7. Detect VS Code
# ============================================================

Write-Host ""
Write-Host "============================================================"
Write-Host "                 VS CODE INTEGRATION"
Write-Host "============================================================"
Write-Host ""

$vsCodeAvailable = $false

if (Get-Command code -ErrorAction SilentlyContinue) {
    Success "VS Code command detected."
    Info "VS Code command: $((Get-Command code).Source)"
    $vsCodeAvailable = $true
} else {
    Warn "VS Code command 'code' was not found."
    Warn "Claude Code VS Code extension will be skipped."
    Warn "Install/open VS Code and make sure the 'code' command"
    Warn "is available in your PATH if you want the extension."
}

# ============================================================
# 8. Install Claude Code VS Code extension
# ============================================================

if ($vsCodeAvailable) {
    $claudeExtensionId = "anthropic.claude-code"

    Info "Checking Claude Code VS Code extension..."

    $installedExtensions = code --list-extensions 2>$null

    if ($installedExtensions -contains $claudeExtensionId) {
        Success "Claude Code VS Code extension is already installed."
    } else {
        Info "Installing Claude Code VS Code extension..."

        code --install-extension $claudeExtensionId

        if ($LASTEXITCODE -eq 0) {
            Success "Claude Code VS Code extension installed."
        } else {
            Warn "VS Code extension installation failed."
            Warn "The CLI installation will continue."
        }
    }
}

# ============================================================
# 9. AgentRouter configuration
# ============================================================

Write-Host ""
Write-Host "============================================================"
Write-Host "                AGENTROUTER CONFIGURATION"
Write-Host "============================================================"
Write-Host ""

Write-Host "This script does NOT install AgentRouter."
Write-Host ""
Write-Host "It configures Claude Code to use your AgentRouter backend."
Write-Host ""
Write-Host "You will provide:"
Write-Host ""
Write-Host "  1. AgentRouter Base URL"
Write-Host "  2. AgentRouter Auth Token"
Write-Host "  3. AgentRouter Model"
Write-Host ""

# ============================================================
# 10. Read existing configuration
# ============================================================

$existingBaseUrl = $env:ANTHROPIC_BASE_URL
$existingToken = $env:ANTHROPIC_AUTH_TOKEN
$existingModel = $env:ANTHROPIC_MODEL

# ============================================================
# 11. Ask for Base URL
# ============================================================

if ($existingBaseUrl) {
    Write-Host "Existing ANTHROPIC_BASE_URL:"
    Write-Host $existingBaseUrl
    Write-Host ""
    $agentRouterBaseUrl = Read-Host "Press ENTER to keep it, or enter a new Base URL"
    if (-not $agentRouterBaseUrl) { $agentRouterBaseUrl = $existingBaseUrl }
} else {
    $agentRouterBaseUrl = Read-Host "AgentRouter Base URL"
}

if (-not $agentRouterBaseUrl) {
    ErrorMsg "AgentRouter Base URL cannot be empty."
    exit 1
}

# ============================================================
# 12. Ask for Auth Token
# ============================================================
# Read-Host -AsSecureString masks input the same way bash's 'read -rsp'
# does. It's converted back to plain text below since it needs to be
# written to settings.json and the environment as normal text.

if ($existingToken) {
    Write-Host ""
    Write-Host "An existing ANTHROPIC_AUTH_TOKEN was detected."
    Write-Host ""
    $secureToken = Read-Host "Press ENTER to keep it, or enter a new Auth Token" -AsSecureString
    $agentRouterAuthToken = [System.Runtime.InteropServices.Marshal]::PtrToStringAuto([System.Runtime.InteropServices.Marshal]::SecureStringToBSTR($secureToken))
    if (-not $agentRouterAuthToken) { $agentRouterAuthToken = $existingToken }
} else {
    Write-Host ""
    $secureToken = Read-Host "AgentRouter Auth Token" -AsSecureString
    $agentRouterAuthToken = [System.Runtime.InteropServices.Marshal]::PtrToStringAuto([System.Runtime.InteropServices.Marshal]::SecureStringToBSTR($secureToken))
}

if (-not $agentRouterAuthToken) {
    ErrorMsg "AgentRouter Auth Token cannot be empty."
    exit 1
}

# ============================================================
# 13. Ask for Model
# ============================================================

if ($existingModel) {
    Write-Host ""
    Write-Host "Existing ANTHROPIC_MODEL:"
    Write-Host $existingModel
    Write-Host ""
    $agentRouterModel = Read-Host "Press ENTER to keep it, or enter a new Model"
    if (-not $agentRouterModel) { $agentRouterModel = $existingModel }
} else {
    Write-Host ""
    $agentRouterModel = Read-Host "AgentRouter Model"
}

if (-not $agentRouterModel) {
    ErrorMsg "AgentRouter Model cannot be empty."
    exit 1
}

# ============================================================
# 14. Save AgentRouter variables persistently
# ============================================================

Info "Saving AgentRouter variables to your Windows User environment..."

Set-PersistentEnvVar -Name "ANTHROPIC_BASE_URL" -Value $agentRouterBaseUrl
Set-PersistentEnvVar -Name "ANTHROPIC_AUTH_TOKEN" -Value $agentRouterAuthToken
Set-PersistentEnvVar -Name "ANTHROPIC_MODEL" -Value $agentRouterModel

Success "AgentRouter variables saved."

# ============================================================
# 15. Create %USERPROFILE%\.claude directory
# ============================================================

$claudeDir = Join-Path $env:USERPROFILE ".claude"
$claudeSettings = Join-Path $claudeDir "settings.json"

New-Item -ItemType Directory -Path $claudeDir -Force | Out-Null

Success "Claude directory: $claudeDir"

# ============================================================
# 16. Configure settings.json
# ============================================================
# The original's Node.js heredoc is genuine JavaScript, not bash — it
# runs identically here. Only how it's *invoked* changes: PowerShell's
# here-string (@'...'@) replaces bash's <<'NODE' syntax, and values are
# passed as command-line arguments instead of environment variables,
# since that's simpler to do correctly from PowerShell.

Info "Configuring Claude Code settings..."

if (Test-Path $claudeSettings) {
    $backupFile = "$claudeSettings.backup.$(Get-Date -Format 'yyyyMMdd_HHmmss')"
    Copy-Item -Path $claudeSettings -Destination $backupFile
    Success "Existing settings backed up:"
    Write-Host "       $backupFile"
}

$mergeScript = @'
const fs = require("fs");

const settingsFile = process.argv[2];
const baseUrl = process.argv[3];
const authToken = process.argv[4];
const model = process.argv[5];

let settings = {};

if (fs.existsSync(settingsFile)) {
    try {
        const content = fs.readFileSync(settingsFile, "utf8").trim();
        if (content.length > 0) {
            settings = JSON.parse(content);
        }
    } catch (error) {
        console.error("Existing settings.json is not valid JSON.");
        console.error("A backup was already created by the setup script.");
        console.error("Creating a fresh settings.json.");
        settings = {};
    }
}

if (!settings || typeof settings !== "object" || Array.isArray(settings)) {
    settings = {};
}

if (!settings.env || typeof settings.env !== "object" || Array.isArray(settings.env)) {
    settings.env = {};
}

settings.env.ANTHROPIC_BASE_URL = baseUrl;
settings.env.ANTHROPIC_AUTH_TOKEN = authToken;
settings.env.ANTHROPIC_MODEL = model;

settings.model = model;

settings["$schema"] = "https://json.schemastore.org/claude-code-settings.json";

fs.writeFileSync(settingsFile, JSON.stringify(settings, null, 2) + "\n");
'@

$tempScriptPath = Join-Path $env:TEMP "claude-settings-merge.js"
Set-Content -Path $tempScriptPath -Value $mergeScript

node $tempScriptPath $claudeSettings $agentRouterBaseUrl $agentRouterAuthToken $agentRouterModel

Remove-Item -Path $tempScriptPath -Force

# Restrict the settings file to the current user only — the closest
# Windows equivalent to the original's chmod 0o600 (owner read/write only).
$acl = Get-Acl $claudeSettings
$acl.SetAccessRuleProtection($true, $false)
$rule = New-Object System.Security.AccessControl.FileSystemAccessRule($env:USERNAME, "FullControl", "Allow")
$acl.AddAccessRule($rule)
Set-Acl -Path $claudeSettings -AclObject $acl

Success "Claude settings configured:"
Write-Host "       $claudeSettings"

# ============================================================
# 17. Verify settings.json
# ============================================================

if (Test-Path $claudeSettings) {
    try {
        Get-Content $claudeSettings -Raw | ConvertFrom-Json | Out-Null
        Success "settings.json is valid JSON."
    } catch {
        ErrorMsg "settings.json validation failed."
        exit 1
    }
} else {
    ErrorMsg "settings.json was not created."
    exit 1
}

# ============================================================
# 18. Verify Claude Code
# ============================================================

Write-Host ""
Write-Host "============================================================"
Write-Host "                    VERIFICATION"
Write-Host "============================================================"
Write-Host ""

Success "Operating system : Windows"
Success "Node.js          : $(node --version)"
Success "npm              : $(npm --version)"

if (Get-Command claude -ErrorAction SilentlyContinue) {
    Success "Claude Code      : installed"
    Info "Claude path      : $((Get-Command claude).Source)"
} else {
    ErrorMsg "Claude Code      : NOT FOUND"
    exit 1
}

if ($vsCodeAvailable) {
    $installedExtensions = code --list-extensions 2>$null
    if ($installedExtensions -contains "anthropic.claude-code") {
        Success "Claude VS Code extension : installed"
    } else {
        Warn "Claude VS Code extension : not detected"
    }
} else {
    Warn "Claude VS Code extension : skipped (VS Code not detected)"
}

# ============================================================
# 19. Verify AgentRouter settings
# ============================================================

Write-Host ""
Write-Host "AgentRouter configuration:"
Write-Host ""
Write-Host "  Base URL : $env:ANTHROPIC_BASE_URL"
Write-Host "  Model    : $env:ANTHROPIC_MODEL"
Write-Host "  Token    : configured"
Write-Host ""

# ============================================================
# 20. Final message
# ============================================================

Write-Host ""
Write-Host "============================================================"
Write-Host "                 SETUP COMPLETE"
Write-Host "============================================================"
Write-Host ""
Write-Host "Claude Code CLI:"
Write-Host ""
Write-Host "  claude"
Write-Host ""
Write-Host "Claude settings:"
Write-Host ""
Write-Host "  $claudeSettings"
Write-Host ""
Write-Host "IMPORTANT:"
Write-Host ""
Write-Host "Close this terminal and open a new one so your environment"
Write-Host "variables load correctly. Unlike bash's 'source', Windows"
Write-Host "environment variables need a fresh session to take effect."
Write-Host ""
Write-Host "Then verify Claude:"
Write-Host ""
Write-Host "  claude --version"
Write-Host ""
Write-Host "Then start Claude Code:"
Write-Host ""
Write-Host "  claude"
Write-Host ""
Write-Host "============================================================"
