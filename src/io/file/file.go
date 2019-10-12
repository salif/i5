// SPDX-License-Identifier: GPL-3.0-or-later
package file

import (
	"io/ioutil"
	"path/filepath"

	"github.com/i5/i5/src/constants"

	"github.com/i5/i5/src/io/console"
)

func GetFilesToRun(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		console.ThrowError(1, console.Format(constants.FILE_CANNOT_READ_DIR, dir))
	}
	filesToRun := []string{}
	for _, f := range files {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == constants.I5_FILE_EXT {
				filesToRun = append(filesToRun, f.Name())
			}
		}
	}
	return filesToRun
}
