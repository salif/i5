// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package object

type Object interface {
	Type() string
	StringValue() string
}

type Key struct {
	Type  string
	Value Object
}

func (this Key) GetValue() Object {
	return this.Value
}

type Mappable interface {
	GenKey() Key
}
