# Nagnos Development Prompt

## Goal
Implement the full Nagnos project: an AI-assisted tmux diagnostic environment with automated static/dynamic analysis, real-time command suggestions, session recording, and future remote-driven diagnostics.

## Requirements
- Use Python as primary language
- Use Bash for minimal collectors
- Use Go for cross-compiled dynamic agent (later phase)
- Clean modular architecture
- All modules testable and packaged
- Works on any Linux server via SSH

## Architecture
- tmux session with 3 panes
- Suggestion engine parses user commands live
- Static analyzer runs local checks
- Dynamic analyzer queries remote KB
- Unified config system (YAML)
- Session recording with secure storage
- Plugin-ready design

## Core Modules
1. **tmux layout**
2. **Suggestion Engine**
3. **Static Analyzer**
4. **Dynamic Analyzer**
5. **Agent Layer**
6. **Recording**
7. **CLI / installer**
8. **Configuration manager**

## Deliverables
- Python package `nagnos/`
- Bash runner
- Go agent skeleton
- Full tmux integration
- JSON/YAML rule files
- Docs, tests, examples

## Design Rules
- No hardcoded paths
- JSON/YAML for rules
- Use subprocess carefully
- All output machine-parsable
- Use asyncio for dynamic calls

## Output Requirements
- tmux starts cleanly
- Suggestions appear instantly
- Static analysis runs at startup
- Dynamic analysis updates async
- Logs always recorded
- Safe-mode for dangerous commands
