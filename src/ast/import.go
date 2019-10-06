// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type ImportExpr struct {
	Value string
	Body  Expression
}

func (i ImportExpr) StringValue() string {
	var out bytes.Buffer
	out.WriteString(i.Value)
	out.WriteString("(")
	out.WriteString(i.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (i ImportExpr) expression() {}

type ImportStatement struct {
	Value string
	Body  Expression
}

func (is ImportStatement) StringValue() string {
	var out bytes.Buffer
	out.WriteString(is.Value)
	out.WriteString("(")
	out.WriteString(is.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (is ImportStatement) statement() {}
