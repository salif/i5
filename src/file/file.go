// SPDX-License-Identifier: GPL-3.0-or-later
package file

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/i5/i5/src/constants"
)

func GetFilesToRun(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	filesToRun := []string{}
	for _, f := range files {
		if !f.IsDir() {
			if filepath.Ext(f.Name()) == constants.I5_FILE_EXT {
				if err != nil {
					return nil, err
				}
				filesToRun = append(filesToRun, fmt.Sprintf("%v/%v", dir, f.Name()))
			}
		}
	}
	return filesToRun, nil
}
