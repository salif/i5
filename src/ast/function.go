// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"strings"
)

type Function struct {
	line   uint32
	token  string
	params Identifiers
	body   Node
}

func (this Function) GetType() string {
	return FUNCTION
}

func (this Function) Debug() string {
	var result strings.Builder
	result.WriteString("((")
	result.WriteString(this.params.Debug())
	result.WriteString(") ")
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.body.Debug())
	result.WriteString(")")
	return result.String()
}

func (this Function) GetLine() uint32 {
	return this.line
}

func (this Function) Init(line uint32, token string) Function {
	this.line = line
	this.token = token
	return this
}

func (this Function) GetParams() Identifiers {
	return this.params
}

func (this *Function) SetParams(params Identifiers) {
	this.params = params
}

func (this Function) GetBody() Node {
	return this.body
}

func (this *Function) SetBody(body Node) {
	this.body = body
}
