// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Throw struct {
	Value string
	Body  Expression
}

func (t Throw) StringValue() string {
	var out bytes.Buffer
	out.WriteString(t.Value)
	out.WriteString("(")
	out.WriteString(t.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (t Throw) statement() {}
