// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

import (
	"testing"

	"github.com/i5/i5/src/constants"
)

type mtest struct {
	T string
	W int
}

var tests []mtest = []mtest{
	mtest{
		T: "", W: 1,
	},
	mtest{
		T: `
		`, W: 2,
	},
	mtest{
		T: `
		main = () => {
			$print(1)
		}
		`, W: 16,
	},
	mtest{
		T: `
		a = (() => {})()
		`, W: 14,
	},
}

func TestRun(t *testing.T) {
	for _, tt := range tests {
		result, err := Run("test code", []byte(tt.T))
		if err != nil {
			t.Errorf("error: %v\n", err.Error())
		}
		if result.Size() != tt.W {
			t.Errorf(constants.TEST_GOT_WANT, result.Size(), tt.W)
		}
	}
}
