// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"github.com/i5/i5/src/io/console"
)

type Call struct {
	line      int
	caller    Node
	arguments []Node
}

func (this Call) GetType() string {
	return CALL
}

func (this Call) Print() {
	this.caller.Print()
	console.Print("(")
	for _, a := range this.arguments {
		a.Print()
		console.Print(", ")
	}
	if len(this.arguments) > 0 {
		console.Print("\u0008\u0008")
	}
	console.Print(")")
}

func (this Call) GetLine() int {
	return this.line
}

func (this Call) Init(line int, caller Node) Call {
	this.line = line
	this.caller = caller
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
