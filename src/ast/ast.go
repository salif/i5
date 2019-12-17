// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Node interface {
	GetType() string
	GetLine() uint32
	Debug() string
}

const (
	ASSIGN       string = "assign"
	BLOCK        string = "block"
	BUILTIN      string = "builtin"
	BREAK        string = "break"
	CALL         string = "call"
	FLOAT        string = "float"
	FUNCTION     string = "function"
	FUNCTIONEXPR string = "function expression"
	IDENTIFIER   string = "identifier"
	IF           string = "if"
	IMPORT       string = "import"
	INDEX        string = "index"
	INFIX        string = "infix"
	INTEGER      string = "integer"
	LOOP         string = "loop"
	POSTFIX      string = "postfix"
	PREFIX       string = "prefix"
	PROGRAM      string = "program"
	RETURN       string = "return"
	STRING       string = "string"
	SWITCH       string = "switch"
	TERNARY      string = "ternary"
	THROW        string = "throw"
)
