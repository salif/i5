// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalReturn(node ast.Return, env *object.Env) error {
	ev, err := Eval(node.GetBody(), env)
	if err != nil {
		return err
	}
	return constants.Error{Line: node.GetLine(), Type: constants.ERROR_RETURN, Message: "", Value: ev}
}
