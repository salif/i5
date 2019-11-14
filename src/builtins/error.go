// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func _error() object.Object {
	v := object.ClassObject{}.Init()
	v.Set("new", object.BuiltinFunction{MinParams: 0, Function: _error_new})
	return v
}

// TODO
func _error_new(obj ...object.Object) object.Object {
	v := object.Error{IsFatal: false}
	if len(obj) > 0 {
		v.Message = obj[0].(object.String)
	}
	if len(obj) > 1 {
		v.Number = obj[1].(object.Integer)
	}
	return v
}
