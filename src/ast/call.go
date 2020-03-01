// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"
)

type Call struct {
	line      uint32
	caller    Node
	arguments []Node
}

func (this Call) GetType() string {
	return CALL
}

func (this Call) Debug() string {
	var result strings.Builder
	result.WriteString(this.caller.Debug())
	result.WriteString("(")
	var n []string = make([]string, 0)
	for _, arg := range this.arguments {
		n = append(n, arg.Debug())
	}
	result.WriteString(strings.Join(n, ", "))
	result.WriteString(")")
	return result.String()
}

func (this Call) GetLine() uint32 {
	return this.line
}

func (this Call) Init(line uint32, caller Node) Call {
	this.line = line
	this.caller = caller
	this.arguments = make([]Node, 0)
	return this
}

func (this Call) GetCaller() Node {
	return this.caller
}

func (this Call) GetArguments() []Node {
	return this.arguments
}

func (this *Call) SetArguments(arguments []Node) {
	this.arguments = arguments
}
