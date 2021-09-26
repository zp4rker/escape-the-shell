package command

import (
	"golang.org/x/term"
	"zp4rker.com/escape-the-shell/zterm"
)

func handleQuit(terminal *term.Terminal) error {
	zterm.Writeln(terminal, "Exiting escape-the-shell now...")
	return QuitRequest
}
