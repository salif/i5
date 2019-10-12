// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Return struct {
	Line  int
	Value string
	Body  Expression
}

func (this Return) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString(" ")
	if this.Body != nil {
		out.WriteString(this.Body.StringValue())
	}
	return out.String()
}

func (this Return) GetLine() int {
	return this.Line
}

func (this Return) statement() {}
