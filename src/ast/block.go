// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Block struct {
	Body []Statement
}

func (b Block) StringValue() string {
	var out bytes.Buffer
	out.WriteString("{")
	for _, s := range b.Body {
		out.WriteString(s.StringValue())
		out.WriteString("; ")
	}
	out.WriteString("}")
	return out.String()
}

func (b Block) statement() {}
