package lexer

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/types"
)

func Run(code []byte) (tokens types.TokenList) {
	tokens.Init()
	var scanner Scanner
	scanner.Init(code, func(length int, position int, line int) {
		errors.FatalError(errors.F(errors.SCANNER_OUT_OF_RANGE, line, ""), 1)
	})

	for scanner.HasNext() {

		// if char is "{" or "}" or "(" or ")"
		if IsBracket(string(scanner.Peek())) {
			tokens.Add(types.BRACKET, string(scanner.Peek()), scanner.Line())
			scanner.Next()
			continue
		}

		// if char is "\n"
		if scanner.PeekEquals(10) {
			tokens.Add(types.EOL, types.EOL, scanner.Line())
			scanner.NextLine()
			scanner.Next()
			continue
		}

		// if char is "\t" or " " or "\r"
		if scanner.PeekEquals(9) || scanner.PeekEquals(32) || scanner.PeekEquals(13) {
			scanner.Next()
			continue
		}

		// if char is "\"
		if scanner.PeekEquals(92) {
			scanner.Next()
			scanner.Next()
			continue
		}

		// if char is ","
		if scanner.PeekEquals(44) {
			tokens.Add(types.COMMA, types.COMMA, scanner.Line())
			scanner.Next()
			continue
		}

		// if char is "."
		if scanner.PeekEquals(46) {
			tokens.Add(types.DOT, types.DOT, scanner.Line())
			scanner.Next()
			continue
		}

		// if char is number(0-9)
		if scanner.PeekBetween(48, 57) {
			var value string = ""

			// if char is number(0-9) or "."
			for ; scanner.HasNext() && (scanner.PeekBetween(48, 57) || scanner.PeekEquals(46)); scanner.Next() {
				value += string(scanner.Peek())
			}
			tokens.Add(types.NUMBER, value, scanner.Line())
			continue
		}

		// if char is "#"
		if scanner.PeekEquals(35) {
			scanner.Next()
			for ; scanner.HasNext() && scanner.Until(10); scanner.Next() {
			}
			continue
		}

		// if char is "`"
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
				// if char is "\n"
				if scanner.PeekEquals(10) {
					scanner.NextLine()
				}
				// if char is "\"
				// TODO: if equals(char, 92) {}
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
				// if char is "\n"
				if scanner.PeekEquals(10) {
					scanner.NextLine()
				}
				// if char is "\"
				// TODO: if equals(char, 92) {}
				value += string(scanner.Peek())
			}
			scanner.Next()
			tokens.Add(types.STRING, value, scanner.Line())
			continue
		}

		// if char is "$"
		if scanner.PeekEquals(36) {
			var value string = ""

			// if char is "_" or string(a-z) or number(0-9)
			for ; scanner.HasNext() && (scanner.PeekEquals(95) || scanner.PeekEquals(36) ||
				scanner.PeekBetween(97, 122) || scanner.PeekBetween(48, 57)); scanner.Next() {
				value += string(scanner.Peek())
			}

			tokens.Add(types.BUILTIN, value, scanner.Line())
			continue
		}

		// if char is "_" or string(a-z) or string(A-Z)
		if scanner.PeekEquals(95) || scanner.PeekBetween(97, 122) || scanner.PeekBetween(65, 90) {
			var value string = ""

			// if char is "_" or string(a-z) or string(A-Z) or number(0-9)
			for ; scanner.HasNext() && (scanner.PeekEquals(95) || scanner.PeekBetween(97, 122) ||
				scanner.PeekBetween(65, 90) || scanner.PeekBetween(48, 57)); scanner.Next() {
				value += string(scanner.Peek())
			}

			if IsKeyword(value) {
				tokens.Add(types.KEYWORD, value, scanner.Line())
			} else {
				tokens.Add(types.IDENTIFIER, value, scanner.Line())
			}
			continue
		}

		if IsOperator(string(scanner.Peek())) {
			var value string = ""

			for ; scanner.HasNext() && IsOperator(string(scanner.Peek())); scanner.Next() {
				value += string(scanner.Peek())
			}
			tokens.Add(types.OPERATOR, value, scanner.Line())
			continue
		}

		errors.FatalError(errors.F(errors.LEXER_UNEXPECTED_TOKEN, scanner.Line(), string(scanner.Peek())), 1)
	}
	tokens.Add(types.EOF, types.EOF, scanner.Line())
	return tokens
}
