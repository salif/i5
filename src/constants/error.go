// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package constants

import "fmt"

type Error struct {
	Line    uint32
	Message string
	Type    byte
	Value   interface{}
}

func (this Error) Error() string {
	return this.Message
}

const (
	ERROR_FATAL  byte = 0
	ERROR_RETURN byte = 1
	ERROR_BREAK  byte = 2
)

type SyntaxError struct {
	Message string
	In      string
}

func (this SyntaxError) Error() string {
	return fmt.Sprintf(SYNTAX_ERROR, this.Message, this.In)
}
