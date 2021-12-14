package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	e1, e2 byte
}

type sequence map[pair]uint64

type ruleSet map[pair]byte

func buildSequence(init string) sequence {
	ret := make(sequence)

	for i := 0; i < len(init)-1; i++ {
		key := pair{
			init[i],
			init[i+1],
		}

		if n, ok := ret[key]; ok {
			ret[key] = n + 1
		} else {
			ret[key] = 1
		}
	}

	return ret
}

func readRules(s *bufio.Scanner) ruleSet {
	ret := make(ruleSet)

	// example BC -> B
	for s.Scan() {
		parts := strings.FieldsFunc(s.Text(), func(r rune) bool {
			return r == ' ' || r == '-' || r == '>'
		})

		// parts = ["BC", "B"]
		ret[pair{parts[0][0], parts[0][1]}] = parts[1][0]
	}

	return ret
}

func insertOrAdd(seq sequence, key pair, count uint64) {
	if n, ok := seq[key]; ok {
		seq[key] = n + count
	} else {
		seq[key] = count
	}
}

func applyRules(seq sequence, rules ruleSet) sequence {
	ret := make(sequence)

	for keyPair, count := range seq {
		if insert, ok := rules[keyPair]; ok {
			key1 := pair{
				keyPair.e1,
				insert,
			}
			key2 := pair{
				insert,
				keyPair.e2,
			}

			insertOrAdd(ret, key1, count)
			insertOrAdd(ret, key2, count)
		}
	}

	return ret
}

func insertOrAddCounts(counts map[byte]uint64, key byte, count uint64) {
	if n, ok := counts[key]; ok {
		counts[key] = n + count
	} else {
		counts[key] = count
	}
}

func countChars(seq sequence) map[byte]uint64 {
	ret := make(map[byte]uint64)

	for keyPair, count := range seq {
		insertOrAddCounts(ret, keyPair.e1, count)
		insertOrAddCounts(ret, keyPair.e2, count)
	}

	return ret
}

func main() {
	infile, err := os.Open("14.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)
	s.Scan()

	init := s.Text()
	startChar := init[0]
	endChar := init[len(init)-1]

	seq := buildSequence(init)

	// skip blank line
	s.Scan()

	rules := readRules(s)

	for i := 0; i < 40; i++ {
		fmt.Println(i)
		seq = applyRules(seq, rules)
		//fmt.Println(seq)
	}

	counts := countChars(seq)

	for char, count := range counts {
		counts[char] = count / 2
	}

	counts[startChar] = counts[startChar] + 1
	counts[endChar] = counts[endChar] + 1

	for char, count := range counts {
		fmt.Printf("%c: %d\n", char, count)
	}
}
