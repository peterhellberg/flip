flip (╯°□°）╯︵ʇxǝʇ
===================

Go library used to flip text.

[![GoDoc](https://godoc.org/github.com/peterhellberg/flip?status.svg)](https://godoc.org/github.com/peterhellberg/flip)

## Command line tool

### Installation

```bash
go get -u github.com/peterhellberg/flip/cmd/flip
```

## Examples

### table.go

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peterhellberg/flip"
)

func main() {
	fmt.Println(flip.Table(strings.Join(os.Args[1:], " ")))
}
```
