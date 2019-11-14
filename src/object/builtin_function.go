// SPDX-License-Identifier: GPL-3.0-or-later
package object

type BuiltinFunction struct {
	Function  func(args ...Object) Object
	MinParams int
}

func (this BuiltinFunction) Type() TYPE {
	return BUILTIN
}

func (this BuiltinFunction) StringValue() string {
	return "[type: builtin]"
}
