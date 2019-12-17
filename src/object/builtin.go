// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type BuiltinFunction struct {
	Function  func(args ...Object) Object
	MinParams int
}

func (this BuiltinFunction) Type() string {
	return constants.TYPE_BUILTIN
}

func (this BuiltinFunction) StringValue() string {
	return fmt.Sprintf("[type: %v, min-params: %v]", this.Type(), this.MinParams)
}
