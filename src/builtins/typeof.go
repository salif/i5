// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"fmt"

	"github.com/i5/i5/src/object"
)

func _typeof(obj ...object.Object) object.Object {
	elem := obj[0]
	return _String(fmt.Sprintf("%v", elem.Type()))
}
