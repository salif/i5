// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func _i5() object.Object {
	v := _Map()
	v.Set(_String("author"), _String("Salif Mehmed"))
	v.Set(_String("version"), _String(constants.MINOR_VERSION))
	return v
}
