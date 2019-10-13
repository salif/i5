// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
)

func _typeof(obj ...object.Object) object.Object {
	if len(obj) == 1 {
		return _String(console.Format("%v", obj[0].Type()))
	} else {
		return _Void()
	}
}
