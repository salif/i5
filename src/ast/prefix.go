// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Prefix struct {
	Line     int
	Operator string
	Right    Expression
}

func (this Prefix) StringValue() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(this.Operator)
	out.WriteString(this.Right.StringValue())
	out.WriteString(")")
	return out.String()
}

func (this Prefix) GetLine() int {
	return this.Line
}

func (this Prefix) expression() {}
