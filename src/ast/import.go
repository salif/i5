// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Import struct {
	Line  int
	Value string
	Body  Expression
}

func (this Import) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString("(")
	out.WriteString(this.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (this Import) GetLine() int {
	return this.Line
}

func (this Import) expression() {}
