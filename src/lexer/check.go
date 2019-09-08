package lexer

import "github.com/i5/i5/src/types"

const (
	EQ       = 61
	NOT      = 33
	PLUS     = 43
	MINUS    = 45
	MULTIPLY = 42
	DIVIDE   = 47
	MODULO   = 37
	AND      = 38
	OR       = 124
	XOR      = 94
	BNOT     = 126
	LT       = 60
	GT       = 62
	COLON    = 58
	DOT      = 46
	COMMA    = 44
	LPAREN   = 40
	RPAREN   = 41
	LBRACE   = 123
	RBRACE   = 125
	LBRACKET = 91
	RBRACKET = 93
	QM       = 63
)

func IsKeyword(char string) (string, bool) {
	switch char {
	case "if":
		return types.IF, true
	case "elif":
		return types.ELIF, true
	case "else":
		return types.ELSE, true
	case "for":
		return types.FOR, true
	case "break":
		return types.BREAK, true
	case "continue":
		return types.CONTINUE, true
	case "fn":
		return types.FN, true
	case "return":
		return types.RETURN, true
	case "import":
		return types.IMPORT, true
	case "true":
		return types.TRUE, true
	case "false":
		return types.FALSE, true
	case "nil":
		return types.NIL, true
	default:
		return "", false
	}
}
