package exec

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"

	"strings"

	"github.com/x-color/monkey/evaluator"
	"github.com/x-color/monkey/lexer"
	"github.com/x-color/monkey/object"
	"github.com/x-color/monkey/parser"
)

const (
	prompt        = ">> "
	promptInBlock = ".. "
)

// Repl starts monkey programing language prompt
func Repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		for strings.Count(line, "{")-strings.Count(line, "}") > 0 {
			fmt.Print(promptInBlock)
			if !scanner.Scan() {
				return
			}
			line += scanner.Text()
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// ExecFile executes monkey programing language source file
func ExecFile(fileName string, out io.Writer) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	code := string(bytes)

	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParseErrors(out, p.Errors())
	}

	evaluator.Eval(program, object.NewEnvironment())
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
