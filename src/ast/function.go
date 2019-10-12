// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
	"strings"
)

type Function struct {
	Line      int
	Value     string
	Anonymous bool
	Params    []*Identifier
	Body      *Block
}

func (this Function) StringValue() string {
	var out bytes.Buffer
	if this.Anonymous {
		out.WriteString(this.Value)
	}
	params := []string{}
	for _, p := range this.Params {
		params = append(params, p.StringValue())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, " "))
	out.WriteString(") ")
	out.WriteString(this.Body.StringValue())
	return out.String()
}

func (this Function) GetLine() int {
	return this.Line
}

func (this Function) expression() {}
