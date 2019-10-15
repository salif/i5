// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Return struct {
	line  int
	token string
	body  Node
}

func (this Return) GetType() string {
	return RETURN
}

func (this Return) Print() {
	console.Print(this.token)
	console.Print(" ")
	if this.body != nil {
		this.body.Print()
	}
}

func (this Return) GetLine() int {
	return this.line
}

func (this Return) GetBody() Node {
	return this.body
}

func (this Return) Init(line int, token string) Return {
	this.line = line
	this.token = token
	return this
}

func (this *Return) SetBody(body Node) {
	this.body = body
}
