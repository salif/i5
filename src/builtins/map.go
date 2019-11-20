// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"fmt"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func _map() object.Object {
	v := object.ClassObject{}.Init()
	v.Set("new", object.BuiltinFunction{MinParams: 0, Function: _map_new})
	v.Set("get", object.BuiltinFunction{MinParams: 2, Function: _map_get})
	v.Set("set", object.BuiltinFunction{MinParams: 3, Function: _map_set})
	v.Set("toString", object.BuiltinFunction{MinParams: 1, Function: _map_to_string})
	return v
}

func _map_new(obj ...object.Object) object.Object {
	return object.Map{}.Init()
}

func _map_get(obj ...object.Object) object.Object {
	v := obj[0]
	key := obj[1]
	if v.Type() != object.MAP {
		return _Error(false, constants.ERROR_TYPE, fmt.Sprint(constants.IR_IS_NOT_A_MAP, v.Type()))
	}
	vmap := v.(object.Map)
	return vmap.Get(key)
}

func _map_set(obj ...object.Object) object.Object {
	v := obj[0]
	key := obj[1]
	value := obj[2]

	if v.Type() != object.MAP {
		return _Error(false, constants.ERROR_TYPE, fmt.Sprint(constants.IR_IS_NOT_A_MAP, v.Type()))
	}

	vmap := v.(object.Map)
	r := vmap.Set(key, value)
	if r {
		return vmap
	} else {
		return _Error(false, constants.ERROR_TYPE, fmt.Sprint(constants.IR_CANNOT_BE_KEY, key.Type()))
	}
}

func _map_to_string(args ...object.Object) object.Object {
	v := args[0]
	if v.Type() == object.MAP {
		m := v.(object.Map)
		return m.ToString()
	} else {
		return _Error(false, constants.ERROR_TYPE, fmt.Sprint(constants.IR_IS_NOT_A_MAP, v.Type()))
	}
}
