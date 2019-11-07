// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "strings"

type Postfix struct {
	line     uint32
	left     Node
	operator string
}

func (this Postfix) GetType() string {
	return POSTFIX
}

func (this Postfix) Debug() string {
	var result strings.Builder
	result.WriteString("(")
	result.WriteString(this.left.Debug())
	result.WriteString(this.operator)
	result.WriteString(")")
	return result.String()
}

func (this Postfix) GetLine() uint32 {
	return this.line
}

func (this Postfix) GetLeft() Node {
	return this.left
}

func (this Postfix) GetOperator() string {
	return this.operator
}

func (this Postfix) Init(line uint32, operator string, left Node) Postfix {
	this.line = line
	this.operator = operator
	this.left = left
	return this
}
