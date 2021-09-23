package command

import (
	"golang.org/x/term"
	"strings"
	"zp4rker.com/escape-the-shell/fs"
	"zp4rker.com/escape-the-shell/termio"
)

func handleLs(terminal *term.Terminal) error {
	var output string
	for _, u := range fs.CurrentDir.Contents {
		output += u.Name() + "\n"
	}
	termio.Writeln(terminal, strings.TrimSpace(output))
	return nil
}
