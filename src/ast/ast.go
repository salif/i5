// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Node interface {
	GetType() string
	GetLine() uint32
	Debug() string
}

const (
	ASSIGN      string = "assign"
	BLOCK       string = "block"
	BUILTIN     string = "builtin"
	CALL        string = "call"
	FLOAT       string = "float"
	FUNCTION    string = "function"
	IDENTIFIER  string = "identifier"
	IDENTIFIERS string = "identifiers"
	IF          string = "if"
	INDEX       string = "index"
	INFIX       string = "infix"
	INTEGER     string = "integer"
	LOOP        string = "loop"
	POSTFIX     string = "postfix"
	PREFIX      string = "prefix"
	PROGRAM     string = "program"
	RETURN      string = "return"
	STRING      string = "string"
	SWITCH      string = "switch"
	TERNARY     string = "ternary"
	THROW       string = "throw"
)
