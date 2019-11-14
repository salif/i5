// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/i5/i5/src/constants"
)

type Map struct {
	Value map[Key]Object
}

func (this Map) Type() TYPE {
	return MAP
}

func (this Map) StringValue() string {
	var out bytes.Buffer
	out.WriteString("{")
	result := []string{}
	for i, v := range this.Value {
		left := i.GetValue().StringValue()
		right := v
		if i.GetValue().Type() == STRING {
			left = "\"" + left + "\""
		}
		if right.Type() == STRING {
			result = append(result, fmt.Sprintf("%v: \"%v\"", left, right.StringValue()))
		} else {
			result = append(result, fmt.Sprintf("%v: %v", left, right.StringValue()))
		}
	}
	out.WriteString(strings.Join(result, ", "))
	out.WriteString("}")
	return out.String()
}

func (this Map) Init() Map {
	this.Value = make(map[Key]Object, 0)
	return this
}

func (this *Map) Get(key Object) Object {
	rkey := key.(Mappable)
	if value, ok := this.Value[rkey.GenKey()]; ok {
		return value
	} else {
		return Error{IsFatal: false, Line: 0, Number: Integer{Value: constants.ERROR_NIL}, Message: String{Value: "nil"}}
	}
}

func (this *Map) Set(key Object, value Object) bool {
	rkey, ok := key.(Mappable)
	if !ok {
		return false
	} else {
		this.Value[rkey.GenKey()] = value
		return true
	}
}
