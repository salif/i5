// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type While struct {
	Line      int
	Value     string
	Condition Expression
	Body      *Block
}

func (this While) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString(" ")
	out.WriteString(this.Condition.StringValue())
	out.WriteString(" ")
	out.WriteString(this.Body.StringValue())
	return out.String()
}

func (this While) GetLine() int {
	return this.Line
}

func (this While) statement() {}
