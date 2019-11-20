// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type BuiltinFunction struct {
	Function  func(args ...Object) Object
	MinParams int
}

func (this BuiltinFunction) Type() TYPE {
	return BUILTIN
}

func (this BuiltinFunction) StringValue() string {
	return fmt.Sprintf("[type: %v]", this.Type())
}
