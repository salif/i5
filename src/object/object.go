// SPDX-License-Identifier: GPL-3.0-or-later
package object

type TYPE string

type Object interface {
	Type() TYPE
	StringValue() string
}

type Hashable interface {
	Hash() string
}

type Immutable interface {
	Clone() Object
}

const (
	NUMBER   = "number"
	STRING   = "string"
	BOOL     = "bool"
	NIL      = "nil"
	RETURN   = "return"
	ERROR    = "error"
	FUNCTION = "function"
	BUILTIN  = "builtin"
	ARRAY    = "array"
	MAP      = "map"
	MODULE   = "module"
)
