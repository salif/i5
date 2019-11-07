// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

type Scanner struct {
	code     []byte
	length   int
	position int
	line     uint32
}

func (sc *Scanner) Init(code []byte) {
	sc.code = code
	sc.length = len(code)
	sc.position = 0
	sc.line = 1
}

func (sc *Scanner) NextLine() {
	sc.line++
}

func (sc *Scanner) Line() uint32 {
	return sc.line
}

func (sc *Scanner) HasNext() bool {
	return (sc.position < sc.length)
}

func (sc *Scanner) Next() {
	sc.position++
}

func (sc *Scanner) Peek() byte {
	return sc.code[sc.position]
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
