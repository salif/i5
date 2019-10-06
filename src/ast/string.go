// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type String struct {
	Value string
}

func (s String) StringValue() string {
	return "\"" + s.Value + "\""
}

func (s String) expression() {}
