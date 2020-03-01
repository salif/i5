// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
