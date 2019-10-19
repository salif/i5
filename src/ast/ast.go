// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Node interface {
	GetType() string
	Print()
	GetLine() int
}

const (
	ASSIGN     = "assign"
	BLOCK      = "block"
	BOOL       = "bool"
	BUILTIN    = "builtin"
	CALL       = "call"
	EXPRESSION = "expression"
	FLOAT      = "float"
	FUNCTION   = "function"
	IDENTIFIER = "identifier"
	IF         = "if"
	IMPORT     = "import"
	INDEX      = "index"
	INFIX      = "infix"
	INTEGER    = "integer"
	LOOP       = "loop"
	POSTFIX    = "postfix"
	PREFIX     = "prefix"
	PROGRAM    = "program"
	RETURN     = "return"
	STRING     = "string"
	SWITCH     = "switch"
	THROW      = "throw"
	TRY        = "try"
)
