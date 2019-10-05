// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Nil struct {
}

func (n *Nil) Type() TYPE {
	return NIL
}

func (n *Nil) StringValue() string {
	return "nil"
}
