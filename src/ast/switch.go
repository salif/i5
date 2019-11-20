// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"strings"

	"github.com/i5/i5/src/types"
)

type Switch struct {
	line      uint32
	token     string
	condition Node
	cases     []Case
}

func (this Switch) GetType() string {
	return SWITCH
}

func (this Switch) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.condition.Debug())
	result.WriteString("{")
	var n []string
	for _, i := range this.cases {
		n = append(n, i.Debug())
	}
	result.WriteString(strings.Join(n, "; "))
	result.WriteString("}")
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

func (this Switch) GetCondition() Node {
	return this.condition
}

func (this Switch) GetCases() []Case {
	return this.cases
}

func (this *Switch) SetCondition(condition Node) {
	this.condition = condition
}

func (this *Switch) SetCases(cases []Case) {
	this.cases = cases
}

type Case struct {
	line  uint32
	_case Node
	body  Block
}

func (this Case) Debug() string {
	var result strings.Builder
	result.WriteString(types.CASE)
	result.WriteString(" ")
	result.WriteString(this._case.Debug())
	result.WriteString(" => ")
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

func (this Case) GetCase() Node {
	return this._case
}

func (this Case) GetBody() Node {
	return this.body
}

func (this *Case) SetCase(_case Node) {
	this._case = _case
}

func (this *Case) SetBody(body Block) {
	this.body = body
}
