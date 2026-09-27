# Dangerous Command Reference

> ⚠️ **For educational purposes only — never run on production systems.**

## Danger Rating Legend

- **9–10/10 — Critical:** Immediate, unrecoverable destruction
- **7–8/10 — High:** Severe damage, very difficult to recover
- **5–6/10 — Medium:** Significant disruption, recoverable with effort
- **1–4/10 — Low:** Disruptive but limited damage

## Linux Commands

| Command | OS | Effects | Legitimate Use | Danger |
|---|---|---|---|---|
| `rm -rf /` | Linux | Recursively force-deletes everything from root. Wipes entire OS, all files, no recovery. | System destruction / stress testing (VM only) | 10/10 |
| `:(){ :|:& };:` | Linux | Fork bomb. Creates infinite processes, freezes system instantly due to resource exhaustion. | Security research / testing system limits | 9/10 |
| `dd if=/dev/zero of=/dev/sda` | Linux | Overwrites entire hard drive with zeros. Permanent, unrecoverable data destruction. | Secure disk wiping before disposal | 9/10 |
| `chmod -R 000 /` | Linux | Removes all read/write/execute permissions from every file. System becomes completely unusable. | None — purely destructive | 9/10 |
| `mkfs.ext4 /dev/sda` | Linux | Formats entire drive, destroys all data and partitions instantly. | Formatting new/wiped drives for fresh install | 8/10 |
| `shred -vfz /dev/sda` | Linux | Overwrites drive multiple times with random data. Even forensics can't recover data. | Secure permanent data erasure | 8/10 |
| `cat /dev/urandom > /dev/mem` | Linux | Writes random data directly into RAM. Causes instant kernel panic and system crash. | None — hardware damage risk | 9/10 |
| `chown -R nobody /` | Linux | Strips ownership from every file. System can no longer identify file owners. Breaks everything. | None in production — destructive | 8/10 |
| `wget -O- url \| sh` | Linux | Downloads and immediately executes a script. Blind code execution — malware risk. | Quick installs (risky without verifying source) | 7/10 |
| `mv / /dev/null` | Linux | Attempts to move entire root filesystem to null. Destroys system structure completely. | None — purely destructive | 9/10 |
| `echo 1 > /proc/sysrq-trigger` | Linux | Triggers immediate kernel-level system crash with no warning. | Emergency shutdown when system is frozen | 6/10 |
| `iptables -F` | Linux | Flushes all firewall rules. System becomes completely open to network attacks immediately. | Resetting firewall rules during config | 6/10 |
| `fdisk /dev/sda (delete all)` | Linux | Destroys all partition tables. OS can no longer find or access the drive. | Repartitioning drives during setup | 7/10 |
| `history \| sh` | Linux | Re-executes every command in your shell history. Any past dangerous command runs again. | Scripting repeated workflows (carefully) | 7/10 |
| `python -c 'import os; os.system("rm -rf /")'` | Linux | Python wrapper for full system deletion. Bypasses some shell-level protections. | None — same as `rm -rf /` but harder to spot | 10/10 |

## Windows Commands

| Command | OS | Effects | Legitimate Use | Danger |
|---|---|---|---|---|
| `format C: /q` | Windows | Quick formats the C: drive. Destroys Windows OS, all installed programs and user data. | Wiping drives before OS reinstall | 9/10 |
| `del /f /s /q C:\*.*` | Windows | Force-deletes every file in C: drive silently. No confirmation, no recycle bin. | Batch file cleanup in controlled environments | 9/10 |
| `rd /s /q C:\` | Windows | Removes every directory from C: drive recursively. System cannot recover. | None in production — destructive | 9/10 |
| `reg delete HKLM\SYSTEM` | Windows | Deletes the system registry hive. Windows cannot boot at all after this. | None — catastrophic | 10/10 |
| `reg delete HKLM\SOFTWARE` | Windows | Removes all software registry entries. Every installed program stops working. | None — catastrophic | 9/10 |
| `diskpart > clean` | Windows | Wipes the entire disk partition table. All data, all partitions gone. | Preparing drives for fresh OS installation | 8/10 |
| `netsh firewall set opmode disable` | Windows | Completely disables Windows Firewall. System exposed to all network attacks. | Firewall troubleshooting (temporarily) | 7/10 |
| `cipher /w:C` | Windows | Overwrites all deleted files on C: making forensic recovery impossible. | Secure deletion of sensitive files | 6/10 |
| `shutdown /s /f /t 0` | Windows | Forces immediate shutdown with zero warning, killing all unsaved work. | Remote admin forced shutdown | 5/10 |
| `Get-Process \| Stop-Process` | Windows PowerShell | Terminates every running process including system processes. | Bulk closing apps in testing environments | 7/10 |
| `Remove-Item -Recurse -Force C:\` | Windows PowerShell | PowerShell equivalent of deleting entire C: drive. Same effect as `del /f /s /q`. | None in production — destructive | 9/10 |
| `Set-ExecutionPolicy Unrestricted` | Windows | Allows ANY script to run with no security checks. Opens door for malware. | Development environments only | 6/10 |
| `bcdedit /deletevalue {default} nx` | Windows | Disables Data Execution Prevention (DEP). Makes system vulnerable to code injection attacks. | Legacy software compatibility testing | 7/10 |
| `netsh int ip reset` | Windows | Resets all TCP/IP settings. Kills all network connections until reconfigured. | Fixing corrupted network stack | 5/10 |
| `wmic shadowcopy delete` | Windows | Deletes ALL system restore points and shadow copies. No rollback possible. | Disk space cleanup (removes recovery options) | 8/10 |

## Safety Note

Always test destructive commands in isolated VMs. Never run on live systems without full backups.
