// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Assign struct {
	Value string
	Left  Expression
	Right Expression
}

func (a Assign) StringValue() string {
	var out bytes.Buffer
	out.WriteString(a.Left.StringValue())
	out.WriteString(" ")
	out.WriteString(a.Value)
	out.WriteString(" ")
	out.WriteString(a.Right.StringValue())
	return out.String()
}

func (a Assign) expression() {}
