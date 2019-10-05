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

func (i Infix) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(" " + i.Operator + " ")
	out.WriteString(i.Right.String())
	out.WriteString(")")
	return out.String()
}

func (i Infix) expression() {}
