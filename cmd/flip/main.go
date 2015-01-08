/*

flip (╯°□°）╯︵ʇxǝʇ

A command line client used to flip text.

Installation

Just go get the command:

    go get -u github.com/peterhellberg/flip/cmd/flip

Usage

You can flip a string on its own or decorate it with a named emoji

    $ flip table foo  #=> (╯°□°）╯︵ooɟ
    $ flip gopher bar #=> ʕ╯◔ϖ◔ʔ╯︵ɹɐq
    $ flip baz        #=> zɐq

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
		if args[0] == "table" && len(args) > 1 {
			return flip.Table(strings.Join(args[1:], " "))
		}

		if args[0] == "gopher" && len(args) > 1 {
			return flip.Gopher(strings.Join(args[1:], " "))
		}

		return flip.UpsideDown(strings.Join(args, " "))
	}

	return ""
}
