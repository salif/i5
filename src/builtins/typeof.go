// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"fmt"

	"github.com/i5/i5/src/object"
)

func _typeof(obj ...object.Object) object.Object {
	return object.String{Value: fmt.Sprintf("%v", obj[0].Type())}
}
