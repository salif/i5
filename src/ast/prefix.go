// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Prefix struct {
	line     int
	operator string
	right    Node
}

func (this Prefix) GetType() int {
	return PREFIX
}

func (this Prefix) Print() {
	console.Print("(")
	console.Print(this.operator)
	this.right.Print()
	console.Print(")")
}

func (this Prefix) GetLine() int {
	return this.line
}

func (this Prefix) GetOperator() string {
	return this.operator
}

func (this Prefix) GetRight() Node {
	return this.right
}

func (this Prefix) Init(line int, operator string) Prefix {
	this.line = line
	this.operator = operator
	return this
}

func (this *Prefix) SetRight(right Node) {
	this.right = right
}
