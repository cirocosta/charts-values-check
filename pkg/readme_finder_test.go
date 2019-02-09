package pkg_test

import (
	"testing"

	. "github.com/cirocosta/charts-values-check/pkg"
)

func TestTablesFinder(t *testing.T) {
	var testCases = []struct {
		desc        string
		input       string
		expected    []string
		shouldError bool
	}{
		{
			desc:        "empty",
			input:       ``,
			shouldError: true,
		},
		{
			desc:     "table without rows",
			expected: []string{},
			input: `
|a|b|
|---|----|`,
		},
		{
			desc:        "table without code",
			shouldError: true,
			input: `
|a|b|
|---|----|
|this.that|else|`,
		},
		{
			desc:     "table with code",
			expected: []string{"this.that"},
			input: `
|a|b|
|---|----|
|` + "`this.that`" + `| else |`,
		},
	}

	var (
		actual []string
		err    error
	)
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual, err = (&ReadmeFinder{}).Find([]byte(tc.input))
			if tc.shouldError {
				if err == nil {
					t.Errorf("should've errored")
				}

				return
			}

			if err != nil {
				t.Errorf("should not have errored")
			}

			if !SlicesEqual(actual, tc.expected) {
				t.Errorf("%+v != %v\n", actual, tc.expected)
			}
		})
	}
}
