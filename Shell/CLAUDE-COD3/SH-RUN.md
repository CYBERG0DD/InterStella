### Power shell command to run claude code

```ps1
cd $env:USERPROFILE\Downloads 
& "$env:USERPROFILE\Downloads\claude-code-agentrouter-setup.ps1"

Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
.\claude-code-agentrouter-setup.ps1

.\uninstall-claude-code-thorough.ps1
```

```bash
chmod +x claude_install.sh
chmod +x claude_uninstall.sh

./claude_install.sh
./claude_uninstall.sh
```