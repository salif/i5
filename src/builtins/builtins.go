// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func Get(str string, env *object.Env) (object.Object, bool) {
	builtin, ok := Builtins[str]
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

var Builtins = map[string]object.Builtin{
	// Strings
	"version": object.Builtin{RealType: object.STRING, Value: _version()},

	// Maps
	"array": object.Builtin{RealType: object.MAP, Value: _array()},
	"i5":    object.Builtin{RealType: object.MAP, Value: _i5()},
	"map":   object.Builtin{RealType: object.MAP, Value: _map()},

	// Functions
	"print":  object.Builtin{RealType: object.FUNCTION, Function: _print},
	"typeof": object.Builtin{RealType: object.FUNCTION, Function: _typeof},
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

func _Builtin(t object.TYPE, f func(args ...object.Object) object.Object) object.Builtin {
	return object.Builtin{RealType: t, Function: f}
}

func _Array() object.Array {
	return object.Array{Value: make([]object.Object, 0)}
}

func _Map() object.Map {
	return object.Map{Value: make(map[string]object.Object, 0)}
}
