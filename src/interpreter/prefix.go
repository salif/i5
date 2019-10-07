// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import "github.com/i5/i5/src/object"

func evalPrefix(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		if right.Type() == object.BOOL {
			switch right {
			case TRUE:
				return FALSE
			case FALSE:
				return TRUE
			case NIL:
				return TRUE
			default:
				return FALSE
			}
		} else {
			return newError("unknown operator: %s%s", operator, right.Type())
		}
	case "~", "-":
		value := right.(*object.Number).Value
		switch operator {
		case "!":
			return FALSE
		case "~":
			return &object.Number{Value: ^value}
		case "-":
			return &object.Number{Value: -value}
		default:
			return newError("unknown operator: %s", operator)
		}
	case "++", "--":
		return newError("not implemented yet: %s%s", operator, right.Type())
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}
