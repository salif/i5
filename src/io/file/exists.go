// SPDX-License-Identifier: GPL-3.0-or-later
package file

import "os"

// return true if file/dir exists
func Exists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
