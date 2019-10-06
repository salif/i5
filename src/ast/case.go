// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Case struct {
	Cases []Expression
	Body  *Block
}

func (c Case) StringValue() string {
	var out bytes.Buffer
	for _, i := range c.Cases {
		out.WriteString("case " + i.StringValue() + ";")
	}
	out.WriteString(c.Body.StringValue())
	return out.String()
}
