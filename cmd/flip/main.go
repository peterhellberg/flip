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
