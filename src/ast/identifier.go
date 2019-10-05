// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Identifier struct {
	Value string
}

func (i Identifier) String() string {
	return i.Value
}

func (i Identifier) expression() {}
