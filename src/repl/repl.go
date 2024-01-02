package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tymbaca/gomonkey/src/lexer"
	"github.com/tymbaca/gomonkey/src/token"
)

var PROMPT = "In [%d]: "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for i := 0; ; i++ {
		fmt.Printf(PROMPT, i)
		ok := scanner.Scan()
		if !ok {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		tok := l.NextToken()
		for tok.Type != token.EOF {
			fmt.Printf("%+v\n", tok)
			tok = l.NextToken()
		}
	}
}
