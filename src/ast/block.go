// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "strings"

type Block struct {
	line uint32
	body []Node
}

func (this Block) GetType() string {
	return BLOCK
}

func (this Block) Debug() string {
	var result strings.Builder
	result.WriteString("{")
	var n []string = make([]string, 0)
	for _, node := range this.body {
		n = append(n, node.Debug())
	}
	result.WriteString(strings.Join(n, "; "))
	result.WriteString("}")
	return result.String()
}

func (this Block) Init(line uint32) Block {
	this.line = line
	this.body = make([]Node, 0)
	return this
}

func (this Block) GetLine() uint32 {
	return this.line
}

func (this Block) GetBody() []Node {
	return this.body
}

func (this Block) Set(line uint32, body []Node) Block {
	this.line = line
	this.body = body
	return this
}

func (this *Block) SetBody(body []Node) {
	this.body = body
}
