// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/colors"
)

func getIdentType(value string) string {
	switch value {
	case constants.TOKEN_ANDAND:
		return constants.TOKEN_ANDAND
	case constants.TOKEN_OROR:
		return constants.TOKEN_OROR
	case constants.TOKEN_FN:
		return constants.TOKEN_FN
	case constants.TOKEN_RETURN:
		return constants.TOKEN_RETURN
	case constants.TOKEN_IF:
		return constants.TOKEN_IF
	case constants.TOKEN_ELIF:
		return constants.TOKEN_ELIF
	case constants.TOKEN_ELSE:
		return constants.TOKEN_ELSE
	case constants.TOKEN_SWITCH:
		return constants.TOKEN_SWITCH
	case constants.TOKEN_CASE:
		return constants.TOKEN_CASE
	case constants.TOKEN_LOOP:
		return constants.TOKEN_LOOP
	case constants.TOKEN_BREAK:
		return constants.TOKEN_BREAK
	case constants.TOKEN_THROW:
		return constants.TOKEN_THROW
	case constants.TOKEN_IMPORT:
		return constants.TOKEN_IMPORT
	case constants.TOKEN_AS:
		return constants.TOKEN_AS
	default:
		return constants.TOKEN_IDENTIFIER
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
