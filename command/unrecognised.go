package command

import (
	"golang.org/x/term"
	"zp4rker.com/escape-the-shell/termio"
)

func handleUnrecognised(terminal *term.Terminal,cmd string) error {
	termio.Writeln(terminal, "Unrecognised command:", cmd)
	return nil
}
