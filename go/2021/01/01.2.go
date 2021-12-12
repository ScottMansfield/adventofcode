package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var buf [3]uint64

func next(num uint64) uint64 {
	buf[0] = buf[1]
	buf[1] = buf[2]
	buf[2] = num

	return buf[0] + buf[1] + buf[2]
}

func main() {
	infile, err := os.Open("01.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	var prev uint64
	var seen uint64
	var count uint64

	for s.Scan() {
		num, err := strconv.ParseUint(s.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		num = next(num)

		seen++
		if seen <= 3 {
			continue
		}

		if prev < num {
			count++
		}

		prev = num
	}

	fmt.Println(count)
}
