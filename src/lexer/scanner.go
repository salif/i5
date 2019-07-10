package lexer

import (
	"github.com/i5-lang/i5/src/error"
)

type Scanner struct {
	code     []byte
	lenght   int
	position int
	line     int
}

func (sc *Scanner) Init(code []byte) {
	sc.code = code
	sc.lenght = len(code)
	sc.position = 0
	sc.line = 1
}

func (sc *Scanner) HasNext() bool {
	return (sc.position < sc.lenght)
}

func (sc *Scanner) Next() {
	sc.position++
}

func (sc *Scanner) Peek() byte {
	if sc.HasNext() {
		return sc.code[sc.position]
	} else {
		error.FatalLexerError("error: line %v: %v: index out of range\n", sc.line, "", 1)
		return 0
	}
}
