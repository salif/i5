// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func _array() object.Object {
	v := _Map()
	v.Set("new", _Builtin(object.ARRAY, _array_new))
	v.Set("push", _Builtin(object.BOOL, _array_push))
	return v
}

func _array_new(obj ...object.Object) object.Object {
	return _Array()
}

func _array_push(obj ...object.Object) object.Object {
	if len(obj) == 2 {
		arr := obj[0]
		if arr.Type() == object.ARRAY {
			arr := arr.(*object.Array)
			return &object.Array{Value: arr.Push(obj[1])}
		} else {
			return _Void()
		}
	} else {
		return _Void()
	}
}
