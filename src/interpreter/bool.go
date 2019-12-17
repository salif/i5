// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import "github.com/i5/i5/src/object"

var (
	TRUE  object.Boolean = object.Boolean{Value: true}
	FALSE object.Boolean = object.Boolean{Value: false}
)

func nativeToBool(input bool) object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func isTrue(obj object.Object) bool {
	if obj == TRUE {
		return true
	} else if obj == FALSE {
		return false
	} else {
		return false
	}
}
