// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalIf(node ast.If, env *object.Env) error {
	ev, err := Eval(node.GetCondition(), env)

	if err != nil {
		return err
	}

	if ev.Type() != constants.TYPE_BOOLEAN {
		return constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_IS_NOT_A_BOOL, ev.Type())}
	}

	if isTrue(ev) {
		_, err := Eval(node.GetConsequence(), env)
		if err != nil {
			return err
		}
	} else if node.HaveAlternative() {
		_, err := Eval(node.GetAlternative(), env)
		if err != nil {
			return err
		}
	}

	return nil
}
