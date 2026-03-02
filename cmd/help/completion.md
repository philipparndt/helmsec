# helmsec completion — generate shell completion scripts

Generate shell completion scripts for helmsec. Supported shells: `bash`, `zsh`, `fish`,
and `powershell`.

## Usage

```
helmsec completion [bash|zsh|fish|powershell]
```

## Setup

### Bash

```bash
source <(helmsec completion bash)
# Or to persist:
helmsec completion bash > /etc/bash_completion.d/helmsec
```

### Zsh

```zsh
helmsec completion zsh > "${fpath[1]}/_helmsec"
```

### Fish

```fish
helmsec completion fish | source
# Or to persist:
helmsec completion fish > ~/.config/fish/completions/helmsec.fish
```

### PowerShell

```powershell
helmsec completion powershell | Out-String | Invoke-Expression
```
