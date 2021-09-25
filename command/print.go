package command

import (
	"golang.org/x/term"
	"zp4rker.com/escape-the-shell/fs"
	"zp4rker.com/escape-the-shell/termio"
)

func handlePrint(terminal *term.Terminal, args []string) error {
	if len(args) < 1 {
		termio.Writeln(terminal, "You need to provide a file to print!")
		return nil
	}

	search := fs.Find(args[0])
	if search == nil {
		termio.Writeln(terminal, "Couldn't find file named:", args[0])
		return nil
	}

	switch (*search).(type) {
	case fs.Directory:
		termio.Writeln(terminal, "You can't print out a directory!")
	case fs.File:
		file := (*search).(fs.File)
		termio.Writeln(terminal, file.Read())
	}
	return nil
}
