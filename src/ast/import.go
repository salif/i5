// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"strings"
)

type Import struct {
	line  uint32
	token string
	body  Node
	as Identifier
}

func (this Import) GetType() string {
	return IMPORT
}

func (this Import) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.body.Debug())
	return result.String()
}

func (this Import) GetLine() uint32 {
	return this.line
}

func (this Import) GetBody() Node {
	return this.body
}

func (this Import) Init(line uint32, token string) Import {
	this.line = line
	this.token = token
	return this
}

func (this *Import) SetBody(body Node) {
	this.body = body
}

func (this *Import) SetAs(body Node) {
	this.body = body
}
