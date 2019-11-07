// SPDX-License-Identifier: GPL-3.0-or-later
package file

import (
	"io/ioutil"
	"os"
)

// Write file and return result or error
func Write(name string, content string, perm os.FileMode) error {
	err := ioutil.WriteFile(name, []byte(content), perm)
	if err != nil {
		return err
	} else {
		return nil
	}
}
