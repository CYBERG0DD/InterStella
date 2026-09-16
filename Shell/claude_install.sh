#!/usr/bin/env bash

set -e

# ============================================================
# Claude Code + AgentRouter + VS Code Complete Setup
# ============================================================

info() {
    echo "[INFO] $1"
}

success() {
    echo "[OK]   $1"
}

warn() {
    echo "[WARN] $1"
}

error() {
    echo "[ERROR] $1"
}

echo
echo "============================================================"
echo "       CLAUDE CODE COMPLETE AUTOMATED SETUP"
echo "============================================================"
echo

# ============================================================
# 1. Detect operating system
# ============================================================

OS="$(uname -s)"

case "$OS" in
    Linux*)
        OS_NAME="Linux"
        ;;
    Darwin*)
        OS_NAME="macOS"
        ;;
    *)
        error "Unsupported operating system: $OS"
        exit 1
        ;;
esac

success "Operating system: $OS_NAME"

# ============================================================
# 2. Detect shell
# ============================================================

CURRENT_SHELL="$(basename "${SHELL:-}")"

case "$CURRENT_SHELL" in
    zsh)
        SHELL_CONFIG="$HOME/.zshrc"
        ;;
    bash)
        SHELL_CONFIG="$HOME/.bashrc"
        ;;
    *)
        error "Unsupported shell: $CURRENT_SHELL"
        error "Supported shells: bash and zsh"
        exit 1
        ;;
esac

touch "$SHELL_CONFIG"

success "Shell: $CURRENT_SHELL"
success "Shell configuration: $SHELL_CONFIG"

# ============================================================
# 3. Helper to add a line without duplicates
# ============================================================

add_shell_line() {
    local line="$1"

    if ! grep -Fqx "$line" "$SHELL_CONFIG" 2>/dev/null; then
        printf '%s\n' "$line" >> "$SHELL_CONFIG"
    fi
}

# ============================================================
# 4. Helper to set an exported shell variable
# ============================================================

set_shell_variable() {
    local variable="$1"
    local value="$2"

    # Remove existing export for this variable.
    sed -i.bak \
        "/^[[:space:]]*export[[:space:]]\+${variable}=.*/d" \
        "$SHELL_CONFIG"

    rm -f "${SHELL_CONFIG}.bak"

    # Add the new value safely.
    printf 'export %s=%q\n' "$variable" "$value" >> "$SHELL_CONFIG"
}

# ============================================================
# 5. Check curl
# ============================================================

if ! command -v curl >/dev/null 2>&1; then
    error "curl is required but was not found."
    error "Please ask your system administrator to install curl."
    exit 1
fi

success "curl is available."

# ============================================================
# 6. Install / load NVM
# ============================================================

export NVM_DIR="${NVM_DIR:-$HOME/.nvm}"

if [ -s "$NVM_DIR/nvm.sh" ]; then

    info "NVM is already installed."

else

    info "NVM was not found."
    info "Installing NVM..."

    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash

    export NVM_DIR="$HOME/.nvm"

    if [ ! -s "$NVM_DIR/nvm.sh" ]; then
        error "NVM installation failed."
        exit 1
    fi

fi

source "$NVM_DIR/nvm.sh"

success "NVM is ready."

# ============================================================
# 7. Configure NVM permanently
# ============================================================

add_shell_line 'export NVM_DIR="$HOME/.nvm"'
add_shell_line '[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"'
add_shell_line '[ -s "$NVM_DIR/bash_completion" ] && . "$NVM_DIR/bash_completion"'

# ============================================================
# 8. Install Node.js 22
# ============================================================

NODE_REQUIRED_MAJOR=22

if command -v node >/dev/null 2>&1; then

    NODE_VERSION="$(node --version)"
    NODE_MAJOR="$(node --version | sed 's/^v//' | cut -d. -f1)"

    info "Node.js detected: $NODE_VERSION"

    if [ "$NODE_MAJOR" -ge "$NODE_REQUIRED_MAJOR" ]; then

        success "Node.js version is compatible."

    else

        warn "Current Node.js version is older than 22."
        info "Installing Node.js 22..."

        nvm install 22
        nvm alias default 22
        nvm use 22

    fi

else

    info "Node.js was not found."
    info "Installing Node.js 22..."

    nvm install 22
    nvm alias default 22
    nvm use 22

fi

success "Node.js: $(node --version)"
success "npm: $(npm --version)"

# ============================================================
# 9. Install Claude Code CLI
# ============================================================

if command -v claude >/dev/null 2>&1; then

    success "Claude Code is already installed."
    info "Claude location: $(command -v claude)"

else

    info "Claude Code is not installed."
    info "Installing Claude Code..."

    npm install -g @anthropic-ai/claude-code

    if ! command -v claude >/dev/null 2>&1; then
        error "Claude Code installation failed."
        exit 1
    fi

    success "Claude Code installed."

fi

# ============================================================
# 10. Detect VS Code
# ============================================================

echo
echo "============================================================"
echo "                 VS CODE INTEGRATION"
echo "============================================================"
echo

if command -v code >/dev/null 2>&1; then

    success "VS Code command detected."
    info "VS Code command: $(command -v code)"

    VS_CODE_AVAILABLE="true"

else

    warn "VS Code command 'code' was not found."
    warn "Claude Code VS Code extension will be skipped."
    warn "Install/open VS Code and make sure the 'code' command"
    warn "is available in your PATH if you want the extension."

    VS_CODE_AVAILABLE="false"

fi

# ============================================================
# 11. Install Claude Code VS Code extension
# ============================================================

if [ "$VS_CODE_AVAILABLE" = "true" ]; then

    CLAUDE_EXTENSION_ID="anthropic.claude-code"

    info "Checking Claude Code VS Code extension..."

    if code --list-extensions 2>/dev/null | \
        grep -Fxq "$CLAUDE_EXTENSION_ID"; then

        success "Claude Code VS Code extension is already installed."

    else

        info "Installing Claude Code VS Code extension..."

        if code --install-extension "$CLAUDE_EXTENSION_ID"; then

            success "Claude Code VS Code extension installed."

        else

            warn "VS Code extension installation failed."
            warn "The CLI installation will continue."

        fi

    fi

fi

# ============================================================
# 12. AgentRouter configuration
# ============================================================

echo
echo "============================================================"
echo "                AGENTROUTER CONFIGURATION"
echo "============================================================"
echo

echo "This script does NOT install AgentRouter."
echo
echo "It configures Claude Code to use your AgentRouter backend."
echo
echo "You will provide:"
echo
echo "  1. AgentRouter Base URL"
echo "  2. AgentRouter Auth Token"
echo "  3. AgentRouter Model"
echo

# ============================================================
# 13. Read existing configuration
# ============================================================

EXISTING_BASE_URL="${ANTHROPIC_BASE_URL:-}"
EXISTING_TOKEN="${ANTHROPIC_AUTH_TOKEN:-}"
EXISTING_MODEL="${ANTHROPIC_MODEL:-}"

# ============================================================
# 14. Ask for Base URL
# ============================================================

if [ -n "$EXISTING_BASE_URL" ]; then

    echo "Existing ANTHROPIC_BASE_URL:"
    echo "$EXISTING_BASE_URL"
    echo

    read -rp "Press ENTER to keep it, or enter a new Base URL: " AGENTROUTER_BASE_URL

    if [ -z "$AGENTROUTER_BASE_URL" ]; then
        AGENTROUTER_BASE_URL="$EXISTING_BASE_URL"
    fi

else

    read -rp "AgentRouter Base URL: " AGENTROUTER_BASE_URL

fi

if [ -z "$AGENTROUTER_BASE_URL" ]; then
    error "AgentRouter Base URL cannot be empty."
    exit 1
fi

# ============================================================
# 15. Ask for Auth Token
# ============================================================

if [ -n "$EXISTING_TOKEN" ]; then

    echo
    echo "An existing ANTHROPIC_AUTH_TOKEN was detected."
    echo

    read -rsp "Press ENTER to keep it, or enter a new Auth Token: " AGENTROUTER_AUTH_TOKEN
    echo

    if [ -z "$AGENTROUTER_AUTH_TOKEN" ]; then
        AGENTROUTER_AUTH_TOKEN="$EXISTING_TOKEN"
    fi

else

    echo
    read -rsp "AgentRouter Auth Token: " AGENTROUTER_AUTH_TOKEN
    echo

fi

if [ -z "$AGENTROUTER_AUTH_TOKEN" ]; then
    error "AgentRouter Auth Token cannot be empty."
    exit 1
fi

# ============================================================
# 16. Ask for Model
# ============================================================

if [ -n "$EXISTING_MODEL" ]; then

    echo
    echo "Existing ANTHROPIC_MODEL:"
    echo "$EXISTING_MODEL"
    echo

    read -rp "Press ENTER to keep it, or enter a new Model: " AGENTROUTER_MODEL

    if [ -z "$AGENTROUTER_MODEL" ]; then
        AGENTROUTER_MODEL="$EXISTING_MODEL"
    fi

else

    echo
    read -rp "AgentRouter Model: " AGENTROUTER_MODEL

fi

if [ -z "$AGENTROUTER_MODEL" ]; then
    error "AgentRouter Model cannot be empty."
    exit 1
fi

# ============================================================
# 17. Save AgentRouter variables to shell config
# ============================================================

info "Saving AgentRouter variables to $SHELL_CONFIG..."

set_shell_variable \
    "ANTHROPIC_BASE_URL" \
    "$AGENTROUTER_BASE_URL"

set_shell_variable \
    "ANTHROPIC_AUTH_TOKEN" \
    "$AGENTROUTER_AUTH_TOKEN"

set_shell_variable \
    "ANTHROPIC_MODEL" \
    "$AGENTROUTER_MODEL"

success "AgentRouter variables saved."

# ============================================================
# 18. Export variables for current script
# ============================================================

export ANTHROPIC_BASE_URL="$AGENTROUTER_BASE_URL"
export ANTHROPIC_AUTH_TOKEN="$AGENTROUTER_AUTH_TOKEN"
export ANTHROPIC_MODEL="$AGENTROUTER_MODEL"

# ============================================================
# 19. Create ~/.claude directory
# ============================================================

CLAUDE_DIR="$HOME/.claude"
CLAUDE_SETTINGS="$CLAUDE_DIR/settings.json"

mkdir -p "$CLAUDE_DIR"

success "Claude directory: $CLAUDE_DIR"

# ============================================================
# 20. Configure ~/.claude/settings.json
# ============================================================

info "Configuring Claude Code settings..."

# Make a backup if settings already exists.
if [ -f "$CLAUDE_SETTINGS" ]; then

    BACKUP_FILE="${CLAUDE_SETTINGS}.backup.$(date +%Y%m%d_%H%M%S)"

    cp "$CLAUDE_SETTINGS" "$BACKUP_FILE"

    success "Existing settings backed up:"
    echo "       $BACKUP_FILE"

fi

# ------------------------------------------------------------
# Use Node.js to safely merge settings.
# Existing settings are preserved.
# ------------------------------------------------------------

export SETUP_CLAUDE_SETTINGS="$CLAUDE_SETTINGS"
export SETUP_ANTHROPIC_BASE_URL="$AGENTROUTER_BASE_URL"
export SETUP_ANTHROPIC_AUTH_TOKEN="$AGENTROUTER_AUTH_TOKEN"
export SETUP_ANTHROPIC_MODEL="$AGENTROUTER_MODEL"

node <<'NODE'
const fs = require("fs");

const settingsFile = process.env.SETUP_CLAUDE_SETTINGS;
const baseUrl = process.env.SETUP_ANTHROPIC_BASE_URL;
const authToken = process.env.SETUP_ANTHROPIC_AUTH_TOKEN;
const model = process.env.SETUP_ANTHROPIC_MODEL;

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

/*
 * Keep existing settings and update only the fields
 * required for AgentRouter.
 */

if (!settings.env || typeof settings.env !== "object" || Array.isArray(settings.env)) {
    settings.env = {};
}

settings.env.ANTHROPIC_BASE_URL = baseUrl;
settings.env.ANTHROPIC_AUTH_TOKEN = authToken;
settings.env.ANTHROPIC_MODEL = model;

/*
 * Set the default Claude Code model.
 */
settings.model = model;

/*
 * Schema support for VS Code autocomplete/validation.
 */
settings["$schema"] =
    "https://json.schemastore.org/claude-code-settings.json";

fs.writeFileSync(
    settingsFile,
    JSON.stringify(settings, null, 2) + "\n",
    { mode: 0o600 }
);

fs.chmodSync(settingsFile, 0o600);
NODE

success "Claude settings configured:"
echo "       $CLAUDE_SETTINGS"

# ============================================================
# 21. Verify settings.json
# ============================================================

if [ -f "$CLAUDE_SETTINGS" ]; then

    if node -e '
        const fs = require("fs");
        const file = process.argv[1];
        JSON.parse(fs.readFileSync(file, "utf8"));
    ' "$CLAUDE_SETTINGS"; then

        success "settings.json is valid JSON."

    else

        error "settings.json validation failed."
        exit 1

    fi

else

    error "settings.json was not created."
    exit 1

fi

# ============================================================
# 22. Verify Claude Code
# ============================================================

echo
echo "============================================================"
echo "                    VERIFICATION"
echo "============================================================"
echo

success "Operating system : $OS_NAME"
success "Shell            : $CURRENT_SHELL"
success "Node.js          : $(node --version)"
success "npm              : $(npm --version)"

if command -v claude >/dev/null 2>&1; then
    success "Claude Code      : installed"
    info "Claude path      : $(command -v claude)"
else
    error "Claude Code      : NOT FOUND"
    exit 1
fi

if [ "$VS_CODE_AVAILABLE" = "true" ]; then

    if code --list-extensions 2>/dev/null | \
        grep -Fxq "anthropic.claude-code"; then

        success "Claude VS Code extension : installed"

    else

        warn "Claude VS Code extension : not detected"

    fi

else

    warn "Claude VS Code extension : skipped (VS Code not detected)"

fi

# ============================================================
# 23. Verify AgentRouter settings
# ============================================================

echo
echo "AgentRouter configuration:"
echo
echo "  Base URL : $ANTHROPIC_BASE_URL"
echo "  Model    : $ANTHROPIC_MODEL"
echo "  Token    : configured"
echo

# ============================================================
# 24. Final message
# ============================================================

echo
echo "============================================================"
echo "                 SETUP COMPLETE"
echo "============================================================"
echo
echo "Claude Code CLI:"
echo
echo "  claude"
echo
echo "Claude settings:"
echo
echo "  $CLAUDE_SETTINGS"
echo
echo "Shell configuration:"
echo
echo "  $SHELL_CONFIG"
echo
echo "IMPORTANT:"
echo
echo "Run this command once:"
echo
echo "  source $SHELL_CONFIG"
echo
echo "Then verify Claude:"
echo
echo "  claude --version"
echo
echo "Then start Claude Code:"
echo
echo "  claude"
echo
echo "============================================================"
