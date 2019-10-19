// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Loop struct {
	line  int
	token string
	body  Block
}

func (this Loop) GetType() string {
	return LOOP
}

func (this Loop) Print() {
	console.Print(this.token)
	console.Print(" ")
	this.body.Print()
}

func (this Loop) GetLine() int {
	return this.line
}

func (this Loop) GetBody() Block {
	return this.body
}

func (this Loop) Init(line int, token string) Loop {
	this.line = line
	this.token = token
	return this
}

func (this *Loop) SetBody(body Block) {
	this.body = body
}
