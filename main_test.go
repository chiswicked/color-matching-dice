package main

import (
	"testing"
)

var printColorTests = []struct {
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

func TestPrintColor(t *testing.T) {
	for _, tc := range printColorTests {
		t.Run(tc.out, func(t *testing.T) {
			res := printColor(tc.in)
			if res != tc.out {
				t.Errorf("\ngot  %v\nwant %v", res, tc.out)
			}
		})
	}
}
