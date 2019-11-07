// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "strings"

type Throw struct {
	line  uint32
	token string
	body  Node
}

func (this Throw) GetType() string {
	return THROW
}

func (this Throw) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.body.Debug())
	return result.String()
}

func (this Throw) GetLine() uint32 {
	return this.line
}

func (this Throw) Init(line uint32, token string) Throw {
	this.line = line
	this.token = token
	return this
}

func (this Throw) GetBody() Node {
	return this.body
}

func (this *Throw) SetBody(body Node) {
	this.body = body
}
