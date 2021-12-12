package main

import (
	"bufio"
	"fmt"
	"os"
)

type field [10][10]uint8
type point struct{ x, y int }

func printField(f *field) {
	for _, row := range f {
		fmt.Println(row)
	}
	fmt.Println()
}

func flashInc(f *field, x, y int) int {
	// avoid out of bounds
	if x < 0 || y < 0 || x > 9 || y > 9 {
		return 0
	}

	// flash once only
	if f[y][x] < 10 {
		f[y][x]++
		return 1
	}

	return 0
}

func flashPropagate(f *field, x, y int) int {
	var ret int

	ret += flashInc(f, x-1, y-1)
	ret += flashInc(f, x, y-1)
	ret += flashInc(f, x+1, y-1)
	ret += flashInc(f, x+1, y)
	ret += flashInc(f, x+1, y+1)
	ret += flashInc(f, x, y+1)
	ret += flashInc(f, x-1, y+1)
	ret += flashInc(f, x-1, y)

	return ret
}

// runs a step and returns the number of flashes
func step(f *field) int {
	// step 1: increment
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			f[y][x]++
		}
	}

	// fmt.Println("AFTER INCREMENT")
	// printField(f)

	// step 2: propagate
	// just loop until no updates have occurred
	var flashed [10][10]bool
	for {
		var changes int

		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if f[y][x] == 10 && !flashed[y][x] {
					changes += flashPropagate(f, x, y)
					flashed[y][x] = true
				}
			}
		}

		if changes == 0 {
			break
		}
	}

	// fmt.Println("AFTER PROPAGATE")
	// printField(f)

	// step 2: count flashes and reset values
	var acc int

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if f[y][x] >= 10 {
				acc++
				f[y][x] = 0
			}
		}
	}

	return acc
}

func checkSimFlash(f *field) bool {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if f[y][x] > 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	infile, err := os.Open("11.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	var f field
	var y int

	for s.Scan() {
		for x, c := range s.Text() {
			f[y][x] = uint8(c) - '0'
		}

		y++
	}

	printField(&f)

	var i int

	for {
		i++
		step(&f)
		//printField(&f)
		if checkSimFlash(&f) {
			fmt.Println(i)
			break
		}
	}
}
