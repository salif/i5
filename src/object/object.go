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
	INTEGER  TYPE = "integer"
	FLOAT    TYPE = "float"
	STRING   TYPE = "string"
	BOOL     TYPE = "bool"
	ERROR    TYPE = "error"
	FUNCTION TYPE = "function"
	BUILTIN  TYPE = "builtin"
	ARRAY    TYPE = "array"
	MAP      TYPE = "map"
	RETURN   TYPE = "return"
	THROW    TYPE = "throw"
)
