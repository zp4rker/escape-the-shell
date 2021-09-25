package main

import (
	"golang.org/x/term"
	"io"
	"os"
	"strings"
	"zp4rker.com/escape-the-shell/command"
	"zp4rker.com/escape-the-shell/termio"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic("Encountered an error when initialising a term instance!")
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic("Unable to get terminal size!")
	}

	terminal := term.NewTerminal(os.Stdin, "> ")
	if err := terminal.SetSize(width, height); err != nil {
		panic("Unable to set terminal size!")
	}

	termio.Writeln(terminal, "Welcome to escape-the-shell!")
	termio.Writeln(terminal)

	shell: for {
		input, err := terminal.ReadLine()
		if err != nil {
			if err == io.EOF {
				terminal.SetPrompt("")
				termio.Writeln(terminal, "Exiting escape-the-shell now...")
				break shell
			}
		}

		if strings.TrimSpace(input) == "" {
			continue
		}

		args := strings.Fields(input)
		cmd := strings.ToLower(args[0])
		args = args[1:]

		if err = command.Handle(terminal, cmd, args); err != nil {
			if err == command.QuitRequest {
				break shell
			} else {
				panic("Encountered an unexpected error!")
			}
		}
	}
}