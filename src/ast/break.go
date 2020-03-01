// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

type Break struct {
	line  uint32
	token string
}

func (this Break) GetType() string {
	return BREAK
}

func (this Break) Debug() string {
	return this.token
}

func (this Break) GetLine() uint32 {
	return this.line
}

func (this Break) Init(line uint32, token string) Break {
	this.line = line
	this.token = token
	return this
}
