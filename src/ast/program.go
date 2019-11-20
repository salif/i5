// SPDX-License-Identifier: GPL-3.0-or-later
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
