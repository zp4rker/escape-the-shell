package main

import (
	"golang.org/x/term"
	"os"
	"strings"
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

	writeLine(terminal, "Welcome to escape-the-shell!")
	writeLine(terminal)

	shell: for {
		input, err := terminal.ReadLine()
		if err != nil {
			panic("Failed to read user input!")
		}

		args := strings.Fields(input)
		cmd := strings.ToLower(args[0])
		args = args[1:]

		switch cmd {
		case "exit":
			writeLine(terminal, "Exiting escape-the-shell now...")
			break shell
		default:
			writeLine(terminal, "Unrecognised command:", cmd)
		}
	}
}

func writeLine(t *term.Terminal, strings ...string) {
	strings = append(strings, "\n")
	write(t, strings...)
}

func write(t *term.Terminal, strings ...string) {
	var s string
	if len(strings) > 0 {
		s = strings[0]
		for i := 1; i < len(strings); i++ {
			s += " " + strings[i]
		}
	}

	if _, err := t.Write([]byte(s)); err != nil {
		panic("Failed to write to terminal!")
	}
}