// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Node interface {
	GetType() int
	Print()
	GetLine() int
}

const (
	_ int = iota
	ASSIGN
	BLOCK
	BOOL
	BREAK
	BUILTIN
	CALL
	CONTINUE
	EXPRESSION
	FLOAT
	FUNCTION
	IDENTIFIER
	IF
	IMPORT
	INDEX
	INFIX
	INTEGER
	POSTFIX
	PREFIX
	PROGRAM
	RETURN
	STRING
	SWITCH
	THROW
	TRY
	WHILE
)
