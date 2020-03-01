// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import "strings"

type Prefix struct {
	line     uint32
	operator string
	right    Node
}

func (this Prefix) GetType() string {
	return PREFIX
}

func (this Prefix) Debug() string {
	var result strings.Builder
	result.WriteString("(")
	result.WriteString(this.operator)
	result.WriteString(this.right.Debug())
	result.WriteString(")")
	return result.String()
}

func (this Prefix) GetLine() uint32 {
	return this.line
}

func (this Prefix) GetOperator() string {
	return this.operator
}

func (this Prefix) GetRight() Node {
	return this.right
}

func (this Prefix) Init(line uint32, operator string) Prefix {
	this.line = line
	this.operator = operator
	return this
}

func (this *Prefix) SetRight(right Node) {
	this.right = right
}
