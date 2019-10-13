// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Builtin struct {
	RealType TYPE
	Function func(args ...Object) Object
	Value    Object
}

func (this Builtin) Type() TYPE {
	return BUILTIN
}

func (this Builtin) StringValue() string {
	return "[type: builtin]"
}
