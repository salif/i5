// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Assign struct {
	Line  int
	Value string
	Left  Expression
	Right Expression
}

func (this Assign) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Left.StringValue())
	out.WriteString(" ")
	out.WriteString(this.Value)
	out.WriteString(" ")
	out.WriteString(this.Right.StringValue())
	return out.String()
}

func (this Assign) GetLine() int {
	return this.Line
}

func (this Assign) expression() {}
