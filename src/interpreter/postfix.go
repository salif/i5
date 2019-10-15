// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
)

func evalPostfix(operator string, right object.Object, env *object.Env, line int) object.Object {
	switch operator {
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_POSTFIX, right.Type(), operator), Line: line}
	}
}
