# helmsec dec — decrypt a SOPS-encrypted file

Decrypt a SOPS-encrypted file and write the plaintext to `<filename>.dec`.
The `.dec` file should be added to `.gitignore` to avoid committing secrets.

## Usage

```
helmsec dec <file-or-pattern>
```

## Output

The decrypted content is written to `<input-file>.dec`. For example:

| Input file              | Output file                 |
|-------------------------|-----------------------------|
| `secrets-dev-local.yaml` | `secrets-dev-local.yaml.dec` |
| `config/prod.yaml`      | `config/prod.yaml.dec`      |

Files already ending in `.dec` are skipped.

## Examples

```bash
# Decrypt a single file
helmsec dec secrets-dev-local.yaml

# Decrypt all encrypted YAML files in the current directory
helmsec dec "*.yaml"

# Decrypt a file in a subdirectory
helmsec dec config/secrets.yaml
```

## Tip

Add `*.dec` to your `.gitignore` to ensure plaintext files are never committed:

```gitignore
*.dec
```
