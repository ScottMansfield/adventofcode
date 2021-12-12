package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	infile, err := os.Open("03.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	// 12 bits in each number
	var ones = make([]int, 12)
	var total int

	for s.Scan() {
		total++

		for i, c := range s.Text() {
			if c == '1' {
				ones[i]++
			}
		}
	}

	var gamma int

	for _, count := range ones {
		gamma <<= 1

		if count > total/2 {
			gamma |= 1
		}
	}

	epsilon := (^gamma) & 0xFFF

	fmt.Println(total)
	fmt.Println(ones)
	fmt.Println(gamma, epsilon, gamma*epsilon)
}
