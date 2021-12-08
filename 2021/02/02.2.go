package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	infile, err := os.Open("02.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	var dist, depth, aim uint64

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")

		num, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "forward":
			dist += num
			depth += aim * num

		case "down":
			aim += num

		case "up":
			aim -= num

		default:
			panic(parts[0])
		}
	}

	fmt.Println(dist, depth, aim, dist*depth)
}
