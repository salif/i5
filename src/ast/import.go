// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Import struct {
	Value string
	Body  Expression
}

func (i Import) String() string {
	var out bytes.Buffer
	out.WriteString(i.Value)
	out.WriteString("(")
	out.WriteString(i.Body.String())
	out.WriteString(")")
	return out.String()
}

func (i Import) expression() {}
