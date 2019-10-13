// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"github.com/i5/i5/src/io/console"
)

type Switch struct {
	line      int
	token     string
	condition Node
	cases     []Case
	_else     Block
}

func (this Switch) GetType() int {
	return SWITCH
}

func (this Switch) Print() {
	console.Print(this.token)
	console.Print(" ")
	this.condition.Print()
	console.Print(" {")
	for _, i := range this.cases {
		i.Print()
		console.Print(";")
	}
	if this._else.body != nil {
		console.Print("else ")
		this._else.Print()
	}
}

func (this Switch) GetLine() int {
	return this.line
}

func (this Switch) Init(line int, token string) Switch {
	this.line = line
	this.token = token
	return this
}

func (this *Switch) SetCondition(condition Node) {
	this.condition = condition
}

func (this *Switch) SetCases(cases []Case) {
	this.cases = cases
}

func (this *Switch) SetElse(_else Block) {
	this._else = _else
}

type Case struct {
	line  int
	cases []Node
	body  Block
}

func (this Case) Print() {
	for _, i := range this.cases {
		console.Print("case ")
		i.Print()
		console.Print(";")
	}
	this.body.Print()
}

func (this Case) GetLine() int {
	return this.line
}

func (this Case) Init(line int) Case {
	this.line = line
	return this
}

func (this *Case) Append(_case Node) {
	this.cases = append(this.cases, _case)
}

func (this *Case) SetBody(body Block) {
	this.body = body
}
