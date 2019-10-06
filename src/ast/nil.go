// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Nil struct {
	Value string
}

func (n Nil) StringValue() string {
	return n.Value
}

func (n Nil) expression() {}
