// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Prefix struct {
	Operator string
	Right    Expression
}

func (p Prefix) StringValue() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(p.Operator)
	out.WriteString(p.Right.StringValue())
	out.WriteString(")")
	return out.String()
}
func (p Prefix) expression() {}
