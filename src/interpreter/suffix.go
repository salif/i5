// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalSuffix(operator string, right object.Object, env *object.Env) object.Object {
	switch operator {
	default:
		return newError(constants.IR_INVALID_SUFFIX, right.Type(), operator)
	}
}
