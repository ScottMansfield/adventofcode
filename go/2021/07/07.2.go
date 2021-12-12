package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func max(crabs []uint16) uint16 {
	var ret uint16

	for _, c := range crabs {
		if c > ret {
			ret = c
		}
	}

	return ret
}

func distCost(dist uint32) uint32 {
	// closed form of sum of first n digits is n(n+1)/2
	return (dist * (dist + 1)) / 2
}

func cost(crabs []uint16, pos uint16) uint32 {
	var ret uint32

	for _, c := range crabs {
		if c > pos {
			ret += distCost(uint32(c - pos))
		} else {
			ret += distCost(uint32(pos - c))
		}
	}

	return ret
}

func main() {
	infile, err := os.Open("07.input")
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadAll(infile)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), ",")

	crabs := make([]uint16, len(parts))

	for i, p := range parts {
		temp, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			panic(err)
		}
		crabs[i] = uint16(temp)
	}

	fmt.Println(crabs)

	maxCrab := max(crabs)

	fmt.Println(maxCrab)

	minCost := uint32(math.MaxUint32)

	for pos := uint16(0); pos <= maxCrab; pos++ {
		curCost := cost(crabs, pos)

		if curCost < minCost {
			minCost = curCost
		}
	}

	fmt.Println(minCost)
}
