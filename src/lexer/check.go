// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import (
	"fmt"

	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/types"
)

const (
	EQ       byte = 61
	NOT      byte = 33
	PLUS     byte = 43
	MINUS    byte = 45
	MULTIPLY byte = 42
	DIVIDE   byte = 47
	MODULO   byte = 37
	AND      byte = 38
	OR       byte = 124
	XOR      byte = 94
	BNOT     byte = 126
	LT       byte = 60
	GT       byte = 62
	COLON    byte = 58
	DOT      byte = 46
	COMMA    byte = 44
	LPAREN   byte = 40
	RPAREN   byte = 41
	LBRACE   byte = 123
	RBRACE   byte = 125
	LBRACKET byte = 91
	RBRACKET byte = 93
	QM       byte = 63
)

func IsKeyword(char string) (string, bool) {
	switch char {
	case types.ANDAND:
		return types.ANDAND, true
	case types.OROR:
		return types.OROR, true
	case types.IF:
		return types.IF, true
	case types.ELIF:
		return types.ELIF, true
	case types.ELSE:
		return types.ELSE, true
	case types.SWITCH:
		return types.SWITCH, true
	case types.CASE:
		return types.CASE, true
	case types.LOOP:
		return types.LOOP, true
	case types.RETURN:
		return types.RETURN, true
	case types.THROW:
		return types.THROW, true
	default:
		return "", false
	}
}

func escape(char byte) string {
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

func newError(fileName string, text string, format ...interface{}) error {
	return fmt.Errorf("%v%v\n%v%v\n", colors.Red("syntax error: "), fmt.Sprintf(text, format...), colors.Red("in: "), fileName)
}
