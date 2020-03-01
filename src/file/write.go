// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
