// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Throw struct {
	Line  int
	Value string
	Body  Expression
}

func (this Throw) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString("(")
	out.WriteString(this.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (this Throw) GetLine() int {
	return this.Line
}

func (this Throw) statement() {}
