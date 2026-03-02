package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

func renderMarkdown(text string) string {
	if !isTerminal() {
		return text
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		return text
	}

	out, err := renderer.Render(text)
	if err != nil {
		return text
	}

	return out
}

func isTerminal() bool {
	stat, _ := os.Stdout.Stat()
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func printWarning(msg string) {
	if !isTerminal() {
		fmt.Fprintln(os.Stderr, "Warning: "+msg)
		return
	}
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("214")).
		Foreground(lipgloss.Color("214")).
		Padding(0, 1)
	fmt.Fprintln(os.Stderr, style.Render("⚠  "+msg))
}

func printError(msg string) {
	if !isTerminal() {
		fmt.Fprintln(os.Stderr, "Error: "+msg)
		return
	}
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("196")).
		Foreground(lipgloss.Color("196")).
		Padding(0, 1)
	fmt.Fprintln(os.Stderr, style.Render("✗  "+msg))
}
