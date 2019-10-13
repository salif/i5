// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Import struct {
	line  int
	token string
	body  Node
}

func (this Import) GetType() int {
	return IMPORT
}

func (this Import) Print() {
	console.Print(this.token)
	console.Print("(")
	this.body.Print()
	console.Print(")")
}

func (this Import) GetLine() int {
	return this.line
}

func (this Import) Init(line int, token string) Import {
	this.line = line
	this.token = token
	return this
}

func (this *Import) SetBody(body Node) {
	this.body = body
}
