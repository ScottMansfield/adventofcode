package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type board [5][5]byte

const sentinel byte = 255

func readMoves(s *bufio.Scanner) []byte {
	// Read in the list of called numbers
	s.Scan()
	moveStrings := strings.Split(s.Text(), ",")
	ret := make([]byte, len(moveStrings))

	for i, ns := range moveStrings {
		temp, err := strconv.ParseUint(ns, 10, 8)
		if err != nil {
			panic(err)
		}

		ret[i] = byte(temp)
	}

	return ret
}

func readBoards(s *bufio.Scanner) []board {
	var ret []board

	// Read in all the boards
	// one blank line then 5 lines of 5 numbers
	for s.Scan() {
		var b [5][5]byte
		for i := 0; i < 5; i++ {
			s.Scan()
			parts := strings.Fields(s.Text())
			for j, ls := range parts {
				temp, err := strconv.ParseUint(ls, 10, 8)
				if err != nil {
					panic(err)
				}

				b[i][j] = byte(temp)
			}
		}

		ret = append(ret, b)
	}

	return ret
}

func applyMove(b []board, idx int, move byte) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[idx][i][j] == move {
				b[idx][i][j] = sentinel
				return
			}
		}
	}
}

func won(b []board, idx int) bool {
	var colLose, rowLose [5]bool

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[idx][i][j] != sentinel {
				rowLose[j] = true
				colLose[i] = true
			}
		}
	}

	for i := 0; i < 5; i++ {
		if !colLose[i] || !rowLose[i] {
			return true
		}
	}

	return false
}

func score(b []board, idx int, move byte) int {
	var acc int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[idx][i][j] != sentinel {
				acc += int(b[idx][i][j])
			}
		}
	}

	return acc * int(move)
}

func main() {
	infile, err := os.Open("04.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	moves := readMoves(s)
	fmt.Println(moves)

	boards := readBoards(s)
	fmt.Println(len(boards))

	removed := make(map[int]struct{})

	for _, m := range moves {
		fmt.Println(m)

		for i := range boards {
			// Mask off boards that have won
			if _, ok := removed[i]; ok {
				continue
			}

			applyMove(boards, i, m)
			if won(boards, i) {
				if len(removed) == len(boards)-1 {
					fmt.Println(score(boards, i, m))
					return
				}

				removed[i] = struct{}{}
			}
		}
	}
}
