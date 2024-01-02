package main

import (
	"os"

	"github.com/tymbaca/gomonkey/src/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
