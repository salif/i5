// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
	"strings"
)

type ExprList struct {
	Value string
	Body  []Expression
}

func (e ExprList) StringValue() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range e.Body {
		args = append(args, a.StringValue())
	}
	out.WriteString(strings.Join(args, ", "))
	return out.String()
}

func (e ExprList) expression() {}
