// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Import struct {
	Value string
	Body  Expression
}

func (i Import) StringValue() string {
	var out bytes.Buffer
	out.WriteString(i.Value)
	out.WriteString("(")
	out.WriteString(i.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (i Import) expression() {}
