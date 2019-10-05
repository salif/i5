// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Expr struct {
	Body Expression
}

func (ex Expr) String() string {
	if ex.Body != nil {
		return ex.Body.String()
	}
	return ""
}

func (ex Expr) statement() {}
