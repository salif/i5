// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Block struct {
	Line int
	Body []Statement
}

func (this Block) StringValue() string {
	var out bytes.Buffer
	out.WriteString("{")
	for _, s := range this.Body {
		out.WriteString(s.StringValue())
		out.WriteString("; ")
	}
	out.WriteString("}")
	return out.String()
}

func (this Block) GetLine() int {
	return this.Line
}

func (this Block) statement() {}
