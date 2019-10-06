// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Try struct {
	Value   string
	Body    *Block
	Err     *Identifier
	Catch   *Block
	Finally *Block
}

func (t Try) StringValue() string {
	var out bytes.Buffer
	out.WriteString(t.Value)
	out.WriteString(" ")
	out.WriteString(t.Body.StringValue())
	if t.Catch != nil {
		out.WriteString(" catch ")
		if t.Err != nil {
			out.WriteString(t.Err.StringValue())
			out.WriteString(" ")
		}
		out.WriteString(t.Catch.StringValue())
		if t.Finally != nil {
			out.WriteString(" finally ")
			out.WriteString(t.Finally.StringValue())
		}
	}
	return out.String()
}
func (t Try) statement() {}
