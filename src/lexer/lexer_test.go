// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import (
	"testing"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

var (
	test_text1 string = `
	fn main() {
		$print(1)
	}
	`
	test_s1 int    = 15
	test1   []byte = []byte(test_text1)
)

func TestRun(t *testing.T) {
	var res1 types.TokenList = Run(test1)
	if res1.Size() != test_s1 {
		t.Errorf(constants.TEST_GOT_WANT, res1.Size(), test_s1)
	}
}
