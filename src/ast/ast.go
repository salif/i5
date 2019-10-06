// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Node interface {
	StringValue() string
}

type Statement interface {
	Node
	statement()
}

type Expression interface {
	Node
	expression()
}
