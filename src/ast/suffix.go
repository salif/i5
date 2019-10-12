// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Suffix struct {
	Line     int
	Left     Expression
	Operator string
}

func (this Suffix) StringValue() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(this.Left.StringValue())
	out.WriteString(this.Operator)
	out.WriteString(")")
	return out.String()
}

func (this Suffix) GetLine() int {
	return this.Line
}

func (this Suffix) expression() {}
