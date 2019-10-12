// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Expr struct {
	Line int
	Body Expression
}

func (this Expr) StringValue() string {
	if this.Body != nil {
		return this.Body.StringValue()
	}
	return ""
}

func (this Expr) GetLine() int {
	return this.Line
}

func (this Expr) statement() {}
