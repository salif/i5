// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func array(args ...object.Object) object.Object {
	result := object.Array{Value: make([]object.Object, 0)}
	if len(args) > 0 {
		// TODO
	}
	return result
}

func array_clear(args ...object.Object) object.Object {
	return Null
}

func array_every(args ...object.Object) object.Object {
	return Null
}

func array_fill(args ...object.Object) object.Object {
	return Null
}

func array_filter(args ...object.Object) object.Object {
	return Null
}

func array_for_each(args ...object.Object) object.Object {
	return Null
}

func array_get(args ...object.Object) object.Object {
	return Null
}

func array_index(args ...object.Object) object.Object {
	return Null
}

func array_join(args ...object.Object) object.Object {
	return Null
}

func array_pop(args ...object.Object) object.Object {
	return Null
}

func array_push(args ...object.Object) object.Object {
	arr := args[0]
	elem := args[1]
	if arr.Type() == constants.TYPE_ARRAY {
		arr := arr.(object.Array)
		return object.Array{Value: arr.Push(elem)}
	} else {
		return newException(constants.EXCEPTION_TYPE, constants.IR_IS_NOT_AN_ARRAY, arr.Type())
	}
}

func array_reduce(args ...object.Object) object.Object {
	return Null
}

func array_reduce_right(args ...object.Object) object.Object {
	return Null
}

func array_reverse(args ...object.Object) object.Object {
	return Null
}

func array_set(args ...object.Object) object.Object {
	return Null
}

func array_shift(args ...object.Object) object.Object {
	return Null
}

func array_slice(args ...object.Object) object.Object {
	return Null
}

func array_sort(args ...object.Object) object.Object {
	return Null
}
