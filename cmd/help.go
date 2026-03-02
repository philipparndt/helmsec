package cmd

import _ "embed"

//go:embed help/root.md
var helpRoot string

//go:embed help/enc.md
var helpEnc string

//go:embed help/dec.md
var helpDec string

//go:embed help/version.md
var helpVersion string

//go:embed help/completion.md
var helpCompletion string

func GetHelp(topic string) string {
	switch topic {
	case "root":
		return renderMarkdown(helpRoot)
	case "enc":
		return renderMarkdown(helpEnc)
	case "dec":
		return renderMarkdown(helpDec)
	case "version":
		return renderMarkdown(helpVersion)
	case "completion":
		return renderMarkdown(helpCompletion)
	default:
		return ""
	}
}
