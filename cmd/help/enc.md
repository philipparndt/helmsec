# helmsec enc — encrypt a file using SOPS

Encrypt a plaintext file using SOPS. The `.sops.yaml` configuration in the current or parent
directory is used to determine the encryption keys.

## Usage

```
helmsec enc <file-or-pattern>
```

## File Resolution

| Input file              | `.dec` exists?  | Action                                                  |
|-------------------------|-----------------|---------------------------------------------------------|
| `secrets.yaml`          | yes             | Encrypts `secrets.yaml.dec` → `secrets.yaml`            |
| `secrets.yaml`          | no              | Encrypts `secrets.yaml` in-place                        |
| `secrets.yaml.dec`      | —               | Encrypts `secrets.yaml.dec` → `secrets.yaml`            |

## Examples

```bash
# Encrypt after editing the .dec file (most common workflow)
helmsec enc secrets-dev-local.yaml

# Encrypt a .dec file directly
helmsec enc secrets-dev-local.yaml.dec

# Encrypt all .dec files in a directory
helmsec enc "*.dec"

# Encrypt a newly created plaintext file in-place
helmsec enc new-secrets.yaml
```
