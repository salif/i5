// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Expression struct {
	line int
	body Node
}

func (this Expression) GetType() int {
	return EXPRESSION
}

func (this Expression) Print() {
	this.body.Print()
}

func (this Expression) GetLine() int {
	return this.line
}

func (this Expression) GetBody() Node {
	return this.body
}

func (this Expression) Init(line int, body Node) Expression {
	this.line = line
	this.body = body
	return this
}
