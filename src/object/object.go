// SPDX-License-Identifier: GPL-3.0-or-later
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
