// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Index struct {
	line     int
	left     Node
	operator string
	right    Node
}

func (this Index) GetType() string {
	return INDEX
}

func (this Index) Print() {
	console.Print("(")
	this.left.Print()
	console.Print(this.operator)
	this.right.Print()
	console.Print(")")
}

func (this Index) GetLine() int {
	return this.line
}

func (this Index) GetLeft() Node {
	return this.left
}

func (this Index) GetRight() Node {
	return this.right
}

func (this Index) GetOperator() string {
	return this.operator
}

func (this Index) Init(line int, left Node, operator string) Index {
	this.line = line
	this.left = left
	this.operator = operator
	return this
}

func (this *Index) SetRight(right Node) {
	this.right = right
}
