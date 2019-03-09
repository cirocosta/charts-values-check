package pkg_test

import (
	"testing"

	. "github.com/cirocosta/charts-values-check/pkg"
)

func TestValuesFinder(t *testing.T) {
	var testCases = []struct {
		desc        string
		input       string
		expected    []string
		shouldError bool
	}{
		{
			desc:        "non-yaml",
			input:       `;this-is-wrong`,
			shouldError: true,
		},
		{
			desc:     "no values",
			input:    `---`,
			expected: []string{},
		},
		{
			desc:     "single with string",
			input:    `foo: bar`,
			expected: []string{"foo"},
		},
		{
			desc:     "single with map",
			input:    `foo: {baz: caz}`,
			expected: []string{"foo.baz"},
		},
		{
			desc:     "single with nil",
			input:    `foo: null`,
			expected: []string{"foo"},
		},
		{
			desc:     "single empty map",
			input:    `foo: {}`,
			expected: []string{"foo"},
		},
		{
			desc:     "single empty array",
			input:    `foo: []`,
			expected: []string{"foo"},
		},
	}

	var (
		actual []string
		err    error
	)
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual, err = (&ValuesFinder{}).Find([]byte(tc.input))
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
