// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"bytes"
	"fmt"
	"strings"
)

type Array struct {
	Value []Object
}

func (this Array) Type() TYPE {
	return ARRAY
}

func (this Array) StringValue() string {
	return fmt.Sprintf("[type: %v]", this.Type())
}

func (this Array) ToString() String {
	var out bytes.Buffer
	out.WriteString("[")
	result := []string{}
	for _, v := range this.Value {
		if v.Type() == STRING {
			result = append(result, "\""+v.StringValue()+"\"")
		} else {
			result = append(result, v.StringValue())
		}
	}
	out.WriteString(strings.Join(result, ", "))
	out.WriteString("]")
	return String{Value: out.String()}
}

func (this Array) Init() Array {
	this.Value = make([]Object, 0)
	return this
}

func (this *Array) Push(obj Object) []Object {
	return append(this.Value, obj)
}
