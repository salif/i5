package lexer

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

func IsKeyword(char string) bool {
	switch char {
	case "if":
		return true
	case "elif":
		return true
	case "else":
		return true
	case "for":
		return true
	case "break":
		return true
	case "continue":
		return true
	case "fn":
		return true
	case "return":
		return true
	case "try":
		return true
	case "catch":
		return true
	case "import":
		return true
	case "true":
		return true
	case "false":
		return true
	case "nil":
		return true
	default:
		return false
	}
}
