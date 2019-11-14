// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func _i5() object.Object {
	v := object.ClassObject{}.Init()
	v.Set("author", object.String{Value: "Salif Mehmed"})
	v.Set("version", object.String{Value: constants.MINOR_VERSION})
	return v
}
