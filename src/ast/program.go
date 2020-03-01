// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"
)

type Program struct {
	line uint32
	body []Function
}

func (this Program) GetType() string {
	return PROGRAM
}

func (this Program) Debug() string {
	var result strings.Builder
	var n []string
	for _, s := range this.body {
		n = append(n, s.Debug())
	}
	result.WriteString(strings.Join(n, "; "))
	return result.String()
}

func (this Program) Init(line uint32) Program {
	this.line = line
	this.body = make([]Function, 0)
	return this
}

func (this Program) GetLine() uint32 {
	return this.line
}

func (this Program) GetBody() []Function {
	return this.body
}

func (this *Program) SetBody(body []Function) {
	this.body = body
}
