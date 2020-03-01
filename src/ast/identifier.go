// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

type Identifier struct {
	line  uint32
	token string
}

func (this Identifier) GetType() string {
	return IDENTIFIER
}

func (this Identifier) Debug() string {
	return this.token
}

func (this Identifier) GetLine() uint32 {
	return this.line
}

func (this Identifier) GetValue() string {
	return this.token
}

func (this Identifier) Init(line uint32, token string) Identifier {
	this.line = line
	this.token = token
	return this
}
