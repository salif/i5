// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Throw struct {
	line  int
	token string
	body  Node
}

func (this Throw) GetType() string {
	return THROW
}

func (this Throw) Print() {
	console.Print(this.token)
	console.Print("(")
	this.body.Print()
	console.Print(")")
}

func (this Throw) GetLine() int {
	return this.line
}

func (this Throw) Init(line int, token string) Throw {
	this.line = line
	this.token = token
	return this
}

func (this Throw) GetBody() Node {
	return this.body
}

func (this *Throw) SetBody(body Node) {
	this.body = body
}
