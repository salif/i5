// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalIdentifier(node ast.Identifier, env *object.Env) (object.Object, error) {
	if val, ok := env.Get(node.GetValue()); ok {
		return val, nil
	} else {
		return nil, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_IDENT_NOT_FOUND, node.GetValue())}
	}
}
