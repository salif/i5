// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func _array() object.Object {
	v := _Map()
	v.Set(_String("new"), _Builtin(object.ARRAY, 0, _array_new))
	v.Set(_String("add"), _Builtin(object.ARRAY, 2, _array_add))
	// v.Set(_String("remove"), _Builtin(object.ARRAY, 2, _array_remove))
	// v.Set(_String("get"), _Builtin(object.VOID, 2, _array_get))
	// v.Set(_String("set"), _Builtin(object.ARRAY, 3, _array_set))
	// v.Set(_String("join"), _Builtin(object.STRING, 2, _array_join))
	return v
}

func _array_new(args ...object.Object) object.Object {
	if len(args) > 0 {
	}
	return _Array()
}

func _array_add(args ...object.Object) object.Object {
	arr := args[0]
	elem := args[1]
	if arr.Type() == object.ARRAY {
		arr := arr.(object.Array)
		return object.Array{Value: arr.Push(elem)}
	} else {
		return _Void()
	}

}
