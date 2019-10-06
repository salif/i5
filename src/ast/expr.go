// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Expr struct {
	Body Expression
}

func (ex Expr) StringValue() string {
	if ex.Body != nil {
		return ex.Body.StringValue()
	}
	return ""
}

func (ex Expr) statement() {}
