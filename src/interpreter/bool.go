// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
