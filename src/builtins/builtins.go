// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/object"
)

func Get(str string) (object.Object, bool) {
	builtin, ok := Builtins[str]
	return builtin, ok
}

var Builtins = map[string]*object.Builtin{
	"print": &object.Builtin{RealType: object.FUNCTION, Function: Print},
}
