package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/x-color/monkey/lexer"
	"github.com/x-color/monkey/token"
)

const prompt = ">> "

// Start starts monkey programing language prompt
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.Eof; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
