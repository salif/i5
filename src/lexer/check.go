package lexer

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

func IsOperator(char string) bool {
	switch char {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	case "=":
		return true
	case "&":
		return true
	case "|":
		return true
	case "%":
		return true
	case "!":
		return true
	case "<":
		return true
	case ">":
		return true
	case ":":
		return true
	default:
		return false
	}
}

func IsBracket(char string) bool {
	switch char {
	case "{":
		return true
	case "}":
		return true
	case "(":
		return true
	case ")":
		return true
	case "[":
		return true
	case "]":
		return true
	default:
		return false
	}
}
