package command

import (
	"golang.org/x/term"
	"zp4rker.com/escape-the-shell/zterm"
)

func handleUnrecognised(terminal *term.Terminal,cmd string) error {
	zterm.Writeln(terminal, "Unrecognised command:", cmd)
	return nil
}
