// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
