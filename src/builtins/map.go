// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package builtins

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func _map(obj ...object.Object) object.Object {
	return object.Map{Value: make(map[object.Key]object.Object, 0)}
}

func map_clear(args ...object.Object) object.Object {
	return Null
}

func map_get(obj ...object.Object) object.Object {
	v := obj[0]
	key := obj[1]
	if v.Type() != constants.TYPE_MAP {
		return newException(constants.EXCEPTION_TYPE, constants.IR_IS_NOT_A_MAP, v.Type())
	}
	vmap := v.(object.Map)
	return vmap.Get(key)
}

func map_keys(args ...object.Object) object.Object {
	return Null
}

func map_remove(args ...object.Object) object.Object {
	return Null
}

func map_set(obj ...object.Object) object.Object {
	v := obj[0]
	key := obj[1]
	value := obj[2]

	if v.Type() != constants.TYPE_MAP {
		return newException(constants.EXCEPTION_TYPE, constants.IR_IS_NOT_A_MAP, v.Type())
	}

	vmap := v.(object.Map)
	r := vmap.Set(key, value)
	if r {
		return vmap
	} else {
		return newException(constants.EXCEPTION_TYPE, constants.IR_CANNOT_BE_KEY, key.Type())
	}
}

func map_values(args ...object.Object) object.Object {
	return Null
}
