// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func _class() object.Object {
	v := object.ClassObject{}.Init()
	v.Set("new", object.BuiltinFunction{MinParams: 0, Function: _class_new})
	v.Set("toString", object.BuiltinFunction{MinParams: 1, Function: _class_to_string})
	return v

}

func _class_new(obj ...object.Object) object.Object {
	return object.ClassObject{}.Init()
}

func _class_to_string(args ...object.Object) object.Object {
	v := args[0]
	if v.Type() == object.CLASSOBJECT {
		c := v.(object.ClassObject)
		return c.ToString()
	} else {
		return _Error(false, constants.ERROR_TYPE, fmt.Sprint(constants.IR_IS_NOT_A_CLASS, v.Type()))
	}
}
