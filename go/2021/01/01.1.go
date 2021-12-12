package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	infile, err := os.Open("01.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	var prev uint64
	var count uint32

	for s.Scan() {
		num, err := strconv.ParseUint(s.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		if prev < num {
			count++
		}

		prev = num
	}

	fmt.Println(count)
}
