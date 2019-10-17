// SPDX-License-Identifier: GPL-3.0-or-later
package object

type TYPE string

type Object interface {
	Type() TYPE
	StringValue() string
}

type Mappable interface {
	GenKey() Key
}

const (
	VOID     = "null"
	INTEGER  = "integer"
	FLOAT    = "float"
	STRING   = "string"
	BOOL     = "bool"
	RETURN   = "return"
	THROW    = "throw"
	ERROR    = "error"
	FUNCTION = "function"
	BUILTIN  = "builtin"
	ARRAY    = "array"
	MAP      = "map"
	MODULE   = "module"
)
