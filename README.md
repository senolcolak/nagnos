# AI-Assisted Tmux Diagnostics Environment
### Smart Troubleshooting for System Administrators, DevOps, and SRE Teams

This project provides a next-generation **tmux-based diagnostic environment** that assists system administrators in understanding, analyzing, and troubleshooting host systems instantly — **without manual setup or configuration**.

It combines:
- Real-time command suggestions  
- Static and dynamic system analysis  
- Automatic session recording  
- A knowledgebase-driven remote diagnostics engine  
- Optional multi-user collaboration  

The goal: **give engineers instant situational awareness the moment they SSH into a server.**

---

## 🧩 **Key Features**

### 🔹 Automatic Tmux Workspace
- Three-pane layout:
  1. **Main terminal**
  2. **Suggestion pane**
  3. **Analysis pane**
- Launches automatically when an admin logs in.

### 🔹 Intelligent Command Suggestions
- Detects incomplete or incorrect commands.
- Displays most-used completions and safer alternatives.
- Example: typing `lvcreate` instantly shows full recommended usage.

### 🔹 Static System Analysis (Offline)
- OS, CPU, RAM, inode, disk usage
- Running services, open ports
- Misconfiguration detection
- PHP, Nginx, MySQL, Kubernetes best-practice checks
- Hypervisor information
- Security observations

### 🔹 Dynamic Analysis (Online)
- Queries a remote knowledgebase for tailored recommendations
- Fills analysis templates and generates actionable routines
- Works via:
  - **Bash Runner (v1)**: Collects baseline runtime info
  - **Go Agent (v2)**: Cross-compiled, secure, advanced collector

### 🔹 Session Recording
- Every tmux session is recorded by default
- Choose **local or remote** storage during installation
- Encrypted and access-controlled

### 🔹 Future Features (Roadmap)
- Remote-driven diagnostics (diagnostic packs pushed on demand)
- Multi-user collaboration (tmate-style)
- Plugin system for new service checks
- Kubernetes cluster-wide intelligence

---

## 🚀 **Why Nagnos is the key for you to manage the systems better?**

Nagnos is the Senior System Administrator Assistant that will guide you not to make any mistakes and behave like you, when you are not there.

System administrators spend a significant amount of time manually running the same commands every time they log into a server:

```bash
top
df -h
journalctl -xe
free -m
netstat -tulpen
ps aux
systemctl status <service>
```

This tmux environment does that automatically and adds intelligence:

- Suggests commands
- Analyzes system health
- Detects misconfigurations
- Generates recommended fixes
- Helps junior engineers act like senior engineers
- Provides repeatable, standardized troubleshooting

