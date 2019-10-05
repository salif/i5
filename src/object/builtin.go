// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Builtin struct {
	Name     string
	RealType TYPE
	Function func(args ...Object) Object
	Value    Object
}

func (b *Builtin) Type() TYPE {
	return BUILTIN
}

func (b *Builtin) StringValue() string {
	return "[type: builtin]"
}
