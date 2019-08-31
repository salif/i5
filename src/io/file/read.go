package file

import (
	"fmt"
	"io/ioutil"

	"github.com/i5/i5/src/errors"
)

func Read(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		errors.FatalError(fmt.Sprintf(errors.READER_NOT_FOUND, fileName), 1)
	}
	return file
}
