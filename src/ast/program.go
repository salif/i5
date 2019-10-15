// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Program struct {
	line int
	body []Node
}

func (this Program) GetType() string {
	return PROGRAM
}

func (this Program) Print() {
	for _, s := range this.body {
		s.Print()
		console.Print("; ")
	}
}

func (this Program) GetLine() int {
	return this.line
}

func (this Program) GetBody() []Node {
	return this.body
}

func (this Program) Init(line int, body []Node) Program {
	this.line = line
	this.body = body
	return this
}

func (this *Program) Append(node Node) {
	this.body = append(this.body, node)
}
