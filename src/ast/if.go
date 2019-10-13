// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type If struct {
	line        int
	token       string
	condition   Node
	consequence Block
	alternative Block
}

func (this If) GetType() int {
	return IF
}

func (this If) Print() {
	console.Print(this.token)
	console.Print(" ")
	this.condition.Print()
	console.Print(" ")
	this.consequence.Print()
	if this.alternative.body != nil {
		console.Print(" else ")
		this.alternative.Print()
	}
}

func (this If) GetLine() int {
	return this.line
}

func (this If) GetCondition() Node {
	return this.condition
}

func (this If) GetConsequence() Block {
	return this.consequence
}

func (this If) GetAlternative() Block {
	return this.alternative
}

func (this If) HaveAlternative() bool {
	return this.alternative.body != nil
}

func (this If) Init(line int, token string) If {
	this.line = line
	this.token = token
	return this
}

func (this *If) SetCondition(condition Node) {
	this.condition = condition
}

func (this *If) SetConsequence(consequence Block) {
	this.consequence = consequence
}

func (this *If) SetAlternative(alternative Block) {
	this.alternative = alternative
}
