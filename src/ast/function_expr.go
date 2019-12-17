// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"strings"
)

type FunctionExpr struct {
	line   uint32
	token  string
	params []Identifier
	body   Node
}

func (this FunctionExpr) GetType() string {
	return FUNCTIONEXPR
}

func (this FunctionExpr) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(": (")
	var n []string = make([]string, 0)
	for _, param := range this.params {
		n = append(n, param.Debug())
	}
	result.WriteString(strings.Join(n, " "))
	result.WriteString(") => ")
	result.WriteString(this.body.Debug())
	return result.String()
}

func (this FunctionExpr) GetLine() uint32 {
	return this.line
}

func (this FunctionExpr) Init(line uint32, token string) FunctionExpr {
	this.line = line
	this.token = token
	this.params = make([]Identifier, 0)
	return this
}

func (this FunctionExpr) GetParams() []Identifier {
	return this.params
}

func (this *FunctionExpr) SetParams(params []Identifier) {
	this.params = params
}

func (this FunctionExpr) GetBody() Node {
	return this.body
}

func (this *FunctionExpr) SetBody(body Node) {
	this.body = body
}
