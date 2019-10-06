// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Infix struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (i Infix) StringValue() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.Left.StringValue())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.StringValue())
	out.WriteString(")")
	return out.String()
}

func (i Infix) expression() {}
