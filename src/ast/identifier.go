// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/types"

type Identifier struct {
	Token  types.Token
	Type   types.Token
	Strict bool
	Val    string
}

func (i Identifier) Value() string {
	return i.Token.Value
}

func (i Identifier) String() string {
	result := i.Val
	if i.Strict {
		result += " " + i.Type.Value
	}
	return result
}

func (i Identifier) expression() {}
