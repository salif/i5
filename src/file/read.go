// SPDX-License-Identifier: GPL-3.0-or-later
package file

import (
	"io/ioutil"
)

// Read file or throw error
func Read(name string) ([]byte, error) {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return content, nil
}
