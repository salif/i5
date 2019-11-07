// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func _error() object.Object {
	v := _Map()
	v.Set(_String("new"), _Builtin(object.MAP, 1, _error_new))
	return v
}

func _error_new(obj ...object.Object) object.Object {
	v := _Map()
	v.Set(_String("message"), obj[0])
	return v
}
