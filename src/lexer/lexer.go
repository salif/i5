// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

// Scan code and return TokenList or return error
func Run(fileName string, code []byte) ([]constants.Token, error) {
	var tokens []constants.Token
	var scanner Scanner
	scanner.Init(code)

	for scanner.HasNext() {

		// if char is '\n'
		if scanner.PeekEquals(constants.BYTE_EOL) {
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_EOL, Value: constants.TOKEN_EOL, Line: scanner.Line()})
			scanner.NextLine()
			scanner.Next()
			continue
		}

		// if char is '\t' or ' ' or '\r'
		if scanner.PeekEquals(constants.BYTE_T) || scanner.PeekEquals(constants.BYTE_SPACE) || scanner.PeekEquals(constants.BYTE_R) {
			scanner.Next()
			continue
		}

		// if char is '\'
		if scanner.PeekEquals(constants.BYTE_B) {
			scanner.Next()
			// if char is '\r'
			if scanner.PeekEquals(constants.BYTE_R) {
				scanner.Next()
			}
			// if char is '\n'
			if scanner.PeekEquals(constants.BYTE_EOL) {
				scanner.Next()
				scanner.NextLine()
			} else {
				return tokens, newError(fmt.Sprintf("%v:%d", fileName, scanner.Line()), constants.LEXER_UNEXPECTED_TOKEN, string(constants.BYTE_B), constants.BYTE_B)
			}
			continue
		}

		// if char is number(0-9)
		if scanner.PeekBetween(constants.BYTE_FNUM, constants.BYTE_LNUM) {
			var value string = ""

			// if char is number(0-9)
			for ; scanner.HasNext() && scanner.PeekBetween(constants.BYTE_FNUM, constants.BYTE_LNUM); scanner.Next() {
				value += string(scanner.Peek())
			}

			// if char is '.'
			if scanner.HasNext() && scanner.PeekEquals(constants.BYTE_DOT) {
				value += string(constants.BYTE_DOT)
				scanner.Next()

				// if char is number(0-9)
				for ; scanner.HasNext() && scanner.PeekBetween(constants.BYTE_FNUM, constants.BYTE_LNUM); scanner.Next() {
					value += string(scanner.Peek())
				}
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_FLOAT, Value: value, Line: scanner.Line()})
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_INTEGER, Value: value, Line: scanner.Line()})
			}
			continue
		}

		// if char is '#'
		if scanner.PeekEquals(constants.BYTE_HASH) {
			scanner.Next()
			for ; scanner.HasNext() && scanner.Until(constants.BYTE_EOL); scanner.Next() {
			}
			continue
		}

		// if char is '`'
		if scanner.PeekEquals(constants.BYTE_G) {
			scanner.Next()
			for ; scanner.HasNext() && scanner.Until(constants.BYTE_G); scanner.Next() {
			}
			scanner.Next()
			continue
		}

		// if char is '"'
		if scanner.PeekEquals(constants.BYTE_DQ) {
			scanner.Next()
			var value string = ""

			for ; scanner.HasNext() && scanner.Until(constants.BYTE_DQ); scanner.Next() {
				// if char is '\n'
				if scanner.PeekEquals(constants.BYTE_EOL) {
					scanner.NextLine()
				}
				// if char is '\'
				if scanner.PeekEquals(constants.BYTE_B) {
					scanner.Next()
					value += escape(scanner.Peek())
					continue
				}
				value += string(scanner.Peek())
			}
			scanner.Next()
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_STRING, Value: value, Line: scanner.Line()})
			continue
		}

		// if char is "'"
		if scanner.PeekEquals(constants.BYTE_SQ) {
			scanner.Next()
			var value string = ""

			for ; scanner.HasNext() && scanner.Until(constants.BYTE_SQ); scanner.Next() {
				// if char is '\n'
				if scanner.PeekEquals(constants.BYTE_EOL) {
					scanner.NextLine()
				}
				// if char is '\'
				if scanner.PeekEquals(constants.BYTE_B) {
					scanner.Next()
					value += escape(scanner.Peek())
					continue
				}
				value += string(scanner.Peek())
			}
			scanner.Next()
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_STRING, Value: value, Line: scanner.Line()})
			continue
		}

		// if char is '$'
		if scanner.PeekEquals(constants.BYTE_DOLAR) {
			var value string = ""
			scanner.Next()

			// if char is '_' or '$' or string(a-z) or number(0-9)
			for ; scanner.HasNext() && (scanner.PeekEquals(constants.BYTE_U) || scanner.PeekEquals(constants.BYTE_DOLAR) ||
				scanner.PeekBetween(constants.BYTE_FSC, constants.BYTE_LSC) || scanner.PeekBetween(constants.BYTE_FNUM, constants.BYTE_LNUM)); scanner.Next() {
				value += string(scanner.Peek())
			}

			tokens = append(tokens, constants.Token{Type: constants.TOKEN_BUILTIN, Value: value, Line: scanner.Line()})
			continue
		}

		// if char is '_' or string(a-z) or string(A-Z)
		if scanner.PeekEquals(constants.BYTE_U) || scanner.PeekBetween(constants.BYTE_FSC, constants.BYTE_LSC) || scanner.PeekBetween(constants.BYTE_FBC, constants.BYTE_LBC) {
			var value string = ""

			// if char is '_' or string(a-z) or string(A-Z) or number(0-9)
			for ; scanner.HasNext() && (scanner.PeekEquals(constants.BYTE_U) || scanner.PeekBetween(constants.BYTE_FSC, constants.BYTE_LSC) ||
				scanner.PeekBetween(constants.BYTE_FBC, constants.BYTE_LBC) || scanner.PeekBetween(constants.BYTE_FNUM, constants.BYTE_LNUM)); scanner.Next() {
				value += string(scanner.Peek())
			}

			tokens = append(tokens, constants.Token{Type: getIdentType(value), Value: value, Line: scanner.Line()})
			continue
		}

		switch scanner.Peek() {

		// =
		case constants.BYTE_EQ:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_EQEQ, Value: constants.TOKEN_EQEQ, Line: scanner.Line()})
				scanner.Next()
			} else if scanner.Peek() == constants.BYTE_GT {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_EQGT, Value: constants.TOKEN_EQGT, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_EQ, Value: constants.TOKEN_EQ, Line: scanner.Line()})
			}
		// !
		case constants.BYTE_NOT:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_NOTEQ, Value: constants.TOKEN_NOTEQ, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_NOT, Value: constants.TOKEN_NOT, Line: scanner.Line()})
			}
		// +
		case constants.BYTE_PLUS:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_PLUSEQ, Value: constants.TOKEN_PLUSEQ, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_PLUS, Value: constants.TOKEN_PLUS, Line: scanner.Line()})
			}
		// -
		case constants.BYTE_MINUS:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_MINUSEQ, Value: constants.TOKEN_MINUSEQ, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_MINUS, Value: constants.TOKEN_MINUS, Line: scanner.Line()})
			}
		// *
		case constants.BYTE_MULTIPLY:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_MULTIPLYEQ, Value: constants.TOKEN_MULTIPLYEQ, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_MULTIPLY, Value: constants.TOKEN_MULTIPLY, Line: scanner.Line()})
			}
		// /
		case constants.BYTE_DIVIDE:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_DIVIDEEQ, Value: constants.TOKEN_DIVIDEEQ, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_DIVIDE, Value: constants.TOKEN_DIVIDE, Line: scanner.Line()})
			}
		// %
		case constants.BYTE_MODULO:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_MODULOEQ, Value: constants.TOKEN_MODULOEQ, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_MODULO, Value: constants.TOKEN_MODULO, Line: scanner.Line()})
			}
		// &
		case constants.BYTE_AND:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_AND, Value: constants.TOKEN_AND, Line: scanner.Line()})
			scanner.Next()
		// |
		case constants.BYTE_OR:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_OR, Value: constants.TOKEN_OR, Line: scanner.Line()})
			scanner.Next()
		// ^
		case constants.BYTE_XOR:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_XOR, Value: constants.TOKEN_XOR, Line: scanner.Line()})
			scanner.Next()
		// ~
		case constants.BYTE_BNOT:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_BNOT, Value: constants.TOKEN_BNOT, Line: scanner.Line()})
			scanner.Next()
		// <
		case constants.BYTE_LT:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_LTEQ, Value: constants.TOKEN_LTEQ, Line: scanner.Line()})
				scanner.Next()
			} else if scanner.Peek() == constants.BYTE_LT {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_LTLT, Value: constants.TOKEN_LTLT, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_LT, Value: constants.TOKEN_LT, Line: scanner.Line()})
			}
		// >
		case constants.BYTE_GT:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_GTEQ, Value: constants.TOKEN_GTEQ, Line: scanner.Line()})
				scanner.Next()
			} else if scanner.Peek() == constants.BYTE_GT {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_GTGT, Value: constants.TOKEN_GTGT, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_GT, Value: constants.TOKEN_GT, Line: scanner.Line()})
			}
		// :
		case constants.BYTE_COLON:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_EQ {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_COLONEQ, Value: constants.TOKEN_COLONEQ, Line: scanner.Line()})
				scanner.Next()
			} else if scanner.Peek() == constants.BYTE_COLON {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_COLONCOLON, Value: constants.TOKEN_COLONCOLON, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_COLON, Value: constants.TOKEN_COLON, Line: scanner.Line()})
			}
		// ?
		case constants.BYTE_QM:
			scanner.Next()
			if scanner.Peek() == constants.BYTE_QM {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_QMQM, Value: constants.TOKEN_QMQM, Line: scanner.Line()})
				scanner.Next()
			} else {
				tokens = append(tokens, constants.Token{Type: constants.TOKEN_QM, Value: constants.TOKEN_QM, Line: scanner.Line()})
			}
		// .
		case constants.BYTE_DOT:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_DOT, Value: constants.TOKEN_DOT, Line: scanner.Line()})
			scanner.Next()
		// ,
		case constants.BYTE_COMMA:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_COMMA, Value: constants.TOKEN_COMMA, Line: scanner.Line()})
			scanner.Next()
		// (
		case constants.BYTE_LPAREN:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_LPAREN, Value: constants.TOKEN_LPAREN, Line: scanner.Line()})
			scanner.Next()
		// )
		case constants.BYTE_RPAREN:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_RPAREN, Value: constants.TOKEN_RPAREN, Line: scanner.Line()})
			scanner.Next()
		// {
		case constants.BYTE_LBRACE:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_LBRACE, Value: constants.TOKEN_LBRACE, Line: scanner.Line()})
			scanner.Next()
		// }
		case constants.BYTE_RBRACE:
			tokens = append(tokens, constants.Token{Type: constants.TOKEN_RBRACE, Value: constants.TOKEN_RBRACE, Line: scanner.Line()})
			scanner.Next()
		default:
			return tokens, newError(fmt.Sprintf("%v:%d", fileName, scanner.Line()), constants.LEXER_UNEXPECTED_TOKEN, string(scanner.Peek()), scanner.Peek())
		}
	}
	tokens = append(tokens, constants.Token{Type: constants.TOKEN_EOF, Value: constants.TOKEN_EOF, Line: scanner.Line()})
	return tokens, nil
}
