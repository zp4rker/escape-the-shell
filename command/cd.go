package command

import (
	"golang.org/x/term"
	"strings"
	"zp4rker.com/escape-the-shell/fs"
	"zp4rker.com/escape-the-shell/termio"
)

func handleCd(terminal *term.Terminal, args []string) error {
	if len(args) < 1 {
		termio.Writeln(terminal, "You need to provide a directory to change to!")
		return nil
	}

	if strings.HasSuffix(args[0], "/") {
		args[0] = args[0][:len(args[0]) - 1]
	}

	if args[0] == ".." {
		fs.ChDir(nil)
		return nil
	}

	dir := fs.Find(args[0])

	if dir == nil {
		termio.Writeln(terminal, "There is no directory named:", args[0])
		return nil
	}
	switch (*dir).(type) {
	case fs.File:
		termio.Writeln(terminal, "That's a file, not a directory!")
	case fs.Directory:
		d := (*dir).(fs.Directory)
		fs.ChDir(&d)
	}
	return nil
}
