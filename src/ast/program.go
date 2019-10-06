// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "bytes"

type Program struct {
	Body []Expression
}

func (p Program) StringValue() string {
	var out bytes.Buffer
	for _, s := range p.Body {
		out.WriteString(s.StringValue() + "; ")
	}
	return out.String()
}
