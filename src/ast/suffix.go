// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Suffix struct {
	Left     Expression
	Operator string
}

func (s Suffix) StringValue() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(s.Left.StringValue())
	out.WriteString(s.Operator)
	out.WriteString(")")
	return out.String()
}
func (s Suffix) expression() {}
