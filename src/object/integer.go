// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Integer struct {
	Value int64
}

func (this Integer) Type() string {
	return constants.TYPE_INTEGER
}

func (this Integer) StringValue() string {
	return fmt.Sprint(this.Value)
}

func (this Integer) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
