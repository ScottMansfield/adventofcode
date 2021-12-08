package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func countOnes(vals []string, digit int) int {
	var ret int

	for _, val := range vals {
		if val[digit] == '1' {
			ret++
		}
	}

	return ret
}

func keepVal(val string, digit int, ones int, total int, keepMajority bool) bool {
	// Shortcut if all values the same
	if ones == total || ones == 0 {
		return true
	}

	// Tie breaking with even numbers
	if (total%2 == 0) && ones == total/2 {
		if keepMajority && val[digit] == '1' {
			return true
		}
		if !keepMajority && val[digit] == '0' {
			return true
		}
	}

	onesMajority := ones >= int(math.Ceil(float64(total)/2))

	if val[digit] == '1' {
		if onesMajority && keepMajority {
			return true
		}
		if !onesMajority && !keepMajority {
			return true
		}
	}

	if val[digit] == '0' {
		if !onesMajority && keepMajority {
			return true
		}
		if onesMajority && !keepMajority {
			return true
		}
	}

	return false
}

func tournamentRound(vals []string, digit int, keepMajority bool) []string {
	var ret []string

	ones := countOnes(vals, digit)

	for _, val := range vals {
		if keepVal(val, digit, ones, len(vals), keepMajority) {
			ret = append(ret, val)
		}
	}

	return ret
}

func main() {
	infile, err := os.Open("03.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	var lines []string

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	digit := 0
	oxylines := tournamentRound(lines, 0, true)
	for len(oxylines) > 1 {
		println(len(oxylines))
		digit++
		oxylines = tournamentRound(oxylines, digit, true)
	}

	fmt.Println(oxylines)

	digit = 0
	co2lines := tournamentRound(lines, 0, false)
	for len(co2lines) > 1 {
		println(len(co2lines))
		digit++
		co2lines = tournamentRound(co2lines, digit, false)
	}

	fmt.Println(co2lines)

	oxy, err := strconv.ParseUint(oxylines[0], 2, 64)
	if err != nil {
		panic(err)
	}

	co2, err := strconv.ParseUint(co2lines[0], 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(oxy, co2, oxy*co2)
}
