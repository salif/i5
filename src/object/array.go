// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "bytes"

type Array struct {
	Value []Object
}

func (this Array) Type() TYPE {
	return ARRAY
}

func (this Array) StringValue() string {
	var out bytes.Buffer
	out.WriteString("[")
	for _, v := range this.Value {
		out.WriteString(v.StringValue())
		out.WriteString(", ")
	}
	out.WriteString("]")
	return out.String()
}

func (this *Array) Push(obj Object) []Object {
	return append(this.Value, obj)
}
