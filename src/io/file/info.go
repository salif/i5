// SPDX-License-Identifier: GPL-3.0-or-later
package file

import "os"

// return:
// 1 if file/dir not exists
// 2 if file/dir is dir
// 3 if file/dir is file
// else return 0
func Info(name string) int {
	f, err := os.Stat(name)
	if os.IsNotExist(err) {
		return 1
	}
	switch mode := f.Mode(); {
	case mode.IsDir():
		return 2
	case mode.IsRegular():
		return 3
	default:
		return 0
	}
}
