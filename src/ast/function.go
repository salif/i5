// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
	"strings"

	"github.com/i5/i5/src/types"
)

type Function struct {
	Token     types.Token
	Anonymous bool
	Params    []*Identifier
	Body      *Block
	Return    Expression
	Strict    bool
}

func (f Function) Value() string {
	return f.Token.Value
}

func (f Function) String() string {
	var out bytes.Buffer
	if f.Anonymous {
		out.WriteString(f.Value())
	}
	params := []string{}
	for _, p := range f.Params {
		params = append(params, p.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	if f.Strict {
		out.WriteString(f.Return.String() + " ")
	}
	out.WriteString(f.Body.String())
	return out.String()
}

func (f Function) expression() {}
