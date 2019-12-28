// SPDX-License-Identifier: GPL-3.0-or-later
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
