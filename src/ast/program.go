// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "bytes"

type Program struct {
	Line int
	Body []Expression
}

func (this Program) StringValue() string {
	var out bytes.Buffer
	for _, s := range this.Body {
		out.WriteString(s.StringValue())
		out.WriteString("; ")
	}
	return out.String()
}

func (this Program) GetLine() int {
	return this.Line
}
