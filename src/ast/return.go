// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"
)

type Return struct {
	line  uint32
	token string
	body  Node
}

func (this Return) GetType() string {
	return RETURN
}

func (this Return) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.body.Debug())
	return result.String()
}

func (this Return) GetLine() uint32 {
	return this.line
}

func (this Return) GetBody() Node {
	return this.body
}

func (this Return) Init(line uint32, token string) Return {
	this.line = line
	this.token = token
	return this
}

func (this *Return) SetBody(body Node) {
	this.body = body
}
