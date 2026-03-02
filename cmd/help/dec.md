# helmsec dec — decrypt a SOPS-encrypted file

Decrypt a SOPS-encrypted file and write the plaintext to `<filename>.dec`.
The `.dec` file should be added to `.gitignore` to avoid committing secrets.

## Usage

```
helmsec dec [--force] <file-or-pattern>
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--force` | `-f` | Skip the `.gitignore` safety check and decrypt regardless |

## Output

The decrypted content is written to `<input-file>.dec`. For example:

| Input file              | Output file                 |
|-------------------------|-----------------------------|
| `secrets-dev-local.yaml` | `secrets-dev-local.yaml.dec` |
| `config/prod.yaml`      | `config/prod.yaml.dec`      |

Files already ending in `.dec` are skipped.

## Safety check

Before decrypting, helmsec verifies that the output `.dec` file would be matched by `.gitignore`
to prevent accidentally committing plaintext secrets. If the check fails, an error is shown.

Use `--force` to skip this check entirely, or run outside a git repository (a warning will be shown
but decryption will proceed).

## Examples

```bash
# Decrypt a single file
helmsec dec secrets-dev-local.yaml

# Decrypt all encrypted YAML files in the current directory
helmsec dec "*.yaml"

# Decrypt a file in a subdirectory
helmsec dec config/secrets.yaml

# Skip the .gitignore check
helmsec dec --force secrets-dev-local.yaml
```

## Tip

Add `*.dec` to your `.gitignore` to ensure plaintext files are never committed:

```gitignore
*.dec
```
