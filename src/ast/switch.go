// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Switch struct {
	Line      int
	Value     string
	Condition Expression
	Cases     []Case
	Else      *Block
}

func (this Switch) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString(" ")
	out.WriteString(this.Condition.StringValue())
	out.WriteString(" {")
	for _, i := range this.Cases {
		out.WriteString(i.StringValue())
		out.WriteString(";")
	}
	if this.Else.Body != nil {
		out.WriteString("else ")
		out.WriteString(this.Else.StringValue())
	}
	return out.String()
}
func (this Switch) GetLine() int {
	return this.Line
}

func (this Switch) statement() {}
