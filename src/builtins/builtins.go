// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func Get(str string, env *object.Env) (object.Object, bool) {
	builtin, ok := builtins[str]
	if ok {
		if builtin.RealType == object.FUNCTION {
			return builtin, ok
		} else {
			return builtin.Value, ok
		}
	} else {
		return object.Void{}, ok
	}
}

var builtins = map[string]object.Builtin{
	"null": object.Builtin{RealType: object.VOID, Value: _Void()},
	// Bools
	// Strings
	// Arrays
	// Maps
	"array":   object.Builtin{RealType: object.MAP, Value: _array()},
	"i5":      object.Builtin{RealType: object.MAP, Value: _i5()},
	"map":     object.Builtin{RealType: object.MAP, Value: _map()},
	"my":      object.Builtin{RealType: object.MAP, Value: _my()},
	"console": object.Builtin{RealType: object.MAP, Value: _console()},
	// Functions
	"print":  object.Builtin{RealType: object.FUNCTION, Function: _console_println, MinParams: 0},
	"typeof": object.Builtin{RealType: object.FUNCTION, Function: _typeof, MinParams: 1},
}

func _Void() object.Object {
	return object.Void{}
}

func _Bool(b bool) object.Bool {
	return object.Bool{Value: b}
}

func _String(str string) object.String {
	return object.String{Value: str}
}

func _Builtin(t object.TYPE, minParams int, function func(args ...object.Object) object.Object) object.Builtin {
	return object.Builtin{RealType: t, Function: function, MinParams: minParams}
}

func _Array() object.Array {
	return object.Array{Value: make([]object.Object, 0)}
}

func _Map() object.Map {
	return object.Map{Value: make(map[object.Key]object.Object, 0)}
}

func _Error(e string) object.Error {
	return object.Error{Message: e}
}
