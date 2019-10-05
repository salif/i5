// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/types"

type Expr struct {
	Token types.Token
	Expr  Expression
}

func (ex Expr) Value() string {
	return ex.Token.Value
}

func (ex Expr) String() string {
	if ex.Expr != nil {
		return ex.Expr.String()
	}
	return ""
}

func (ex Expr) statement() {}
