// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"fmt"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func _array() object.Object {
	v := object.ClassObject{}.Init()
	v.Set("new", object.BuiltinFunction{MinParams: 0, Function: _array_new})
	v.Set("add", object.BuiltinFunction{MinParams: 2, Function: _array_add})
	v.Set("toString", object.BuiltinFunction{MinParams: 1, Function: _array_to_string})
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
		return _Error(false, constants.ERROR_TYPE, fmt.Sprint(constants.IR_IS_NOT_AN_ARRAY, arr.Type()))
	}
}

func _array_to_string(args ...object.Object) object.Object {
	v := args[0]
	if v.Type() == object.ARRAY {
		arr := v.(object.Array)
		return arr.ToString()
	} else {
		return _Error(false, constants.ERROR_TYPE, fmt.Sprint(constants.IR_IS_NOT_AN_ARRAY, v.Type()))
	}
}
