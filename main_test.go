package main

import (
	"testing"
)

var fmtColorTests = []struct {
	in  color
	out string
}{
	{
		in:  red,
		out: "Red",
	}, {
		in:  green,
		out: "Green",
	}, {
		in:  blue,
		out: "Blue",
	}, {
		in:  yellow,
		out: "Yellow",
	},
}

func TestFmtColor(t *testing.T) {
	for _, tc := range fmtColorTests {
		t.Run(tc.out, func(t *testing.T) {
			res := fmtColor(tc.in)
			if res != tc.out {
				t.Errorf("\ngot  %v\nwant %v", res, tc.out)
			}
		})
	}
}
