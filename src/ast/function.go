// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"github.com/i5/i5/src/io/console"
)

type Function struct {
	line      int
	token     string
	anonymous bool
	params    []Identifier
	body      Block
}

func (this Function) GetType() string {
	return FUNCTION
}

func (this Function) Print() {
	if this.anonymous {
		console.Print(this.token)
	}
	console.Print("(")
	for _, p := range this.params {
		p.Print()
		console.Print(" ")
	}

	console.Print(") ")
	this.body.Print()
}

func (this Function) GetLine() int {
	return this.line
}

func (this Function) Init(line int, token string) Function {
	this.line = line
	this.token = token
	return this
}

func (this Function) GetAnonymous() bool {
	return this.anonymous
}

func (this *Function) SetAnonymous(anonymous bool) {
	this.anonymous = anonymous
}

func (this Function) GetParams() []Identifier {
	return this.params
}

func (this *Function) SetParams(params []Identifier) {
	this.params = params
}

func (this Function) GetBody() Block {
	return this.body
}

func (this *Function) SetBody(body Block) {
	this.body = body
}
