// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

// Scan code and return TokenList or return error
func Run(fileName string, code []byte) (types.TokenList, error) {
	var tokens types.TokenList
	var scanner Scanner
	tokens.Init()
	scanner.Init(code)

	for scanner.HasNext() {

		// if char is '\n'
		if scanner.PeekEquals(10) {
			tokens.Add(types.EOL, types.EOL, scanner.Line())
			scanner.NextLine()
			scanner.Next()
			continue
		}

		// if char is '\t' or ' ' or '\r'
		if scanner.PeekEquals(9) || scanner.PeekEquals(32) || scanner.PeekEquals(13) {
			scanner.Next()
			continue
		}

		// if char is '\'
		if scanner.PeekEquals(92) {
			scanner.Next()
			// if char is '\r'
			if scanner.PeekEquals(13) {
				scanner.Next()
			}
			// if char is '\n'
			if scanner.PeekEquals(10) {
				scanner.Next()
				scanner.NextLine()
			} else {
				return tokens, newError(fmt.Sprintf("%v:%d", fileName, scanner.Line()), constants.LEXER_UNEXPECTED_TOKEN, string(92), 92)
			}
			continue
		}

		// if char is number(0-9)
		if scanner.PeekBetween(48, 57) {
			var value string = ""

			// if char is number(0-9)
			for ; scanner.HasNext() && scanner.PeekBetween(48, 57); scanner.Next() {
				value += string(scanner.Peek())
			}

			// if char is '.'
			if scanner.HasNext() && scanner.PeekEquals(46) {
				value += "."
				scanner.Next()

				// if char is number(0-9)
				for ; scanner.HasNext() && scanner.PeekBetween(48, 57); scanner.Next() {
					value += string(scanner.Peek())
				}
				tokens.Add(types.FLOAT, value, scanner.Line())
			} else {
				tokens.Add(types.INT, value, scanner.Line())
			}
			continue
		}

		// if char is '#'
		if scanner.PeekEquals(35) {
			scanner.Next()
			for ; scanner.HasNext() && scanner.Until(10); scanner.Next() {
			}
			continue
		}

		// if char is '`'
		if scanner.PeekEquals(96) {
			scanner.Next()
			for ; scanner.HasNext() && scanner.Until(96); scanner.Next() {
			}
			scanner.Next()
			continue
		}

		// if char is '"'
		if scanner.PeekEquals(34) {
			scanner.Next()
			var value string = ""

			for ; scanner.HasNext() && scanner.Until(34); scanner.Next() {
				// if char is '\n'
				if scanner.PeekEquals(10) {
					scanner.NextLine()
				}
				// if char is '\'
				if scanner.PeekEquals(92) {
					scanner.Next()
					value += escape(scanner.Peek())
					continue
				}
				value += string(scanner.Peek())
			}
			scanner.Next()
			tokens.Add(types.STRING, value, scanner.Line())
			continue
		}

		// if char is "'"
		if scanner.PeekEquals(39) {
			scanner.Next()
			var value string = ""

			for ; scanner.HasNext() && scanner.Until(39); scanner.Next() {
				// if char is '\n'
				if scanner.PeekEquals(10) {
					scanner.NextLine()
				}
				// if char is '\'
				if scanner.PeekEquals(92) {
					scanner.Next()
					value += escape(scanner.Peek())
					continue
				}
				value += string(scanner.Peek())
			}
			scanner.Next()
			tokens.Add(types.STRING, value, scanner.Line())
			continue
		}

		// if char is '$'
		if scanner.PeekEquals(36) {
			var value string = ""
			scanner.Next()

			// if char is '_' or string(a-z) or number(0-9)
			for ; scanner.HasNext() && (scanner.PeekEquals(95) || scanner.PeekEquals(36) ||
				scanner.PeekBetween(97, 122) || scanner.PeekBetween(48, 57)); scanner.Next() {
				value += string(scanner.Peek())
			}

			tokens.Add(types.BUILTIN, value, scanner.Line())
			continue
		}

		// if char is '_' or string(a-z) or string(A-Z)
		if scanner.PeekEquals(95) || scanner.PeekBetween(97, 122) || scanner.PeekBetween(65, 90) {
			var value string = ""

			// if char is '_' or string(a-z) or string(A-Z) or number(0-9)
			for ; scanner.HasNext() && (scanner.PeekEquals(95) || scanner.PeekBetween(97, 122) ||
				scanner.PeekBetween(65, 90) || scanner.PeekBetween(48, 57)); scanner.Next() {
				value += string(scanner.Peek())
			}

			if kToken, isKeyword := IsKeyword(value); isKeyword {
				tokens.Add(kToken, value, scanner.Line())
			} else {
				tokens.Add(types.IDENT, value, scanner.Line())
			}
			continue
		}

		switch scanner.Peek() {

		// =
		case EQ:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.EQEQ, types.EQEQ, scanner.Line())
				scanner.Next()
			} else if scanner.Peek() == GT {
				tokens.Add(types.EQGT, types.EQGT, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.EQ, types.EQ, scanner.Line())
			}
		// !
		case NOT:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.NOTEQ, types.NOTEQ, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.NOT, types.NOT, scanner.Line())
			}
		// +
		case PLUS:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.PLUSEQ, types.PLUSEQ, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.PLUS, types.PLUS, scanner.Line())
			}
		// -
		case MINUS:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.MINUSEQ, types.MINUSEQ, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.MINUS, types.MINUS, scanner.Line())
			}
		// *
		case MULTIPLY:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.MULTIPLYEQ, types.MULTIPLYEQ, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.MULTIPLY, types.MULTIPLY, scanner.Line())
			}
		// /
		case DIVIDE:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.DIVIDEEQ, types.DIVIDEEQ, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.DIVIDE, types.DIVIDE, scanner.Line())
			}
		// %
		case MODULO:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.MODULOEQ, types.MODULOEQ, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.MODULO, types.MODULO, scanner.Line())
			}
		// &
		case AND:
			tokens.Add(types.AND, types.AND, scanner.Line())
			scanner.Next()
		// |
		case OR:
			tokens.Add(types.OR, types.OR, scanner.Line())
			scanner.Next()
		// ^
		case XOR:
			tokens.Add(types.XOR, types.XOR, scanner.Line())
			scanner.Next()
		// ~
		case BNOT:
			tokens.Add(types.BNOT, types.BNOT, scanner.Line())
			scanner.Next()
		// <
		case LT:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.LTEQ, types.LTEQ, scanner.Line())
				scanner.Next()
			} else if scanner.Peek() == LT {
				tokens.Add(types.LTLT, types.LTLT, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.LT, types.LT, scanner.Line())
			}
		// >
		case GT:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.GTEQ, types.GTEQ, scanner.Line())
				scanner.Next()
			} else if scanner.Peek() == GT {
				tokens.Add(types.GTGT, types.GTGT, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.GT, types.GT, scanner.Line())
			}
		// :
		case COLON:
			scanner.Next()
			if scanner.Peek() == EQ {
				tokens.Add(types.COLONEQ, types.COLONEQ, scanner.Line())
				scanner.Next()
			} else if scanner.Peek() == COLON {
				tokens.Add(types.COLONCOLON, types.COLONCOLON, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.COLON, types.COLON, scanner.Line())
			}
		// ?
		case QM:
			scanner.Next()
			if scanner.Peek() == QM {
				tokens.Add(types.QMQM, types.QMQM, scanner.Line())
				scanner.Next()
			} else {
				tokens.Add(types.QM, types.QM, scanner.Line())
			}
		// .
		case DOT:
			tokens.Add(types.DOT, types.DOT, scanner.Line())
			scanner.Next()
		// ,
		case COMMA:
			tokens.Add(types.COMMA, types.COMMA, scanner.Line())
			scanner.Next()
		// (
		case LPAREN:
			tokens.Add(types.LPAREN, types.LPAREN, scanner.Line())
			scanner.Next()
		// )
		case RPAREN:
			tokens.Add(types.RPAREN, types.RPAREN, scanner.Line())
			scanner.Next()
		// {
		case LBRACE:
			tokens.Add(types.LBRACE, types.LBRACE, scanner.Line())
			scanner.Next()
		// }
		case RBRACE:
			tokens.Add(types.RBRACE, types.RBRACE, scanner.Line())
			scanner.Next()
		default:
			return tokens, newError(fmt.Sprintf("%v:%d", fileName, scanner.Line()), constants.LEXER_UNEXPECTED_TOKEN, string(scanner.Peek()), scanner.Peek())
		}
	}
	tokens.Add(types.EOF, types.EOF, scanner.Line())
	return tokens, nil
}
