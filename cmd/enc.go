package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var encCmd = &cobra.Command{
	Use:   "enc <file-or-pattern> [<file-or-pattern>...]",
	Short: "Encrypt a plaintext file using SOPS",
	Long:  GetHelp("enc"),
	Args:  cobra.MinimumNArgs(1),
	RunE:  runEnc,
}

func init() {
	rootCmd.AddCommand(encCmd)
}

func runEnc(cmd *cobra.Command, args []string) error {
	var hasError bool
	for _, pattern := range args {
		files, err := filepath.Glob(pattern)
		if err != nil {
			return fmt.Errorf("invalid pattern: %w", err)
		}
		if files == nil {
			files = []string{pattern}
		}

		for _, f := range files {
			if err := encryptFile(f); err != nil {
				fmt.Fprintf(os.Stderr, "error encrypting %s: %v\n", f, err)
				hasError = true
			}
		}
	}
	if hasError {
		return fmt.Errorf("one or more files failed to encrypt")
	}
	return nil
}

func encryptFile(file string) error {
	var src, dst string

	if strings.HasSuffix(file, ".dec") {
		// secrets.yaml.dec → encrypt to secrets.yaml
		src = file
		dst = strings.TrimSuffix(file, ".dec")
	} else {
		// secrets.yaml → check if secrets.yaml.dec exists
		decFile := file + ".dec"
		if _, err := os.Stat(decFile); err == nil {
			// .dec variant exists: encrypt it back to the encrypted filename
			src = decFile
			dst = file
		} else {
			// no .dec file: only encrypt in-place if the file is not already encrypted
			if isEncryptedBySops(file) {
				fmt.Printf("skipping: %s (already encrypted, no .dec file present)\n", file)
				return nil
			}
			src = file
			dst = file
		}
	}

	var sopsArgs []string
	if src == dst {
		sopsArgs = []string{"--encrypt", "--in-place", src}
	} else {
		// src has a .dec extension — sops cannot infer the format from it,
		// so pass the format explicitly based on the destination filename.
		format := sopsFormat(dst)
		sopsArgs = []string{"--encrypt", "--input-type", format, "--output-type", format, "--output", dst, src}
	}

	out, err := exec.Command("sops", sopsArgs...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("sops failed: %w\n%s", err, strings.TrimSpace(string(out)))
	}

	if src != dst {
		fmt.Printf("encrypted: %s → %s\n", src, dst)
	} else {
		fmt.Printf("encrypted: %s\n", dst)
	}
	return nil
}

// isEncryptedBySops checks whether a file already contains SOPS metadata,
// indicating it has been encrypted and should not be re-encrypted in-place.
func isEncryptedBySops(filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false
	}
	ext := strings.ToLower(filepath.Ext(filename))
	content := string(data)
	switch ext {
	case ".json":
		return strings.Contains(content, `"sops":`)
	default: // yaml, env, ini — sops top-level key has no indentation
		for _, line := range strings.Split(content, "\n") {
			if strings.HasPrefix(line, "sops:") {
				return true
			}
		}
		return false
	}
}
