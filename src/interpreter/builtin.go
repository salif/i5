// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/builtins"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalBuiltin(node ast.Builtin, env *object.Env) (object.Object, error) {
	ev, ok := builtins.Get(node.GetValue(), env)
	if ok {
		return ev, nil
	} else {
		return nil, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_BUILTIN_NOT_FOUND, node.GetValue())}
	}
}
