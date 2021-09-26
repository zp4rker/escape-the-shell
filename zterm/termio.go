package zterm

import (
	"golang.org/x/term"
)

func Writeln(t *term.Terminal, strings ...string) {
	strings = append(strings, "\n")
	Write(t, strings...)
}

func Write(t *term.Terminal, strings ...string) {
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