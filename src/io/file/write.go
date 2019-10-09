// SPDX-License-Identifier: GPL-3.0-or-later
package file

import (
	"io/ioutil"
	"os"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
)

// Write file and return result or error
func Write(name string, content string, perm os.FileMode) (result string) {
	err := ioutil.WriteFile(name, []byte(content), perm)
	if err != nil {
		return console.Format(constants.FILE_CANNOT_WRITE, name)
	} else {
		return
	}

}
