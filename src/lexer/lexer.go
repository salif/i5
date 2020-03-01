// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package lexer

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Lexer struct {
	FileName string
	Tokens   []constants.Token
	code     []byte
	length   int
	position int
	peek     byte
	line     uint32
}

func (this *Lexer) Init(fileName string, code []byte) {
	this.FileName = fileName
	this.Tokens = make([]constants.Token, 0)
	this.code = code
	this.length = len(code)
	this.position = 0
	this.peek = 0
	this.line = 1
	if this.peekIsValid() {
		this.peek = code[0]
	}
}

func (this *Lexer) Run() error {
	for this.peekIsValid() {

		// if char is '\t' or ' ' or '\r'
		if this.peek == constants.BYTE_T || this.peek == constants.BYTE_SPACE || this.peek == constants.BYTE_R {
			this.next()
			continue
		}

		// if char is '\n'
		if this.peek == constants.BYTE_EOL {
			this.add(constants.Token{Type: constants.TOKEN_EOL, Value: constants.TOKEN_EOL, Line: this.line})
			this.line++
			this.next()
			continue
		}

		// if char is '\'
		if this.peek == constants.BYTE_B {
			this.next()
			// if char is '\r'
			if this.peekIsValid() && this.peek == constants.BYTE_R {
				this.next()
			}
			if !this.peekIsValid() {
				return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_UNEXPECTED_TOKEN, constants.TOKEN_EOF), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
			}
			// if char is '\n'
			if this.peek == constants.BYTE_EOL {
				this.next()
				this.line++
			} else {
				return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_UNEXPECTED_TOKEN, string(constants.BYTE_B)), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
			}
			continue
		}

		// if char is number(0-9)
		if this.peek >= constants.BYTE_FNUM && this.peek <= constants.BYTE_LNUM {
			var value string = ""

			// if char is number(0-9)
			for ; this.peekIsValid() && (this.peek >= constants.BYTE_FNUM && this.peek <= constants.BYTE_LNUM); this.next() {
				value += string(this.peek)
			}

			// if char is '.'
			if this.peekIsValid() && this.peek == constants.BYTE_DOT {
				value += string(constants.BYTE_DOT)
				this.next()

				// if char is number(0-9)
				for ; this.peekIsValid() && (this.peek >= constants.BYTE_FNUM && this.peek <= constants.BYTE_LNUM); this.next() {
					value += string(this.peek)
				}
				this.add(constants.Token{Type: constants.TOKEN_FLOAT, Value: value, Line: this.line})
			} else {
				this.add(constants.Token{Type: constants.TOKEN_INTEGER, Value: value, Line: this.line})
			}
			continue
		}

		// if char is '#'
		if this.peek == constants.BYTE_HASH {
			this.next()
			for ; this.peekIsValid() && this.peek != constants.BYTE_EOL; this.next() {
			}
			continue
		}

		// if char is '`'
		if this.peek == constants.BYTE_G {
			this.next()
			for {
				if !this.peekIsValid() {
					return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_EXPECTED_FOUND, string(constants.BYTE_G), constants.TOKEN_EOF), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
				}
				if this.peek == constants.BYTE_G {
					this.next()
					break
				} else {
					this.next()
				}
			}
			continue
		}

		// if char is '"'
		if this.peek == constants.BYTE_DQ {
			this.next()
			var value string = ""

			for {
				if !this.peekIsValid() {
					return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_EXPECTED_FOUND, string(constants.BYTE_DQ), constants.TOKEN_EOF), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
				}
				if this.peek == constants.BYTE_DQ {
					this.next()
					break
				}
				// if char is '\n'
				if this.peek == constants.BYTE_EOL {
					this.next()
					this.line++
					continue
				}
				// if char is '\'
				if this.peek == constants.BYTE_B {
					this.next()
					if !this.peekIsValid() {
						return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_UNEXPECTED_TOKEN, string(constants.BYTE_B)), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
					}
					value += this.escape(this.peek)
					continue
				}
				value += string(this.peek)
				this.next()
			}
			this.add(constants.Token{Type: constants.TOKEN_STRING, Value: value, Line: this.line})
			continue
		}

		// if char is "'"
		if this.peek == constants.BYTE_SQ {
			this.next()
			var value string = ""

			for {
				if !this.peekIsValid() {
					return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_EXPECTED_FOUND, string(constants.BYTE_DQ), constants.TOKEN_EOF), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
				}
				if this.peek == constants.BYTE_SQ {
					this.next()
					break
				}
				if this.peek < 32 || this.peek > 126 {
					return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_UNEXPECTED_TOKEN_C, string(this.peek), this.peek), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
				}
				value += string(this.peek)
				this.next()
			}
			this.add(constants.Token{Type: constants.TOKEN_STRING, Value: value, Line: this.line})
			continue
		}

		// if char is '$'
		if this.peek == constants.BYTE_DOLAR {
			this.next()
			var value = this.readIdent()
			this.add(constants.Token{Type: constants.TOKEN_BUILTIN, Value: value, Line: this.line})
			continue
		}

		// if char is '_' or string(a-z) or string(A-Z)
		if this.peek == constants.BYTE_U || (this.peek >= constants.BYTE_FSC && this.peek <= constants.BYTE_LSC) ||
			(this.peek >= constants.BYTE_FBC && this.peek <= constants.BYTE_LBC) {

			var value string = this.readIdent()
			v, e := constants.KEYWORDS[value]
			if e {
				this.add(constants.Token{Type: v, Value: value, Line: this.line})
			} else {
				this.add(constants.Token{Type: constants.TOKEN_IDENTIFIER, Value: value, Line: this.line})
			}
			continue
		}

		switch this.peek {

		// =
		case constants.BYTE_EQ:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_EQEQ, Value: constants.TOKEN_EQEQ, Line: this.line})
				this.next()
			} else if this.peekIsValid() && this.peek == constants.BYTE_GT {
				this.add(constants.Token{Type: constants.TOKEN_EQGT, Value: constants.TOKEN_EQGT, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_EQ, Value: constants.TOKEN_EQ, Line: this.line})
			}
		// !
		case constants.BYTE_NOT:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_NOTEQ, Value: constants.TOKEN_NOTEQ, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_NOT, Value: constants.TOKEN_NOT, Line: this.line})
			}
		// +
		case constants.BYTE_PLUS:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_PLUSEQ, Value: constants.TOKEN_PLUSEQ, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_PLUS, Value: constants.TOKEN_PLUS, Line: this.line})
			}
		// -
		case constants.BYTE_MINUS:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_MINUSEQ, Value: constants.TOKEN_MINUSEQ, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_MINUS, Value: constants.TOKEN_MINUS, Line: this.line})
			}
		// *
		case constants.BYTE_MULTIPLY:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_MULTIPLYEQ, Value: constants.TOKEN_MULTIPLYEQ, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_MULTIPLY, Value: constants.TOKEN_MULTIPLY, Line: this.line})
			}
		// /
		case constants.BYTE_DIVIDE:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_DIVIDEEQ, Value: constants.TOKEN_DIVIDEEQ, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_DIVIDE, Value: constants.TOKEN_DIVIDE, Line: this.line})
			}
		// %
		case constants.BYTE_MODULO:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_MODULOEQ, Value: constants.TOKEN_MODULOEQ, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_MODULO, Value: constants.TOKEN_MODULO, Line: this.line})
			}
		// &
		case constants.BYTE_AND:
			this.add(constants.Token{Type: constants.TOKEN_AND, Value: constants.TOKEN_AND, Line: this.line})
			this.next()
		// |
		case constants.BYTE_OR:
			this.add(constants.Token{Type: constants.TOKEN_OR, Value: constants.TOKEN_OR, Line: this.line})
			this.next()
		// ^
		case constants.BYTE_XOR:
			this.add(constants.Token{Type: constants.TOKEN_XOR, Value: constants.TOKEN_XOR, Line: this.line})
			this.next()
		// ~
		case constants.BYTE_BNOT:
			this.add(constants.Token{Type: constants.TOKEN_BNOT, Value: constants.TOKEN_BNOT, Line: this.line})
			this.next()
		// <
		case constants.BYTE_LT:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_LTEQ, Value: constants.TOKEN_LTEQ, Line: this.line})
				this.next()
			} else if this.peekIsValid() && this.peek == constants.BYTE_LT {
				this.add(constants.Token{Type: constants.TOKEN_LTLT, Value: constants.TOKEN_LTLT, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_LT, Value: constants.TOKEN_LT, Line: this.line})
			}
		// >
		case constants.BYTE_GT:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_GTEQ, Value: constants.TOKEN_GTEQ, Line: this.line})
				this.next()
			} else if this.peekIsValid() && this.peek == constants.BYTE_GT {
				this.add(constants.Token{Type: constants.TOKEN_GTGT, Value: constants.TOKEN_GTGT, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_GT, Value: constants.TOKEN_GT, Line: this.line})
			}
		// :
		case constants.BYTE_COLON:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_EQ {
				this.add(constants.Token{Type: constants.TOKEN_COLONEQ, Value: constants.TOKEN_COLONEQ, Line: this.line})
				this.next()
			} else if this.peekIsValid() && this.peek == constants.BYTE_COLON {
				this.add(constants.Token{Type: constants.TOKEN_COLONCOLON, Value: constants.TOKEN_COLONCOLON, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_COLON, Value: constants.TOKEN_COLON, Line: this.line})
			}
		// ?
		case constants.BYTE_QM:
			this.next()
			if this.peekIsValid() && this.peek == constants.BYTE_QM {
				this.add(constants.Token{Type: constants.TOKEN_QMQM, Value: constants.TOKEN_QMQM, Line: this.line})
				this.next()
			} else {
				this.add(constants.Token{Type: constants.TOKEN_QM, Value: constants.TOKEN_QM, Line: this.line})
			}
		// .
		case constants.BYTE_DOT:
			this.add(constants.Token{Type: constants.TOKEN_DOT, Value: constants.TOKEN_DOT, Line: this.line})
			this.next()
		// ,
		case constants.BYTE_COMMA:
			this.add(constants.Token{Type: constants.TOKEN_COMMA, Value: constants.TOKEN_COMMA, Line: this.line})
			this.next()
		// (
		case constants.BYTE_LPAREN:
			this.add(constants.Token{Type: constants.TOKEN_LPAREN, Value: constants.TOKEN_LPAREN, Line: this.line})
			this.next()
		// )
		case constants.BYTE_RPAREN:
			this.add(constants.Token{Type: constants.TOKEN_RPAREN, Value: constants.TOKEN_RPAREN, Line: this.line})
			this.next()
		// {
		case constants.BYTE_LBRACE:
			this.add(constants.Token{Type: constants.TOKEN_LBRACE, Value: constants.TOKEN_LBRACE, Line: this.line})
			this.next()
		// }
		case constants.BYTE_RBRACE:
			this.add(constants.Token{Type: constants.TOKEN_RBRACE, Value: constants.TOKEN_RBRACE, Line: this.line})
			this.next()
		default:
			return constants.SyntaxError{Message: fmt.Sprintf(constants.SYNTAX_UNEXPECTED_TOKEN_C, string(this.peek), this.peek), In: fmt.Sprintf("%v:%d", this.FileName, this.line)}
		}
	}
	this.add(constants.Token{Type: constants.TOKEN_EOF, Value: constants.TOKEN_EOF, Line: this.line})
	return nil
}

func (this *Lexer) readIdent() string {
	var value string = ""

	for {
		if !this.peekIsValid() {
			return value
		} else if this.peek == constants.BYTE_U || this.peek == constants.BYTE_DOLAR || (this.peek >= constants.BYTE_FSC && this.peek <= constants.BYTE_LSC) || (this.peek >= constants.BYTE_FNUM && this.peek <= constants.BYTE_LNUM) {
			value += string(this.peek)
			this.next()
		} else if this.peek >= constants.BYTE_FBC && this.peek <= constants.BYTE_LBC {
			value += "_" + string(this.peek+32)
			this.next()
		} else {
			break
		}
	}

	return value
}

func (this *Lexer) add(tkn constants.Token) {
	this.Tokens = append(this.Tokens, tkn)
}

func (this Lexer) peekIsValid() bool {
	return this.position < this.length
}

func (this *Lexer) next() {
	this.position++
	if this.peekIsValid() {
		this.peek = this.code[this.position]
	}
}

func (this *Lexer) escape(char byte) string {
	switch char {
	case 116:
		return string(9) // if char is 't' return '\t'
	case 110:
		return string(10) // if char is 'n' return '\n'
	case 114:
		return string(13) // if char is 'r' return '\r'
	default:
		return string(char) // else return string(char)
	}
}
