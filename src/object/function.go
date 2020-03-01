// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package object

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

type Function struct {
	Params []ast.Identifier
	Body   ast.Node
	Env    *Env
}

func (this Function) Type() string {
	return constants.TYPE_FUNCTION
}

func (this Function) StringValue() string {
	return fmt.Sprintf("[type: %v, params: %v]", this.Type(), len(this.Params))
}
