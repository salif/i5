// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Case struct {
	Line  int
	Cases []Expression
	Body  *Block
}

func (this Case) StringValue() string {
	var out bytes.Buffer
	for _, i := range this.Cases {
		out.WriteString("case ")
		out.WriteString(i.StringValue())
		out.WriteString(";")
	}
	out.WriteString(this.Body.StringValue())
	return out.String()
}

func (this Case) GetLine() int {
	return this.Line
}
