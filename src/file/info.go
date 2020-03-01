// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
