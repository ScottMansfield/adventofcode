package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filterChars = map[rune]struct{}{
	' ': {},
	',': {},
	'-': {},
	'>': {},
}

type point struct {
	x, y uint16
}

type line struct {
	start, end point
}

func readLines(s *bufio.Scanner) []line {
	var ret []line

	for s.Scan() {
		// Read each line stripping out all unwanted chars
		// parts will be start x, y end x, y
		parts := strings.FieldsFunc(s.Text(), func(r rune) bool {
			_, ok := filterChars[r]
			return ok
		})

		var nums [4]uint16

		for i := 0; i < 4; i++ {
			temp, err := strconv.ParseUint(parts[i], 10, 16)
			if err != nil {
				panic(err)
			}

			nums[i] = uint16(temp)
		}

		newline := line{
			start: point{
				x: nums[0],
				y: nums[1],
			},
			end: point{
				x: nums[2],
				y: nums[3],
			},
		}

		ret = append(ret, newline)
	}

	return ret
}

func maxes(lines []line) (uint16, uint16) {
	var mx, my uint16

	for _, l := range lines {
		if l.start.x > mx {
			mx = l.start.x
		}
		if l.start.y > my {
			my = l.start.y
		}
		if l.end.x > mx {
			mx = l.end.x
		}
		if l.end.y > my {
			my = l.end.y
		}
	}

	return mx, my
}

func drawLine(field [][]byte, l line) {
	if l.start.x == l.end.x {
		x := l.start.x
		start := l.start.y
		end := l.end.y

		if start > end {
			start, end = end, start
		}

		for y := start; y <= end; y++ {
			field[y][x]++
		}

		return

	}

	if l.start.y == l.end.y {
		// y is the same in both points
		y := l.start.y
		start := l.start.x
		end := l.end.x

		if start > end {
			start, end = end, start
		}

		for x := start; x <= end; x++ {
			field[y][x]++
		}

		return
	}

	// 45 degree diagonal only
	// reverse points if going upward
	if l.start.y > l.end.y {
		l.start, l.end = l.end, l.start
	}

	startx := l.start.x
	// endx := l.end.x

	starty := l.start.y
	endy := l.end.y

	dist := endy - starty
	// println(startx)
	// println(starty)
	// println(endy)
	// println(dist)

	// down right
	if l.start.x < l.end.x {
		for i := uint16(0); i <= dist; i++ {
			x := startx + i
			y := starty + i

			field[y][x]++
		}

		return
	}

	// down left
	// println("DOWN LEFT")
	for i := uint16(0); i <= dist; i++ {
		// println(i)
		x := startx - i
		y := starty + i

		// println(x)
		// println(y)

		field[y][x]++
	}
}

func countCrosses(field [][]byte) int {
	ret := 0

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] > 1 {
				ret++
			}
		}
	}

	return ret
}

func pprint(field [][]byte) {
	for _, temp := range field {
		fmt.Println(temp)
	}
}

func main() {
	infile, err := os.Open("05.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	lines := readLines(s)

	mx, my := maxes(lines)
	mx++
	my++

	// outer index is y (row), inner index is x (col)
	var field [][]byte

	for i := uint16(0); i < my; i++ {
		field = append(field, make([]byte, mx))
	}

	fmt.Println(mx, my)

	for _, l := range lines {
		fmt.Println(l)
		drawLine(field, l)
	}

	fmt.Println(countCrosses(field))
	// pprint(field)
}
