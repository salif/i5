// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "bytes"

type Array struct {
	Elements []Object
}

func (a *Array) Type() TYPE {
	return ARRAY
}

func (a *Array) StringValue() string {
	var out bytes.Buffer
	out.WriteString("[")
	for _, v := range a.Elements {
		out.WriteString(v.StringValue())
	}
	out.WriteString("]")
	return out.String()
}
