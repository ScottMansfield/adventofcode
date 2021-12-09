package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func localMinimum(field [][]byte, x, y int) bool {
	maxx := len(field[0]) - 1
	maxy := len(field) - 1

	val := field[y][x]

	// Corners
	// Top left
	if x == 0 && y == 0 {
		return val < field[0][1] &&
			val < field[1][0]
	}

	// Bottom left
	if x == 0 && y == maxy {
		return val < field[maxy][1] &&
			val < field[maxy-1][0]
	}

	// Top right
	if x == maxx && y == 0 {
		return val < field[0][maxx-1] &&
			val < field[1][maxx]
	}

	// Bottom right
	if x == maxx && y == maxy {
		return val < field[maxy][maxx-1] &&
			val < field[maxy-1][maxx]
	}

	// Edges
	// Left
	if x == 0 {
		return val < field[y][x+1] &&
			val < field[y-1][x] &&
			val < field[y+1][x]
	}

	// Top
	if y == 0 {
		return val < field[y+1][x] &&
			val < field[y][x-1] &&
			val < field[y][x+1]
	}

	// Right
	if x == maxx {
		return val < field[y][x-1] &&
			val < field[y-1][x] &&
			val < field[y+1][x]
	}

	// Bottom
	if y == maxy {
		return val < field[y-1][x] &&
			val < field[y][x-1] &&
			val < field[y][x+1]
	}

	return val < field[y][x-1] &&
		val < field[y][x+1] &&
		val < field[y-1][x] &&
		val < field[y+1][x]
}

func basinSize(field [][]byte, visited [][]bool, x, y int) int {
	if field[y][x] == 9 || visited[y][x] {
		return 0
	}

	visited[y][x] = true

	ret := 1

	// Up
	if y > 0 {
		ret += basinSize(field, visited, x, y-1)
	}
	// Down
	if y < len(field)-1 {
		ret += basinSize(field, visited, x, y+1)
	}
	// Left
	if x > 0 {
		ret += basinSize(field, visited, x-1, y)
	}
	// Right
	if x < len(field[0])-1 {
		ret += basinSize(field, visited, x+1, y)
	}

	return ret
}

func main() {
	infile, err := os.Open("09.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)
	var field [][]byte

	for s.Scan() {
		var line []byte
		for _, c := range s.Text() {
			line = append(line, byte(c)-byte('0'))
		}
		field = append(field, line)
	}

	//fmt.Println(field)

	// an array to avoid infinite recursion
	visited := make([][]bool, len(field))
	for i := range visited {
		visited[i] = make([]bool, len(field[0]))
	}

	var basinSizes []int

	for y, row := range field {
		for x := range row {
			if localMinimum(field, x, y) {
				//fmt.Println(x, y)
				bs := basinSize(field, visited, x, y)
				//fmt.Println(bs)
				basinSizes = append(basinSizes, bs)
			}
		}
	}

	sort.Ints(basinSizes)

	fmt.Println(basinSizes[len(basinSizes)-3:])

	acc := 1
	acc *= basinSizes[len(basinSizes)-1]
	acc *= basinSizes[len(basinSizes)-2]
	acc *= basinSizes[len(basinSizes)-3]

	fmt.Println(acc)
}
