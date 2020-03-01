// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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

func (this Map) Type() string {
	return constants.TYPE_MAP
}

func (this Map) StringValue() string {
	var out bytes.Buffer
	out.WriteString("{")
	result := []string{}
	for i, v := range this.Value {
		left := i.GetValue().StringValue()
		right := v
		if i.GetValue().Type() == constants.TYPE_STRING {
			left = "\"" + left + "\""
		}
		if right.Type() == constants.TYPE_STRING {
			result = append(result, fmt.Sprintf("%v: \"%v\"", left, right.StringValue()))
		} else {
			result = append(result, fmt.Sprintf("%v: %v", left, right.StringValue()))
		}
	}
	out.WriteString(strings.Join(result, ", "))
	out.WriteString("}")
	return out.String()
}

func (this *Map) Get(key Object) Object {
	rkey := key.(Mappable)
	if value, ok := this.Value[rkey.GenKey()]; ok {
		return value
	} else {
		return Exception{Name: String{Value: constants.EXCEPTION_NULL}, Message: String{Value: constants.EXCEPTION_NULL}}
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
