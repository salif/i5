// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
)

func Print(obj ...object.Object) object.Object {
	for _, o := range obj {
		console.Print(o.StringValue())
	}
	console.Println()
	return nil
}
