// SPDX-License-Identifier: GPL-3.0-or-later
package file

import (
	"io/ioutil"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
)

// Read file or throw error
func Read(name string) []byte {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		console.ThrowError(1, constants.FILE_CANNOT_READ, name)
	}
	return content
}
