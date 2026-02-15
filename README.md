# goshell - Unix Shell written in Go

A custom Unix shell built from scratch in Go to study Linux internals: processes, syscalls, file descriptors, pipes, and signals.

## Features

### Implemented

- **REPL** - interactive read-eval-print loop
- **External commands** - runs any program via `fork` + `execve` + `wait` (`os/exec`)
- **Builtins** - `cd`, `pwd`, `exit`, `export`, `unset`, `env`
- **Environment variables** - `export VAR=value`, `$VAR` expansion in arguments
- **I/O redirection** - `>`, `>>`, `<`, `2>`
- **Pipes** - `cmd1 | cmd2 | cmd3` (arbitrary length pipelines)

### Planned

- Signal handling (Ctrl+C, Ctrl+Z)
- Job control (`&`, `fg`, `bg`, `jobs`)
- Globbing (`*`, `?`, `~`)
- Command substitution `$(cmd)`
- Process inspector (`/proc` reader)

## Project structure

```
main.go       - REPL loop, entry point
parser.go     - tokenizer, $VAR expansion, redirect/pipe parsing
executor.go   - process execution, pipes, I/O redirection
builtins.go   - built-in commands (cd, pwd, exit, export, unset, env)
```

## Build & Run

```bash
go run .
```

## Linux internals covered

| Stage | Topic | Key syscalls |
|-------|-------|-------------|
| 1. REPL & processes | Process lifecycle | `fork`, `execve`, `wait4` |
| 2. Environment | Environment passing | `execve(path, argv, envp)` |
| 3. Redirects | File descriptors | `open`, `close`, `dup2` |
| 4. Pipes | IPC, kernel buffers | `pipe`, `dup2`, `close` |
| 5. Signals | Signal delivery | `sigaction`, `kill` |
| 6. Job control | Process groups, sessions | `setpgid`, `tcsetpgrp` |
