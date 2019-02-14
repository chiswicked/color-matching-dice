package main

import (
	"fmt"
	"strings"
	"time"
)

type die [6]color // 0: top, 1: left, 2: front, 3: right, 4: rear, 5: bottom
type dice [4]die

type color int

// 4 colors add up to 15
const (
	red color = 1 << iota
	green
	blue
	yellow
)

var perm int
var sol int

func main() {
	started := time.Now()
	// Base set from which copies are created
	var set dice
	set[0] = die{green, yellow, blue, yellow, red, yellow}
	set[1] = die{green, blue, blue, yellow, red, green}
	set[2] = die{red, yellow, yellow, blue, red, green}
	set[3] = die{yellow, blue, green, red, red, green}

	perm = 0 // Number of permutations checked
	sol = 0  // Number of solutions found

	for d0 := 0; d0 < 4; d0++ { // Die in position 1
		for d1 := 0; d1 < 4; d1++ { // Die in position 2
			if d1 == d0 { // Can't repeat the same die
				continue
			}
			for d2 := 0; d2 < 4; d2++ { // Die in position 3
				if d2 == d1 || d2 == d0 { // Can't repeat the same die
					continue
				}
				for d3 := 0; d3 < 4; d3++ { // Die in position 4
					if d3 == d0 || d3 == d1 || d3 == d2 { // Can't repeat the same die
						continue
					}
					rowPerm := dice{set[d0].copy(), set[d1].copy(), set[d2].copy(), set[d3].copy()}
					rotPerms(rowPerm)
				}
			}
		}
	}
	fmt.Printf("%v permutations checked, %v solutions found. Processing took %v\n",
		perm,
		sol,
		time.Since(started))
}

// Takes an ordered row of dice and iterates through all dice rotation permutations.
// If a match is found, prints out the dice positions.
func rotPerms(row dice) {
	for pos0 := 0; pos0 < 6; pos0++ { // Die 1 - What's on top
		for r0 := 0; r0 < 4; r0++ { // Die 1 - Rotate Y
			for pos1 := 0; pos1 < 6; pos1++ { // Die 2 - What's on top
				for r1 := 0; r1 < 4; r1++ { // Die 2 - Rotate Y
					for pos2 := 0; pos2 < 6; pos2++ { // Die 3 - What's on top
						for r2 := 0; r2 < 4; r2++ { // Die 3 - Rotate Y
							for pos3 := 0; pos3 < 6; pos3++ { // Die 4 - What's on top
								for r3 := 0; r3 < 4; r3++ { // Die 4 - Rotate Y
									perm++
									if row.allSidesUnique() {
										sol++
										row.string(sol)
									}
									row[3].rotateY()
								}
								row[3].position(pos3)
							}
							row[2].rotateY()
						}
						row[2].position(pos2)
					}
					row[1].rotateY()
				}
				row[1].position(pos1)
			}
			row[0].rotateY()
		}
		row[0].position(pos0)
	}
}

// Using different copies for concurrent processing
func (d *die) copy() die {
	var res [6]color
	copy(res[:], d[:])
	return res
}

// 6 alternating X and Z rotations iterate through all sides
func (d *die) position(i int) {
	if i%2 == 0 {
		d.rotateX()
		return
	}
	d.rotateZ()
}

func (d *die) rotateX() {
	d[0], d[2], d[5], d[4] = d[2], d[5], d[4], d[0]
}

func (d *die) rotateY() {
	d[1], d[2], d[3], d[4] = d[4], d[1], d[2], d[3]
}

func (d *die) rotateZ() {
	d[0], d[1], d[5], d[3] = d[3], d[0], d[1], d[5]
}

// Returns whether the dice arrangement gives a match
// (no color repeat itslef on either side)
func (d *dice) allSidesUnique() bool {
	for side := 1; side < 5; side++ {
		if d[0][side]+d[1][side]+d[2][side]+d[3][side] != 15 {
			return false
		}
	}
	return true
}

// Gives the string representation of a color
func fmtColor(i color) string {
	switch i {
	case 1:
		return "Red"
	case 2:
		return "Green"
	case 4:
		return "Blue"
	case 8:
		return "Yellow"
	default:
		panic("?!")
	}
}

// Prints out the dice positions to Stdout
func (d *dice) string(sol int) {
	fmt.Printf("Solution #%v\n", sol)
	fmt.Println(strings.Repeat("-", 78))
	fmt.Printf("|%-10v|%-10v|%-10v|%-10v|%-10v|%-10v|%-10v|\n", "Die", "Top", "Left", "Front", "Right", "Rear", "Bottom")
	fmt.Println(strings.Repeat("-", 78))
	for num, die := range d {
		fmt.Printf("|%-10v|%-10v|%-10v|%-10v|%-10v|%-10v|%-10v|\n",
			fmt.Sprintf("#%v", num+1),
			fmtColor(die[0]),
			fmtColor(die[1]),
			fmtColor(die[2]),
			fmtColor(die[3]),
			fmtColor(die[4]),
			fmtColor(die[5]))
	}
	fmt.Println(strings.Repeat("-", 78) + "\n")
}
