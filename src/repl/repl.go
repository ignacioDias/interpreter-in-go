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
		r.out.Write([]byte(PROMPT)) //decorator ">>" at the start of each line
		line, ok := r.getLineFromInput(scanner)
		if !ok {
			return
		}
		r.processInput(line)
	}
}

func (r *REPL) getLineFromInput(scanner *bufio.Scanner) (string, bool) {
	scanned := scanner.Scan()
	if !scanned {
		return "", false
	}
	line := scanner.Text()
	return line, true
}

func (r *REPL) processInput(line string) {
	l := lexer.New(line)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		output := fmt.Sprintf("%+v\n", tok)
		r.out.Write([]byte(output))
	}
}
