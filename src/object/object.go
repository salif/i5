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
	VOID     = "undefined"
	INTEGER  = "integer"
	FLOAT    = "float"
	STRING   = "string"
	BOOL     = "bool"
	RETURN   = "return"
	BREAK    = "break"
	CONTINUE = "continue"
	THROW    = "throw"
	ERROR    = "error"
	FUNCTION = "function"
	BUILTIN  = "builtin"
	ARRAY    = "array"
	MAP      = "map"
	MODULE   = "module"
)
