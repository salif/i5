// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Block struct {
	line int
	body []Node
}

func (this Block) GetType() int {
	return BLOCK
}

func (this Block) Print() {
	console.Print("{")
	for _, s := range this.body {
		s.Print()
		console.Print("; ")
	}
	console.Print("}")
}

func (this Block) GetLine() int {
	return this.line
}

func (this Block) GetBody() []Node {
	return this.body
}

func (this Block) Init(line int) Block {
	this.line = line
	return this
}

func (this Block) Set(line int, body []Node) Block {
	this.line = line
	this.body = body
	return this
}

func (this *Block) Append(node Node) {
	this.body = append(this.body, node)
}
