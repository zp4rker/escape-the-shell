package command

import (
	"errors"
	"golang.org/x/term"
)

var QuitRequest error = errors.New("requested quit")

func Handle(terminal *term.Terminal, cmd string, args []string) error {
	switch cmd {
	case "exit", "quit":
		return handleQuit(terminal)
	case "ls":
		return handleLs(terminal)
	case "cd":
		return handleCd(terminal, args)
	case "mkdir":
		return handleMkdir(terminal, args)
	default:
		return handleUnrecognised(terminal, cmd)
	}
}