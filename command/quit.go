package command

import (
	"golang.org/x/term"
	"zp4rker.com/escape-the-shell/termio"
)

func handleQuit(terminal *term.Terminal) error {
	termio.Writeln(terminal, "Exiting escape-the-shell now...")
	return QuitRequest
}
