package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/spf13/cobra"
)

var decCmd = &cobra.Command{
	Use:   "dec <file-or-pattern>",
	Short: "Decrypt a SOPS-encrypted file to a .dec plaintext file",
	Long:  GetHelp("dec"),
	Args:  cobra.ExactArgs(1),
	RunE:  runDec,
}

func init() {
	rootCmd.AddCommand(decCmd)
}

func runDec(cmd *cobra.Command, args []string) error {
	pattern := args[0]
	files, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("invalid pattern: %w", err)
	}
	if files == nil {
		files = []string{pattern}
	}

	var hasError bool
	for _, f := range files {
		if strings.HasSuffix(f, ".dec") {
			fmt.Fprintf(os.Stderr, "skipping %s: already a .dec file\n", f)
			continue
		}
		if err := decryptFile(f); err != nil {
			fmt.Fprintf(os.Stderr, "error decrypting %s: %v\n", f, err)
			hasError = true
		}
	}
	if hasError {
		return fmt.Errorf("one or more files failed to decrypt")
	}
	return nil
}

func decryptFile(src string) error {
	format := sopsFormat(src)
	data, err := decrypt.File(src, format)
	if err != nil {
		return err
	}

	out := src + ".dec"
	if err := os.WriteFile(out, data, 0600); err != nil {
		return err
	}
	fmt.Printf("decrypted: %s → %s\n", src, out)
	return nil
}

func sopsFormat(filename string) string {
	name := strings.TrimSuffix(filename, ".dec")
	switch strings.ToLower(filepath.Ext(name)) {
	case ".yaml", ".yml":
		return "yaml"
	case ".json":
		return "json"
	case ".env":
		return "dotenv"
	case ".ini":
		return "ini"
	default:
		return "yaml"
	}
}
