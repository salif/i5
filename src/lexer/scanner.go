package lexer

import "github.com/i5/i5/src/errors"

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

func (sc *Scanner) NextLine() {
	sc.line++
}

func (sc *Scanner) Line() int {
	return sc.line
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
		errors.NewFatalLexerError(errors.SCANNER_OUT_OF_RANGE, sc.line, "", 1)
		return 0
	}
}

func (sc *Scanner) Until(char byte) bool {
	return sc.Peek() != char
}

func (sc *Scanner) PeekEquals(char byte) bool {
	return sc.Peek() == char
}

func (sc *Scanner) PeekBetween(first byte, second byte) bool {
	return (sc.Peek() >= first && sc.Peek() <= second)
}
