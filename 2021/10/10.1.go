package main

import (
	"bufio"
	"fmt"
	"os"
)

var points = map[byte]uint32{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func score(line []byte) uint32 {
	// store on a stack, push and pop, simple as that
	var stack []byte

	for _, c := range line {
		switch c {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case '<':
			stack = append(stack, '>')

		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			if stack[len(stack)-1] != c {
				return points[c]
			}
			stack = stack[:len(stack)-1]
		}
	}

	return 0
}

func main() {
	infile, err := os.Open("10.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)
	var acc uint32

	for s.Scan() {
		acc += score(s.Bytes())
	}

	fmt.Println(acc)
}
