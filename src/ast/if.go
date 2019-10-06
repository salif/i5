// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type If struct {
	Value       string
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

func (i If) StringValue() string {
	var out bytes.Buffer
	out.WriteString(i.Value)
	out.WriteString(" ")
	out.WriteString(i.Condition.StringValue())
	out.WriteString(" ")
	out.WriteString(i.Consequence.StringValue())
	if i.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(i.Alternative.StringValue())
	}
	return out.String()
}
func (i If) statement() {}
