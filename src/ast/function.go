// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"
)

type Function struct {
	line   uint32
	token  string
	name   Identifier
	params []Identifier
	body   Block
}

func (this Function) GetType() string {
	return FUNCTION
}

func (this Function) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.name.Debug())
	result.WriteString("(")
	var n []string = make([]string, 0)
	for _, param := range this.params {
		n = append(n, param.Debug())
	}
	result.WriteString(strings.Join(n, " "))
	result.WriteString(") ")
	result.WriteString(this.body.Debug())
	return result.String()
}

func (this Function) GetLine() uint32 {
	return this.line
}

func (this Function) Init(line uint32, token string) Function {
	this.line = line
	this.token = token
	this.params = make([]Identifier, 0)
	return this
}

func (this Function) GetName() Identifier {
	return this.name
}

func (this *Function) SetName(name Identifier) {
	this.name = name
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
