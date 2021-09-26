package main

import "zp4rker.com/escape-the-shell/zterm"

func main() {
	terminal, cleanup, err := zterm.NewTerminal()
	if err != nil {
		panic("Unable to create terminal!")
	}
	defer cleanup()

	terminal.Writeln("Press any key! Press 'x' to exit.")
	shell: for {
		count, bytes := terminal.Read()
		terminal.Writeln("Count:", count)
		terminal.Write("Input: \"")
		for i, b := range bytes {
			if b == 0 {
				continue
			}
			if i > 0 && i < len(bytes) - 1 {
				terminal.Write(", ")
			}
			terminal.Write(string(b))
		}
		terminal.Writeln("\"")
		terminal.Writeln("Input Raw:", bytes)
		terminal.Writeln()

		if count == 1 && string(bytes[0]) == "x" {
			break shell
		}
	}
	terminal.Writeln("Exiting now...")
}