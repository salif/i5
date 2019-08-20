package file

import (
	"fmt"
	"github.com/i5/i5/src/errors"
	"io/ioutil"
)

func Read(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		errors.NewFatalError(fmt.Sprintf(errors.READER_NOT_FOUND, fileName), 1)
	}
	return file
}
