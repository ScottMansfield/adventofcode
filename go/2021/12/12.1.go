package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stringSet map[string]struct{}

func (s stringSet) copyWith(name string) stringSet {
	ret := make(stringSet)
	for k := range s {
		ret[k] = struct{}{}
	}
	ret[name] = struct{}{}
	return ret
}

type graph map[string]stringSet

func ensureNode(g graph, name string) {
	if _, ok := g[name]; !ok {
		g[name] = make(stringSet)
	}
}

// Assumes nodes exist
func connect(g graph, node1, node2 string) {
	ensureNode(g, node1)
	ensureNode(g, node2)
	g[node1][node2] = struct{}{}
	g[node2][node1] = struct{}{}
}

func small(s string) bool {
	return s[0] >= 'a' && s[0] <= 'z'
}

func dfs(g graph, seen stringSet, node string) int {
	if node == "end" {
		return 1
	}

	var acc int

	for n := range g[node] {
		if _, ok := seen[n]; ok && small(n) {
			continue
		}

		acc += dfs(g, seen.copyWith(node), n)
	}

	return acc
}

func main() {
	infile, err := os.Open("12.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	g := make(graph)

	for s.Scan() {
		parts := strings.Split(s.Text(), "-")
		connect(g, parts[0], parts[1])
	}

	fmt.Println(g)

	fmt.Println(dfs(g, make(stringSet), "start"))
}
