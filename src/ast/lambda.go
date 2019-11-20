// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"strings"
)

type Lambda struct {
	line   uint32
	token  string
	params []Identifier
	body   Node
}

func (this Lambda) GetType() string {
	return LAMBDA
}

func (this Lambda) Debug() string {
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

func (this Lambda) GetLine() uint32 {
	return this.line
}

func (this Lambda) Init(line uint32, token string) Lambda {
	this.line = line
	this.token = token
	this.params = make([]Identifier, 0)
	return this
}

func (this Lambda) GetParams() []Identifier {
	return this.params
}

func (this *Lambda) SetParams(params []Identifier) {
	this.params = params
}

func (this Lambda) GetBody() Node {
	return this.body
}

func (this *Lambda) SetBody(body Node) {
	this.body = body
}
