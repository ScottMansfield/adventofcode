package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	element byte
	next    *node
}

type ruleKey struct {
	e1, e2 byte
}

type ruleSet map[ruleKey]byte

func buildList(init string) *node {
	var ret *node

	// run backwards through string creating linked list
	for i := len(init) - 1; i >= 0; i-- {
		ret = &node{
			element: init[i],
			next:    ret,
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
		ret[ruleKey{parts[0][0], parts[0][1]}] = parts[1][0]
	}

	return ret
}

func applyRules(list *node, r ruleSet) {
	prev := list
	list = list.next

	for list != nil {
		key := ruleKey{
			prev.element,
			list.element,
		}

		if insert, ok := r[key]; ok {
			newNode := &node{
				element: insert,
				next:    list,
			}

			prev.next = newNode
		}

		prev = list
		list = list.next
	}
}

func printList(list *node) {
	for list != nil {
		fmt.Printf("%c", list.element)
		list = list.next
	}
	fmt.Println()
}

func countChars(list *node) map[byte]uint32 {
	ret := make(map[byte]uint32)

	for list != nil {
		if n, ok := ret[list.element]; ok {
			ret[list.element] = n + 1
		} else {
			ret[list.element] = 1
		}

		list = list.next
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

	list := buildList(s.Text())

	// skip blank line
	s.Scan()

	rules := readRules(s)

	printList(list)

	for i := 0; i < 10; i++ {
		applyRules(list, rules)

		if i < 4 {
			printList(list)
		}
	}

	counts := countChars(list)

	fmt.Println(counts)
}
