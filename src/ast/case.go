// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Case struct {
	Cases []Expression
	Body  *Block
}

func (c Case) String() string {
	var out bytes.Buffer
	for _, i := range c.Cases {
		out.WriteString("case " + i.String() + ";")
	}
	out.WriteString(c.Body.String())
	return out.String()
}
