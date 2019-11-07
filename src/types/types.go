// SPDX-License-Identifier: GPL-3.0-or-later
package types

type Token struct {
	Type  string
	Value string
	Line  uint32
}

type TokenList struct {
	list []Token
}

func (ta TokenList) Init() {
	ta.list = make([]Token, 0)
}

func (ta *TokenList) Add(kind string, char string, line uint32) {
	ta.list = append(ta.list, Token{Type: kind, Value: char, Line: line})
}

func (ta TokenList) Get(index int) Token {
	if index >= len(ta.list) {
		return Token{Type: EOF, Value: EOF, Line: 0}
	}
	return ta.list[index]
}

func (ta TokenList) Size() int {
	return len(ta.list)
}
