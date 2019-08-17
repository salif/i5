package lexer

// All tokens:
// ( ) { }
// number string
// keyword identifier
// operator dlm
// eol eof

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/types"
	"strings"
)

var (
	keywords map[string]bool = map[string]bool{
		"if": true, "elif": true, "else": true, "for": true, "break": true,
		"continue": true, "fn": true, "return": true, "try": true, "catch": true, "import": true}
	bools map[string]bool = map[string]bool{
		"true": true, "false": true}
	operators map[string]bool = map[string]bool{
		"+": true, "-": true, "*": true, "/": true, "=": true, "&": true, "|": true,
		"%": true, "!": true, "<": true, ">": true, ".": true, ":": true, "@": true, "?": true}
	bbp map[string]bool = map[string]bool{
		"{": true, "}": true, "(": true, ")": true}
)

func Run(code []byte) (tokens types.TokenList) {
	tokens.Init()
	var scanner Scanner
	scanner.Init(code)

	for scanner.HasNext() {
		char := scanner.Peek()

		if contains(bbp, string(char)) {
			tokens.Add(string(char), string(char), scanner.line)
			scanner.Next()
			continue
		}

		// if char is "\n"
		if equals(char, 10) {
			tokens.Add("eol", "eol", scanner.line)
			scanner.line++
			scanner.Next()
			continue
		}

		// if char is "\t" or " " or "\r"
		if equals(char, 9) || equals(char, 32) || equals(char, 13) {
			scanner.Next()
			continue
		}

		// if char is ","
		if equals(char, 44) {
			tokens.Add("dlm", ",", scanner.line)
			scanner.Next()
			continue
		}

		// if char is number(0-9)
		if between(char, 48, 57) {
			var value string = ""

			// if char is number(0-9) or "."
			for scanner.HasNext() && (between(char, 48, 57) || equals(char, 46)) {
				value += string(char)
				scanner.Next()
				char = scanner.Peek()
			}
			tokens.Add("number", value, scanner.line)
			continue
		}

		// if char is "#"
		if equals(char, 35) {
			for scanner.HasNext() && char != 10 {
				scanner.Next()
				char = scanner.Peek()
			}
			continue
		}

		// if char is '"'
		if equals(char, 34) {
			var value string = ""
			scanner.Next()
			char = scanner.Peek()

			for scanner.HasNext() && char != 34 {
				// if char is "\n"
				if equals(char, 10) {
					scanner.line++
				}
				// if char is "\"
				// TODO: if equals(char, 92) {}
				value += string(char)
				scanner.Next()
				char = scanner.Peek()
			}
			scanner.Next()
			tokens.Add("string", value, scanner.line)
			continue
		}

		// if char is string(A-Z)
		if between(char, 65, 90) {
			char = ([]byte(strings.ToLower(string(char))))[0]
		}

		// if char is "_" or "$" or string(a-z)
		if equals(char, 95) || equals(char, 36) || between(char, 97, 122) {
			var value string = ""

			// if char is "_" or string(a-z) or number(0-9)
			for scanner.HasNext() && (equals(char, 95) || equals(char, 36) || between(char, 97, 122) || between(char, 48, 57)) {
				value += string(char)
				scanner.Next()
				char = scanner.Peek()
			}

			if contains(keywords, value) {
				tokens.Add("keyword", value, scanner.line)
			} else if contains(bools, value) {
				tokens.Add("bool", value, scanner.line)
			} else {
				tokens.Add("identifier", value, scanner.line)
			}
			continue
		}

		if contains(operators, string(char)) {
			var value string = ""

			for scanner.HasNext() && contains(operators, string(char)) {
				value += string(char)
				scanner.Next()
				char = scanner.Peek()
			}
			tokens.Add("operator", value, scanner.line)
			continue
		}

		errors.FatalLexerError("error: line %v: %v: unexpected token\n", scanner.line, string(char), 1)
	}
	tokens.Add("eof", "eof", 0)
	return tokens
}

func contains(par map[string]bool, char string) bool {
	_, dContains := par[char]
	return dContains
}

func equals(first byte, second byte) bool {
	return first == second
}

func between(char byte, first byte, second byte) bool {
	return (char >= first && char <= second)
}
