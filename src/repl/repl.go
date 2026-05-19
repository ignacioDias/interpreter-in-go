package repl

import (
	"bufio"
	"fmt"
	"interpreter/src/monkey/lexer"
	"interpreter/src/monkey/token"
	"io"
)

const PROMPT = ">> "

type REPL struct {
	in  io.Reader
	out io.Writer
}

func New(in io.Reader, out io.Writer) *REPL {
	return &REPL{
		in:  in,
		out: out,
	}
}

func (r *REPL) Start() {
	scanner := bufio.NewScanner(r.in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
