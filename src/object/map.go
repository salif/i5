// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"bytes"

	"github.com/i5/i5/src/io/console"
)

type Map struct {
	Value map[string]Object
}

func (this Map) Type() TYPE {
	return MAP
}

func (this Map) StringValue() string {
	var out bytes.Buffer
	out.WriteString("{")
	for i, v := range this.Value {
		out.WriteString(console.Format("\"%v\": \"%v\", ", i, v.StringValue()))
	}
	out.WriteString("}")
	return out.String()
}

func (this *Map) Get(key string) Object {
	if value, ok := this.Value[key]; ok {
		return value
	} else {
		return &Void{}
	}
}

func (this *Map) Set(key string, value Object) {
	this.Value[key] = value
}
