// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type If struct {
	Line        int
	Value       string
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

func (this If) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString(" ")
	out.WriteString(this.Condition.StringValue())
	out.WriteString(" ")
	out.WriteString(this.Consequence.StringValue())
	if this.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(this.Alternative.StringValue())
	}
	return out.String()
}

func (this If) GetLine() int {
	return this.Line
}

func (this If) statement() {}
