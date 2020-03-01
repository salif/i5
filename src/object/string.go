// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package object

import "github.com/i5/i5/src/constants"

type String struct {
	Value string
}

func (this String) Type() string {
	return constants.TYPE_STRING
}

func (this String) StringValue() string {
	return this.Value
}

func (this String) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
