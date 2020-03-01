// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import (
	"strings"

	"github.com/i5/i5/src/constants"
)

type Ternary struct {
	line        uint32
	token       string
	condition   Node
	consequence Node
	alternative Node
}

func (this Ternary) GetType() string {
	return TERNARY
}

func (this Ternary) Debug() string {
	var result strings.Builder
	result.WriteString(this.condition.Debug())
	result.WriteString(" ")
	result.WriteString(this.token)
	result.WriteString(" ")
	result.WriteString(this.consequence.Debug())
	if this.HaveAlternative() {
		result.WriteString(" ")
		result.WriteString(constants.TOKEN_QMQM)
		result.WriteString(" ")
		result.WriteString(this.alternative.Debug())
	}
	return result.String()
}

func (this Ternary) GetLine() uint32 {
	return this.line
}

func (this Ternary) GetCondition() Node {
	return this.condition
}

func (this Ternary) GetConsequence() Node {
	return this.consequence
}

func (this Ternary) GetAlternative() Node {
	return this.alternative
}

func (this Ternary) HaveAlternative() bool {
	return this.alternative != nil
}

func (this Ternary) Init(line uint32, token string) Ternary {
	this.line = line
	this.token = token
	return this
}

func (this *Ternary) SetCondition(condition Node) {
	this.condition = condition
}

func (this *Ternary) SetConsequence(consequence Node) {
	this.consequence = consequence
}

func (this *Ternary) SetAlternative(alternative Node) {
	this.alternative = alternative
}
