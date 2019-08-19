package lexer

// All tokens:
// ( ) { }
// number string
// keyword bool identifier
// operator dlm
// eol eof

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/types"
)

var (
	keywords map[string]bool = map[string]bool{
		"if": true, "elif": true, "else": true, "for": true, "break": true, "continue": true, "fn": true, "return": true, "try": true, "catch": true}
	bools map[string]bool = map[string]bool{
		"true": true, "false:": true}
	operators map[string]bool = map[string]bool{
		"+": true, "-": true, "*": true, "/": true, "=": true, "&": true, "|": true, "%": true, "!": true, "<": true, ">": true, ".": true, ":": true}
	bbp map[string]bool = map[string]bool{
		"{": true, "}": true, "(": true, ")": true}
)

func Run(code []byte) (tokens types.TokenList) {
	tokens.Init()
	var scanner Scanner
	scanner.Init(code)

	for scanner.HasNext() {

		// if char is "{" or "}" or "(" or ")"
		if contains(bbp, string(scanner.Peek())) {
			tokens.Add(string(scanner.Peek()), string(scanner.Peek()), scanner.Line())
			scanner.Next()
			continue
		}

		// if char is "\n"
		if scanner.PeekEquals(10) {
			tokens.Add("eol", "eol", scanner.Line())
			scanner.NextLine()
			scanner.Next()
			continue
		}

		// if char is "\t" or " " or "\r"
		if scanner.PeekEquals(9) || scanner.PeekEquals(32) || scanner.PeekEquals(13) {
			scanner.Next()
			continue
		}

		// if char is ","
		if scanner.PeekEquals(44) {
			tokens.Add("dlm", ",", scanner.Line())
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
			tokens.Add("number", value, scanner.Line())
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
			tokens.Add("string", value, scanner.Line())
			continue
		}

		// if char is "_" or "$" or string(a-z)
		if scanner.PeekEquals(95) || scanner.PeekEquals(36) || scanner.PeekBetween(97, 122) {
			var value string = ""

			// if char is "_" or string(a-z) or number(0-9)
			for ; scanner.HasNext() && (scanner.PeekEquals(95) || scanner.PeekEquals(36) || scanner.PeekBetween(97, 122) || scanner.PeekBetween(48, 57)); scanner.Next() {
				value += string(scanner.Peek())
			}

			if contains(keywords, value) {
				tokens.Add("keyword", value, scanner.Line())
			} else if contains(bools, value) {
				tokens.Add("bool", value, scanner.Line())
			} else {
				tokens.Add("identifier", value, scanner.Line())
			}
			continue
		}

		if contains(operators, string(scanner.Peek())) {
			var value string = ""

			for ; scanner.HasNext() && contains(operators, string(scanner.Peek())); scanner.Next() {
				value += string(scanner.Peek())
			}
			tokens.Add("operator", value, scanner.Line())
			continue
		}

		errors.FatalLexerError("error: line %v: %v: unexpected token\n", scanner.Line(), string(scanner.Peek()), 1)
	}
	tokens.Add("eof", "eof", 0)
	return tokens
}

func contains(par map[string]bool, char string) bool {
	_, dContains := par[char]
	return dContains
}
