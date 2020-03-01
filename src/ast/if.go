// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"

	"github.com/i5/i5/src/constants"
)

type If struct {
	line        uint32
	token       string
	condition   Node
	consequence Block
	alternative Block
}

func (this If) GetType() string {
	return IF
}

func (this If) Debug() string {
	var result strings.Builder
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.condition.Debug())
	result.WriteString(" ")
	result.WriteString(this.consequence.Debug())
	if this.HaveAlternative() {
		result.WriteString(" ")
		result.WriteString(constants.TOKEN_ELSE)
		result.WriteString(" ")
		result.WriteString(this.alternative.Debug())
	}
	return result.String()
}

func (this If) GetLine() uint32 {
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

func (this If) Init(line uint32, token string) If {
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
