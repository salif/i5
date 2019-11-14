// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func _array() object.Object {
	v := object.ClassObject{}.Init()
	v.Set("new", object.BuiltinFunction{MinParams: 0, Function: _array_new})
	v.Set("add", object.BuiltinFunction{MinParams: 2, Function: _array_add})
	// TODO add "remove", 2, _array_remove
	// TODO add "get", 2, _array_get
	// TODO add "set", 3, _array_set
	// TODO add "join", 2, _array_join
	return v
}

func _array_new(args ...object.Object) object.Object {
	if len(args) > 0 {
		// TODO
	}
	return object.Array{}.Init()
}

func _array_add(args ...object.Object) object.Object {
	arr := args[0]
	elem := args[1]
	if arr.Type() == object.ARRAY {
		arr := arr.(object.Array)
		return object.Array{Value: arr.Push(elem)}
	} else {
		return NIL
	}

}
