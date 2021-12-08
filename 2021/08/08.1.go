package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	infile, err := os.Open("08.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)
	var acc int

	for s.Scan() {
		parts := strings.Split(string(s.Text()), "|")
		parts = strings.Fields(parts[1])

		for _, p := range parts {
			switch len(p) {
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 7:
				acc++
			}
		}
	}

	fmt.Println(acc)
}
