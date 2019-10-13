// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type While struct {
	line      int
	token     string
	condition Node
	body      Block
}

func (this While) GetType() int {
	return WHILE
}

func (this While) Print() {
	console.Print(this.token)
	console.Print(" ")
	this.condition.Print()
	console.Print(" ")
	this.body.Print()
}

func (this While) GetLine() int {
	return this.line
}

func (this While) GetCondition() Node {
	return this.condition
}

func (this While) GetBody() Block {
	return this.body
}

func (this While) Init(line int, token string) While {
	this.line = line
	this.token = token
	return this
}

func (this *While) SetCondition(condition Node) {
	this.condition = condition
}

func (this *While) SetBody(body Block) {
	this.body = body
}
