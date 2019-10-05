// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Import struct {
	Token types.Token
	Val   Expression
}

func (i Import) Value() string {
	return i.Token.Value
}

func (i Import) String() string {
	var out bytes.Buffer
	out.WriteString(i.Value())
	out.WriteString("(")
	out.WriteString(i.Val.String())
	out.WriteString(")")
	return out.String()
}

func (i Import) expression() {}
