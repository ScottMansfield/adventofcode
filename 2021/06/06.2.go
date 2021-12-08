package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func simulate(fish []uint64) {
	temp6, temp8 := fish[0], fish[0]

	fish[0] = fish[1]
	fish[1] = fish[2]
	fish[2] = fish[3]
	fish[3] = fish[4]
	fish[4] = fish[5]
	fish[5] = fish[6]
	fish[6] = fish[7] + temp6
	fish[7] = fish[8]
	fish[8] = temp8
}

func main() {
	infile, err := os.Open("06.input")
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadAll(infile)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), ",")

	fish := make([]uint64, 9)

	for _, p := range parts {
		temp, err := strconv.ParseUint(p, 10, 8)
		if err != nil {
			panic(err)
		}
		fish[temp]++
	}

	fmt.Println(fish)

	for i := 0; i < 256; i++ {
		simulate(fish)
		fmt.Println(fish)
	}

	var acc uint64
	for _, f := range fish {
		acc += f
	}

	fmt.Println(acc)
}
