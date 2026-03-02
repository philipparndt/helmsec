# helmsec

A CLI tool to encrypt and decrypt secret files using [SOPS](https://github.com/getsops/sops).

`helmsec` is a drop-in replacement for the `helm secrets dec` / `helm secrets enc` commands
previously provided by the [helm-secrets](https://github.com/jkroepke/helm-secrets) plugin.
It provides the same `enc`/`dec` workflow as a standalone binary — no Helm plugin required.

`helmsec` follows the convention that decrypted files carry a `.dec` suffix, making it easy
to `.gitignore` plaintext secrets while keeping the encrypted originals in version control.

## Installation

### Homebrew

```bash
brew tap philipparndt/helmsec
brew install philipparndt/helmsec/helmsec
```

### Manual

Download the latest binary for your platform from the
[releases page](https://github.com/philipparndt/helmsec/releases).

## Requirements

- [SOPS](https://github.com/getsops/sops) must be installed and available on your `PATH`
  (used for encryption; decryption uses the SOPS Go library directly)
- A `.sops.yaml` configuration file in your project defining your key provider(s)

## Usage

```
helmsec <command> [flags]
```

### Commands

| Command      | Description                                              |
|--------------|----------------------------------------------------------|
| `enc`        | Encrypt a plaintext file using SOPS                      |
| `dec`        | Decrypt a SOPS-encrypted file to a `.dec` plaintext file |
| `version`    | Print version information                                |
| `completion` | Generate shell completion scripts                        |

## Workflow

```bash
# 1. Decrypt an encrypted secrets file for editing
helmsec dec secrets-dev-local.yaml
# → creates secrets-dev-local.yaml.dec (plaintext, git-ignored)

# 2. Edit the plaintext file
vim secrets-dev-local.yaml.dec

# 3. Encrypt it back — helmsec detects the .dec file automatically
helmsec enc secrets-dev-local.yaml
# → encrypts secrets-dev-local.yaml.dec back to secrets-dev-local.yaml
```

You can also pass the `.dec` file directly or use glob patterns:

```bash
helmsec enc secrets-dev-local.yaml.dec
helmsec dec "*.yaml"
helmsec enc "*.dec"
```

## `enc` file resolution

| Input file          | `.dec` exists? | Action                                              |
|---------------------|----------------|-----------------------------------------------------|
| `secrets.yaml`      | yes            | Encrypts `secrets.yaml.dec` → `secrets.yaml`        |
| `secrets.yaml`      | no             | Encrypts `secrets.yaml` in-place                    |
| `secrets.yaml.dec`  | —              | Encrypts `secrets.yaml.dec` → `secrets.yaml`        |

## Configuration

SOPS encryption is driven by your `.sops.yaml` file. Example using age:

```yaml
creation_rules:
  - path_regex: secrets.*\.yaml$
    age: age1ql3z7hjy54pw3hyww5ayyfg7zqgvc7w3j2elw8zmrj2kg5sfn9aqmcac8p
```

See the [SOPS documentation](https://github.com/getsops/sops#using-sops-yaml-conf-to-select-kmspgp-for-new-files)
for all supported key providers (AWS KMS, GCP KMS, Azure Key Vault, age, PGP).

## `.gitignore` recommendation

```gitignore
*.dec
```

## Shell completion

```bash
# Bash
source <(helmsec completion bash)

# Zsh
helmsec completion zsh > "${fpath[1]}/_helmsec"

# Fish
helmsec completion fish | source
```

## License

[Apache 2.0](LICENSE)
