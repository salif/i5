// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Float struct {
	Value float64
}

func (this Float) Type() string {
	return constants.TYPE_FLOAT
}

func (this Float) StringValue() string {
	return fmt.Sprint(this.Value)
}

func (this Float) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
