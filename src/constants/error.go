// SPDX-License-Identifier: GPL-3.0-or-later
package constants

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
