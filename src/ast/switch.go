// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Switch struct {
	Token     types.Token
	Condition Expression
	Cases     []Case
	Else      *Block
}

func (s Switch) Value() string {
	return s.Token.Value
}

func (s Switch) String() string {
	var out bytes.Buffer
	out.WriteString("switch ")
	out.WriteString(s.Condition.String() + " {")
	for _, i := range s.Cases {
		out.WriteString(i.String() + ";")
	}
	if s.Else.Body != nil {
		out.WriteString("else ")
		out.WriteString(s.Else.String())
	}
	return out.String()
}
func (s Switch) statement() {}
