// Copyright © Steve McCoy.
// Licensed under the MIT License.

package main

import "testing"

type test struct {
	args   []string
	expect string
}

func TestPrefix(t *testing.T) {
	tests := []test{
		{[]string{"a", "b"}, ""},
		{[]string{"aa", "ab", "ac"}, "a"},
		{[]string{"Tommy [Disc 1]", "Tommy [Disc 2]"}, "Tommy"},
	}

	for _, test := range tests {
		p := prefix(test.args)
		if p != test.expect {
			t.Fatalf("Bad prefix %q for %v, expected %q",
				p, test.args, test.expect)
		}
	}
}
