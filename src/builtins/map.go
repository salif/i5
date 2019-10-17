// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func _map() object.Object {
	v := _Map()
	v.Set(_String("new"), _Builtin(object.MAP, 0, _map_new))
	return v
}

func _map_new(obj ...object.Object) object.Object {
	return _Map()
}
