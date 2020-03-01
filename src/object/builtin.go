// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
