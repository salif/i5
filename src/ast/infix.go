// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Infix struct {
	Line     int
	Left     Expression
	Operator string
	Right    Expression
}

func (this Infix) StringValue() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(this.Left.StringValue())
	out.WriteString(" ")
	out.WriteString(this.Operator)
	out.WriteString(" ")
	out.WriteString(this.Right.StringValue())
	out.WriteString(")")
	return out.String()
}

func (this Infix) GetLine() int {
	return this.Line
}

func (this Infix) expression() {}
