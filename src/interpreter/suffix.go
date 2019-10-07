// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import "github.com/i5/i5/src/object"

func evalSuffix(operator string, right object.Object) object.Object {
	switch operator {
	case "++", "--":
		return newError("not implemented yet: %s%s", operator, right.Type())
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}
