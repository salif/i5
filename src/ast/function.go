// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
	"strings"
)

type Function struct {
	Value     string
	Anonymous bool
	Params    []*Identifier
	Body      *Block
}

func (f Function) StringValue() string {
	var out bytes.Buffer
	if f.Anonymous {
		out.WriteString(f.Value)
	}
	params := []string{}
	for _, p := range f.Params {
		params = append(params, p.StringValue())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, " "))
	out.WriteString(") ")
	out.WriteString(f.Body.StringValue())
	return out.String()
}

func (f Function) expression() {}
