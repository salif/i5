// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/types"

type Identifier struct {
	Token types.Token
	Val   string
}

func (i Identifier) Value() string {
	return i.Token.Value
}

func (i Identifier) String() string {
	return i.Val
}

func (i Identifier) expression() {}
