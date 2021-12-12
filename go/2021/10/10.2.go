package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var points = map[byte]uint64{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func score(line []byte) uint64 {
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
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}

	var ret uint64

	for i := len(stack) - 1; i >= 0; i-- {
		ret *= 5
		ret += points[stack[i]]
	}

	return ret
}

func main() {
	infile, err := os.Open("10.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)
	var acc []uint64

	for s.Scan() {
		s := score(s.Bytes())
		if s > 0 {
			acc = append(acc, s)
		}
	}

	sort.Slice(acc, func(i, j int) bool { return acc[i] < acc[j] })
	fmt.Println(acc)
	fmt.Println(acc[len(acc)/2])
}
