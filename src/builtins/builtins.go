// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func Get(str string, env *object.Env) (object.Object, bool) {
	r, ok := builtins[str]
	return r, ok
}

var builtins = map[string]object.Object{
	"null": Null,

	"true":  object.Boolean{Value: true},
	"false": object.Boolean{Value: false},

	"i5_version": object.String{Value: constants.MINOR_VERSION},

	"array":              object.BuiltinFunction{MinParams: 0, Function: array},
	"array_clear":        object.BuiltinFunction{MinParams: 1, Function: array_clear},
	"array_every":        object.BuiltinFunction{MinParams: 2, Function: array_every},
	"array_fill":         object.BuiltinFunction{MinParams: 4, Function: array_fill},
	"array_filter":       object.BuiltinFunction{MinParams: 2, Function: array_filter},
	"array_for_each":     object.BuiltinFunction{MinParams: 2, Function: array_for_each},
	"array_get":          object.BuiltinFunction{MinParams: 2, Function: array_get},
	"array_index":        object.BuiltinFunction{MinParams: 2, Function: array_index},
	"array_join":         object.BuiltinFunction{MinParams: 2, Function: array_join},
	"array_pop":          object.BuiltinFunction{MinParams: 1, Function: array_pop},
	"array_push":         object.BuiltinFunction{MinParams: 2, Function: array_push},
	"array_reduce":       object.BuiltinFunction{MinParams: 2, Function: array_reduce},
	"array_reduce_right": object.BuiltinFunction{MinParams: 2, Function: array_reduce_right},
	"array_reverse":      object.BuiltinFunction{MinParams: 1, Function: array_reverse},
	"array_set":          object.BuiltinFunction{MinParams: 3, Function: array_set},
	"array_shift":        object.BuiltinFunction{MinParams: 1, Function: array_shift},
	"array_slice":        object.BuiltinFunction{MinParams: 2, Function: array_slice},
	"array_sort":         object.BuiltinFunction{MinParams: 1, Function: array_sort},
	"boolean":            object.BuiltinFunction{MinParams: 1, Function: boolean},
	"console_read_line":  object.BuiltinFunction{MinParams: 0, Function: console_read_line},
	"console_write":      object.BuiltinFunction{MinParams: 1, Function: console_write},
	"console_write_line": object.BuiltinFunction{MinParams: 0, Function: console_write_line},
	"integer":            object.BuiltinFunction{MinParams: 1, Function: integer},
	"integer_parse":      object.BuiltinFunction{MinParams: 1, Function: integer_parse},
	"is_null":            object.BuiltinFunction{MinParams: 1, Function: is_null},
	"map":                object.BuiltinFunction{MinParams: 0, Function: _map},
	"map_clear":          object.BuiltinFunction{MinParams: 1, Function: map_clear},
	"map_get":            object.BuiltinFunction{MinParams: 2, Function: map_get},
	"map_keys":           object.BuiltinFunction{MinParams: 1, Function: map_keys},
	"map_remove":         object.BuiltinFunction{MinParams: 2, Function: map_remove},
	"map_set":            object.BuiltinFunction{MinParams: 3, Function: map_set},
	"map_values":         object.BuiltinFunction{MinParams: 1, Function: map_values},
	"print":              object.BuiltinFunction{MinParams: 0, Function: console_write_line},
	"printf":             object.BuiltinFunction{MinParams: 2, Function: printf},
	"string":             object.BuiltinFunction{MinParams: 1, Function: _string},
	"string_char":        object.BuiltinFunction{MinParams: 2, Function: string_char},
	"string_char_code":   object.BuiltinFunction{MinParams: 2, Function: string_char_code},
	"string_code":        object.BuiltinFunction{MinParams: 1, Function: string_code},
	"string_codes":       object.BuiltinFunction{MinParams: 1, Function: string_codes},
	"string_concat":      object.BuiltinFunction{MinParams: 1, Function: string_concat},
	"string_contains":    object.BuiltinFunction{MinParams: 2, Function: string_contains},
	"string_format":      object.BuiltinFunction{MinParams: 2, Function: string_format},
	"string_has_prefix":  object.BuiltinFunction{MinParams: 2, Function: string_has_prefix},
	"string_has_suffix":  object.BuiltinFunction{MinParams: 2, Function: string_has_suffix},
	"string_index":       object.BuiltinFunction{MinParams: 2, Function: string_index},
	"string_repeat":      object.BuiltinFunction{MinParams: 2, Function: string_repeat},
	"string_replace":     object.BuiltinFunction{MinParams: 3, Function: string_replace},
	"string_reverse":     object.BuiltinFunction{MinParams: 1, Function: string_reverse},
	"string_slice":       object.BuiltinFunction{MinParams: 2, Function: string_slice},
	"string_split":       object.BuiltinFunction{MinParams: 2, Function: string_split},
	"string_to_lower":    object.BuiltinFunction{MinParams: 1, Function: string_to_lower},
	"string_to_upper":    object.BuiltinFunction{MinParams: 1, Function: string_to_upper},
	"string_trim":        object.BuiltinFunction{MinParams: 1, Function: string_trim},
	"string_trim_left":   object.BuiltinFunction{MinParams: 1, Function: string_trim_left},
	"string_trim_right":  object.BuiltinFunction{MinParams: 1, Function: string_trim_right},
	"type_of":            object.BuiltinFunction{MinParams: 1, Function: type_of},
}

var Null = object.Exception{Name: object.String{Value: constants.EXCEPTION_NULL}, Message: object.String{Value: constants.EXCEPTION_NULL}}

func newException(name string, text string, format ...interface{}) object.Exception {
	return object.Exception{
		Name:    object.String{Value: name},
		Message: object.String{Value: fmt.Sprintf(text, format...)}}
}

func is_null(args ...object.Object) object.Object {
	return object.Boolean{Value: args[0] == Null}
}

func type_of(obj ...object.Object) object.Object {
	return object.String{Value: fmt.Sprintf("%v", obj[0].Type())}
}
