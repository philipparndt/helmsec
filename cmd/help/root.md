# helmsec — encrypt and decrypt secret files using SOPS

`helmsec` is a drop-in replacement for `helm secrets dec` / `helm secrets enc` from the
[helm-secrets](https://github.com/jkroepke/helm-secrets) plugin — as a standalone binary,
no Helm plugin required. Decrypted files use a `.dec` suffix so they can be easily `.gitignore`d
while keeping the encrypted originals in version control.

## Usage

```
helmsec <command> [flags]
```

## Commands

| Command      | Description                                              |
|--------------|----------------------------------------------------------|
| `enc`        | Encrypt a plaintext file using SOPS                      |
| `dec`        | Decrypt a SOPS-encrypted file to a `.dec` plaintext file |
| `version`    | Print version information                                |
| `completion` | Generate shell completion scripts                        |

## Configuration

`helmsec` uses SOPS under the hood and respects your `.sops.yaml` configuration file for
encryption rules (key groups, key providers, path regexes, etc.).

## Examples

```bash
# Decrypt an encrypted secrets file for editing
helmsec dec secrets-dev-local.yaml

# Edit the decrypted file
vim secrets-dev-local.yaml.dec

# Encrypt it back
helmsec enc secrets-dev-local.yaml
```
