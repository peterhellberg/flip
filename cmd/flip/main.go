/*

flip (╯°□°）╯︵ʇxǝʇ

A command line client used to flip text.

Installation

Just go get the command:

    go get -u github.com/peterhellberg/flip/cmd/flip

Usage

You can flip a string on its own or decorate it with a named emoji

    $ flip table foo         #=> (╯°□°）╯︵ooɟ
    $ flip gopher bar        #=> ʕ╯◔ϖ◔ʔ╯︵ɹɐq
    $ flip '(＃｀д´)ﾉ︵' baz #=> (＃｀д´)ﾉ︵zɐq

*/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peterhellberg/flip"
)

func main() {
	fmt.Println(flipString(os.Args[1:]))
}

func flipString(args []string) string {
	if len(args) > 0 {
		if f, ok := flip.Flippers[args[0]]; len(args) > 1 && ok {
			return f(strings.Join(args[1:], " "))
		}

		if len(args) > 1 {
			return args[0] + flip.UpsideDown(strings.Join(args[1:], " "))
		}

		return flip.UpsideDown(strings.Join(args, " "))
	}

	return ""
}
