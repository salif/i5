// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func Get(str string, env *object.Env) (object.Object, bool) {
	builtin, ok := builtins[str]
	if ok {
		return builtin, ok
	} else {
		return NIL, ok
	}
}

var builtins = map[string]object.Object{
	"nil":     NIL,
	"break":   _Error(false, constants.ERROR_BREAK, "break"),
	"true":    object.Bool{Value: true},
	"false":   object.Bool{Value: false},
	"array":   _array(),
	"error":   _error(),
	"i5":      _i5(),
	"map":     _map(),
	"console": _console(),
	"print":   object.BuiltinFunction{Function: _console_println, MinParams: 0},
	"typeof":  object.BuiltinFunction{Function: _typeof, MinParams: 1},
}

var NIL = object.Error{IsFatal: false, Line: 0, Number: object.Integer{Value: constants.ERROR_NIL}, Message: object.String{Value: "nil"}}

func _Error(isFatal bool, number int64, message string) object.Error {
	return object.Error{IsFatal: isFatal, Line: uint32(0), Number: object.Integer{Value: number}, Message: object.String{Value: message}}
}
