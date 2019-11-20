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

func getIdentType(value string) string {
	switch value {
	case types.ANDAND:
		return types.ANDAND
	case types.OROR:
		return types.OROR
	case types.FN:
		return types.FN
	case types.LAMBDA:
		return types.LAMBDA
	case types.IF:
		return types.IF
	case types.ELIF:
		return types.ELIF
	case types.ELSE:
		return types.ELSE
	case types.SWITCH:
		return types.SWITCH
	case types.CASE:
		return types.CASE
	case types.LOOP:
		return types.LOOP
	case types.RETURN:
		return types.RETURN
	case types.THROW:
		return types.THROW
	default:
		return types.IDENT
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
