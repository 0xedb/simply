package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/thebashshell/simply/lexer"
	"github.com/thebashshell/simply/token"
)

const (
	PROMPT = "~~~"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		lx := lexer.New(line)

		for tok := lx.NextToken(); tok.Kind != token.EOF; tok = lx.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
