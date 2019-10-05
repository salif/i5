// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
	"strings"

	"github.com/i5/i5/src/types"
)

type Call struct {
	Token     types.Token
	Function  Expression
	Arguments []Expression
}

func (c Call) Value() string {
	return c.Token.Value
}

func (c Call) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

func (c Call) expression() {}
