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

var positionTests = []struct {
	in  die
	out int
}{
	{
		in:  die{red, green, blue, yellow, yellow, blue},
		out: int(red + green + blue + yellow + yellow + blue),
	},
	{
		in:  die{red, red, blue, yellow, yellow, red},
		out: int(red + red + blue + yellow + yellow + red),
	},
	{
		in:  die{red, blue, blue, yellow, green, blue},
		out: int(red + blue + blue + yellow + green + blue),
	},
}

func TestPosition(t *testing.T) {
	for _, tc := range positionTests {
		t.Run("Rotate die 6x", func(t *testing.T) {
			res := 0
			for i := 0; i < 6; i++ {
				tc.in.position(i)
				res += int(tc.in[0])
			}
			if res != tc.out {
				t.Errorf("\ngot  %v\nwant %v", res, tc.out)
			}
		})
	}
}

var allSidesUniqueTests = []struct {
	in  dice
	out bool
}{
	{
		in: dice{
			{green, green, green, green, green, green},
			{blue, blue, blue, blue, blue},
			{red, red, red, red, red, red},
			{yellow, yellow, yellow, yellow, yellow, yellow},
		},
		out: true,
	},
	{
		in: dice{
			{green, yellow, red, blue, green, green},
			{blue, green, yellow, red, blue},
			{red, blue, green, yellow, red, red},
			{yellow, red, blue, green, yellow, yellow},
		},
		out: true,
	},
	{
		in: dice{
			{green, green, yellow, green, green, green},
			{blue, blue, blue, blue, blue},
			{red, red, red, red, red, red},
			{yellow, yellow, yellow, yellow, yellow, yellow},
		},
		out: false,
	},
	{
		in: dice{
			{green, yellow, red, blue, green, green},
			{blue, blue, yellow, red, blue},
			{red, blue, green, yellow, red, red},
			{yellow, red, blue, green, yellow, yellow},
		},
		out: false,
	},
}

func TestAllSidesUnique(t *testing.T) {
	for _, tc := range allSidesUniqueTests {
		t.Run("All sides unique", func(t *testing.T) {
			res := tc.in.allSidesUnique()
			if res != tc.out {
				t.Errorf("\ngot  %v\nwant %v", res, tc.out)
			}
		})
	}
}
