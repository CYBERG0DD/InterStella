```bash
rm -rf ~/.claude
npm uninstall -g claude-code
npm uninstall -g @anthropic-ai/claude-code claude-code 2>/dev/null || true
rm -f ~/.local/bin/claude
rm -f ~/.local/share/applications/claude-code-url-handler.desktop
rm -f ~/.local/share/applications/claude.desktop
rm -f /usr/local/bin/claude
rm -f /usr/bin/claude
rm -rf ~/.claude
hash -r
echo "=== Claude command ==="
command -v claude || echo "Claude command: NOT FOUND"
echo
echo "=== Remaining Claude files ==="
find ~/.local ~/.npm /usr/local/bin /usr/bin -iname '*claude*' 2>/dev/null || echo "No Claude files found"
```
