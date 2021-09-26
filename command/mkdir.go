package command

import (
	"golang.org/x/term"
	"zp4rker.com/escape-the-shell/fs"
	"zp4rker.com/escape-the-shell/zterm"
)

func handleMkdir(terminal *term.Terminal, args []string) error {
	if len(args) < 1 {
		zterm.Writeln(terminal, "You need to provide a name for the directory!")
		return nil
	}

	if search := fs.Find(args[0]); search != nil {
		zterm.Writeln(terminal, "A file or directory already exists with that name!")
		return nil
	}

	fs.MkDir(args[0])
	return nil
}
