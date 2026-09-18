# uninstall-claude-code-thorough.ps1
#
# Windows PowerShell port of the original bash uninstall script.
# Run from PowerShell (Administrator not required for these paths).

Write-Host "=== Removing Claude Code ===" -ForegroundColor Cyan

# ~/.claude config directory
Remove-Item -Path "$env:USERPROFILE\.claude" -Recurse -Force -ErrorAction SilentlyContinue

# Try uninstalling every possible npm package name Claude Code has ever
# shipped under. -ErrorAction SilentlyContinue is PowerShell's equivalent
# to bash's '2>/dev/null || true' here — if the package was never
# installed via npm, this just does nothing rather than stopping the script.
npm uninstall -g claude-code 2>$null
npm uninstall -g "@anthropic-ai/claude-code" claude-code 2>$null

# Native installer binary and version folder (the paths Anthropic's own
# uninstall docs specify)
Remove-Item -Path "$env:USERPROFILE\.local\bin\claude.exe" -Force -ErrorAction SilentlyContinue
Remove-Item -Path "$env:USERPROFILE\.local\share\claude" -Recurse -Force -ErrorAction SilentlyContinue

# The original's two .desktop-file lines are deliberately not ported here.
# .desktop files are a Linux desktop-environment concept — application
# launcher and URL-handler registrations — with no Windows equivalent
# file to delete. Anthropic's own official Windows uninstall
# instructions don't mention any registry cleanup either, which is real
# evidence the native CLI doesn't register anything comparable on
# Windows in the first place. If you're also removing the separate
# Claude Desktop app (not Claude Code, the CLI), that's uninstalled
# through Windows' own "Apps & Features" like any other program — a
# different, unrelated process from anything in this script.

# The original's /usr/local/bin and /usr/bin lines are also intentionally
# left out — both are Unix system-wide binary locations with no Windows
# counterpart at all, not just a different path to swap in.

# Re-check ~/.claude in case anything (like a running Claude Desktop
# instance, per Anthropic's own docs) recreated it since the first removal.
Remove-Item -Path "$env:USERPROFILE\.claude" -Recurse -Force -ErrorAction SilentlyContinue

# No equivalent to bash's 'hash -r' is needed here. That command clears
# bash's internal cache of where it last found a command, so a deleted
# binary isn't wrongly reported as still present later in the same
# session. PowerShell doesn't cache command locations that way at all —
# every command lookup re-resolves against the current PATH fresh, so
# there's no stale cache to clear in the first place.

Write-Host ""
Write-Host "=== Claude command ===" -ForegroundColor Cyan
$claudeCmd = Get-Command claude -ErrorAction SilentlyContinue
if ($claudeCmd) {
    Write-Host $claudeCmd.Source
} else {
    Write-Host "Claude command: NOT FOUND"
}

Write-Host ""
Write-Host "=== Remaining Claude files ===" -ForegroundColor Cyan
$searchPaths = @(
    "$env:USERPROFILE\.local",
    "$env:APPDATA\npm",
    "$env:USERPROFILE\AppData\Roaming\npm"
) | Where-Object { Test-Path $_ }

$remaining = Get-ChildItem -Path $searchPaths -Recurse -Filter "*claude*" -ErrorAction SilentlyContinue

if ($remaining) {
    $remaining | ForEach-Object { Write-Host $_.FullName }
} else {
    Write-Host "No Claude files found"
}
