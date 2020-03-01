// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package object

import (
	"bytes"
	"strings"

	"github.com/i5/i5/src/constants"
)

type Array struct {
	Value []Object
}

func (this Array) Type() string {
	return constants.TYPE_ARRAY
}

func (this Array) StringValue() string {
	var out bytes.Buffer
	out.WriteString("[")
	result := []string{}
	for _, v := range this.Value {
		if v.Type() == constants.TYPE_STRING {
			result = append(result, "\""+v.StringValue()+"\"")
		} else {
			result = append(result, v.StringValue())
		}
	}
	out.WriteString(strings.Join(result, ", "))
	out.WriteString("]")
	return out.String()
}

func (this *Array) Push(obj Object) []Object {
	return append(this.Value, obj)
}
