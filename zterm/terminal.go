package zterm

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-runewidth"
	"github.com/mattn/go-tty"
	"io"
	"strings"
)

type Terminal struct {
	in io.Reader
	out io.Writer
	width, height int
	cursorPos     int
}

func NewTerminal() (*Terminal, func(), error) {
	t, err := tty.Open()
	if err != nil {
		return nil, func() {}, err
	}

	clean, err := t.Raw()
	if err != nil {
		return nil, func() {}, err
	}

	in, out := t.Input(), colorable.NewColorable(t.Output())
	width, height, err := t.Size()
	if err != nil {
		return nil, func() {}, err
	}

	terminal := &Terminal{in, out, width, height, 0}
	return terminal, func() {
		t.Close()
		_ = clean()
	}, nil
}

func (t *Terminal) Writeln(args ...interface{}) {
	var s string
	for _, arg := range args {
		if s != "" {
			s += " "
		}
		s += fmt.Sprintf("%v", arg)
	}

	t.Write(strings.TrimSpace(s))
	t.nextLine()
}

func (t *Terminal) Write(args ...interface{}) {
	var s string
	for _, arg := range args {
		s += fmt.Sprintf("%v", arg)
	}

	if size := runewidth.StringWidth(s); size <= t.remainingCols() {
		_, _ = t.out.Write([]byte(s))
		t.cursorPos += size
	} else {
		for _, r := range []rune(s) {
			size := runewidth.RuneWidth(r)
			if size > t.remainingCols() {
				t.nextLine()
			}
			_, _ = t.out.Write([]byte{byte(r)})
			t.cursorPos += size
		}
	}
}

func (t *Terminal) Read() (int, []byte) {
	b := make([]byte, 5)
	c, _ := t.in.Read(b)
	return c, b
}

func (t *Terminal) nextLine() {
	_, _ = t.out.Write([]byte("\r\n"))
	t.cursorPos = 0
}

func (t *Terminal) remainingCols() int {
	return t.width - t.cursorPos
}
