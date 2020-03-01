// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import "strings"

type Loop struct {
	line  uint32
	token string
	body  Block
}

func (this Loop) GetType() string {
	return LOOP
}

func (this Loop) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.body.Debug())
	return result.String()
}

func (this Loop) GetLine() uint32 {
	return this.line
}

func (this Loop) GetBody() Block {
	return this.body
}

func (this Loop) Init(line uint32, token string) Loop {
	this.line = line
	this.token = token
	return this
}

func (this *Loop) SetBody(body Block) {
	this.body = body
}
