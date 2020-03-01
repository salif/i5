// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

type String struct {
	line  uint32
	value string
}

func (this String) GetType() string {
	return STRING
}

func (this String) Debug() string {
	return "\"" + this.value + "\""
}

func (this String) GetLine() uint32 {
	return this.line
}

func (this String) GetValue() string {
	return this.value
}

func (this String) Init(line uint32, value string) String {
	this.line = line
	this.value = value
	return this
}
