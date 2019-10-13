// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Assign struct {
	line     int
	operator string
	left     Node
	right    Node
}

func (this Assign) GetType() int {
	return ASSIGN
}

func (this Assign) Print() {
	this.left.Print()
	console.Print(" ")
	console.Print(this.operator)
	console.Print(" ")
	this.right.Print()
}

func (this Assign) GetLine() int {
	return this.line
}

func (this Assign) Init(line int, operator string, left Node) Assign {
	this.line = line
	this.operator = operator
	this.left = left
	return this
}

func (this Assign) GetLeft() Node {
	return this.left
}

func (this Assign) GetRight() Node {
	return this.right
}

func (this *Assign) SetRight(right Node) {
	this.right = right
}
