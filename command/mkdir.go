package command

import (
	"golang.org/x/term"
	"zp4rker.com/escape-the-shell/fs"
	"zp4rker.com/escape-the-shell/termio"
)

func handleMkdir(terminal *term.Terminal, args []string) error {
	if search := fs.Find(args[0]); search != nil {
		termio.Writeln(terminal, "A file or directory already exists with that name!")
		return nil
	}

	fs.MkDir(args[0])
	return nil
}
