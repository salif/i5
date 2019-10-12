// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type ImportExpr struct {
	Line  int
	Value string
	Body  Expression
}

func (this ImportExpr) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString("(")
	out.WriteString(this.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (this ImportExpr) GetLine() int {
	return this.Line
}

func (this ImportExpr) expression() {}

type ImportStatement struct {
	Line  int
	Value string
	Body  Expression
}

func (this ImportStatement) StringValue() string {
	var out bytes.Buffer
	out.WriteString(this.Value)
	out.WriteString("(")
	out.WriteString(this.Body.StringValue())
	out.WriteString(")")
	return out.String()
}

func (this ImportStatement) GetLine() int {
	return this.Line
}

func (this ImportStatement) statement() {}
