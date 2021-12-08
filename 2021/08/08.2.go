package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strings"
)

func toByte(s string) uint8 {
	var ret uint8

	for _, c := range s {
		switch c {
		case 'a':
			ret |= 0x01
		case 'b':
			ret |= 0x02
		case 'c':
			ret |= 0x04
		case 'd':
			ret |= 0x08
		case 'e':
			ret |= 0x10
		case 'f':
			ret |= 0x20
		case 'g':
			ret |= 0x40
		}
	}

	return ret
}

// 7 - 1     = top
// 8 - 0     = middle
// 9 - 7 - 4 = bottom
// 4 - 3     = topleft
// 8 - 6     = topright
// 8 - 9     = botleft
// 3 - 2     = botright

func buildDecoder(decode []string) [10]uint8 {
	// Putting them in order means the array has:
	// 1:     2 segments
	// 7:     3 segments
	// 4:     4 segments
	// 2,3,5: 5 segments
	// 0,6,9: 6 segments
	// 8:     7 segments
	sort.Slice(decode, func(i, j int) bool {
		return len(decode[i]) < len(decode[j])
	})

	// Convert each string to a byte to avoid sorting characters
	// and make math easier
	decodeBytes := make([]uint8, len(decode))

	for i, ds := range decode {
		decodeBytes[i] = toByte(ds)
	}

	var ret [10]uint8

	// Fill in known values
	ret[1] = decodeBytes[0]
	ret[4] = decodeBytes[2]
	ret[7] = decodeBytes[1]
	ret[8] = decodeBytes[9]

	// 8 - 7 - 4 gives bottom left and bottom (2 bits)
	// 2 + bottom left and bottom = 2
	// 3 and 5 & botleftbot = bottom
	// 1 & 3 == 1
	botleftbot := ret[8] & ^ret[7] & ^ret[4]

	// Array showing found items in the decodeBytes array
	// the odd man out is 5
	found := make([]bool, 6)

	// loop over 2,3,5 area
	var bottom uint8
	for i := 3; i <= 5; i++ {
		if decodeBytes[i]&botleftbot == botleftbot {
			ret[2] = decodeBytes[i]
			found[i] = true
		} else {
			bottom = decodeBytes[i] & botleftbot
		}

		if decodeBytes[i]&ret[1] == ret[1] {
			ret[3] = decodeBytes[i]
			found[i] = true
		}
	}

	// pick out the odd man out (5)
	for i := 3; i <= 5; i++ {
		if !found[i] {
			ret[5] = decodeBytes[i]
		}
	}

	botleft := botleftbot & ^bottom

	// 6 & 1 will give only one bit set, which is the bottom right
	// 9 & botleft = 0
	// loop over 0,6,9 area

	// Array showing found items in the decodeBytes array
	// the odd man out is 5
	found = make([]bool, 9)

	// loop over 6 segment digit area
	for i := 6; i <= 8; i++ {
		temp := decodeBytes[i] & ret[1]
		if bits.OnesCount8(temp) == 1 {
			ret[6] = decodeBytes[i]
			found[i] = true
		}

		temp = decodeBytes[i] & botleft
		if temp == 0 {
			ret[9] = decodeBytes[i]
			found[i] = true
		}
	}

	// pick out the odd man out (0)
	for i := 6; i <= 8; i++ {
		if !found[i] {
			ret[0] = decodeBytes[i]
		}
	}

	return ret
}

func decode(nums []string, decoder [10]uint8) uint32 {
	var acc uint32

outer:
	for _, num := range nums {
		lookup := toByte(num)

		for i := 0; i < 10; i++ {
			if lookup == decoder[i] {
				acc *= 10
				acc += uint32(i)
				continue outer
			}
		}

		panic(lookup)
	}

	return acc
}

func main() {
	infile, err := os.Open("08.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)
	var acc uint32

	for s.Scan() {
		if s.Text() == "" {
			break
		}
		parts := strings.Split(string(s.Text()), "|")
		key := strings.Fields(parts[0])
		outnums := strings.Fields(parts[1])

		decoder := buildDecoder(key)

		fmt.Println(decoder)

		acc += decode(outnums, decoder)
		println(acc)
	}

	fmt.Println(acc)
}
