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

func (a Assign) String() string {
	var out bytes.Buffer
	out.WriteString(a.Left.String())
	out.WriteString(" ")
	out.WriteString(a.Value)
	out.WriteString(" ")
	out.WriteString(a.Right.String())
	return out.String()
}

func (a Assign) expression() {}
