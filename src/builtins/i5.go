// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import "github.com/i5/i5/src/object"

func _i5() object.Object {
	v := _Map()
	v.Set("author", _String("Salif Mehmed"))
	v.Set("github", _String("https://github.com/i5/i5.git"))
	v.Set("version", _version())
	return v
}
