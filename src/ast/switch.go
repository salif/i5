// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"fmt"
	"strings"

	"github.com/i5/i5/src/types"
)

type Switch struct {
	line      uint32
	token     string
	condition Node
	cases     []Case
	_else     Block
}

func (this Switch) GetType() string {
	return SWITCH
}

func (this Switch) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.condition.Debug())
	result.WriteString(" {")
	var n []string
	for _, i := range this.cases {
		n = append(n, i.Debug())
	}
	result.WriteString(strings.Join(n, ";"))
	if this._else.body != nil {
		result.WriteString(types.ELSE)
		result.WriteString(" ")
		result.WriteString(this._else.Debug())
	}
	return result.String()
}

func (this Switch) GetLine() uint32 {
	return this.line
}

func (this Switch) Init(line uint32, token string) Switch {
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
	line  uint32
	cases []Node
	body  Block
}

func (this Case) Debug() string {
	var result strings.Builder
	for _, i := range this.cases {
		result.WriteString(types.CASE)
		result.WriteString(" ")
		result.WriteString(i.Debug())
		fmt.Print(";")
	}
	result.WriteString(this.body.Debug())
	return result.String()
}

func (this Case) GetLine() uint32 {
	return this.line
}

func (this Case) Init(line uint32) Case {
	this.line = line
	return this
}

func (this *Case) Append(_case Node) {
	this.cases = append(this.cases, _case)
}

func (this *Case) SetBody(body Block) {
	this.body = body
}
