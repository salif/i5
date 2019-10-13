// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Infix struct {
	line     int
	left     Node
	operator string
	right    Node
}

func (this Infix) GetType() int {
	return INFIX
}

func (this Infix) Print() {
	console.Print("(")
	this.left.Print()
	console.Print(" ")
	console.Print(this.operator)
	console.Print(" ")
	this.right.Print()
	console.Print(")")
}

func (this Infix) GetLine() int {
	return this.line
}

func (this Infix) GetLeft() Node {
	return this.left
}

func (this Infix) GetRight() Node {
	return this.right
}

func (this Infix) GetOperator() string {
	return this.operator
}

func (this Infix) Init(line int, operator string, left Node) Infix {
	this.line = line
	this.operator = operator
	this.left = left
	return this
}

func (this Infix) Set(line int, left Node, operator string, right Node) Infix {
	this.line = line
	this.left = left
	this.operator = operator
	this.right = right
	return this
}

func (this *Infix) SetRight(right Node) {
	this.right = right
}
