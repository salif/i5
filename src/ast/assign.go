// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Assign struct {
	Token types.Token
	Left  Expression
	Right Expression
}

func (a Assign) Value() string {
	return a.Token.Value
}

func (a Assign) String() string {
	var out bytes.Buffer
	switch a.Right.(type) {
	case Function:
		out.WriteString(a.Right.Value() + " " + a.Left.String() + a.Right.String())
	default:
		out.WriteString(a.Left.String())
		out.WriteString(a.Value())
		out.WriteString(a.Right.String())
	}
	return out.String()
}

func (a Assign) expression() {}
