// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Postfix struct {
	line     int
	left     Node
	operator string
}

func (this Postfix) GetType() int {
	return POSTFIX
}

func (this Postfix) Print() {
	console.Print("(")
	this.left.Print()
	console.Print(this.operator)
	console.Print(")")
}

func (this Postfix) GetLine() int {
	return this.line
}

func (this Postfix) GetLeft() Node {
	return this.left
}

func (this Postfix) GetOperator() string {
	return this.operator
}

func (this Postfix) Init(line int, operator string, left Node) Postfix {
	this.line = line
	this.operator = operator
	this.left = left
	return this
}
