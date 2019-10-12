// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Try struct {
	Line    int
	Value   string
	Body    *Block
	Err     *Identifier
	Catch   *Block
	Finally *Block
}

func (this Try) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString(" ")
	out.WriteString(this.Body.StringValue())
	if this.Catch != nil {
		out.WriteString(" catch ")
		if this.Err != nil {
			out.WriteString(this.Err.StringValue())
			out.WriteString(" ")
		}
		out.WriteString(this.Catch.StringValue())
		if this.Finally != nil {
			out.WriteString(" finally ")
			out.WriteString(this.Finally.StringValue())
		}
	}
	return out.String()
}

func (this Try) GetLine() int {
	return this.Line
}

func (this Try) statement() {}
