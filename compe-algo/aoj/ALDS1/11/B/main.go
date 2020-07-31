package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 1000000
)

var sc *bufio.Scanner

func initScanner(r io.Reader) *bufio.Scanner {
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	return sc
}

func main() {
	sc = initScanner(os.Stdin)
	resolve(parseProblem())
}

type graph struct {
	size  int
	edges [][]int
}

func parseProblem() *graph {
	n := scanInt(sc)
	g := &graph{
		size:  n,
		edges: make([][]int, n),
	}
	for i := 0; i < g.size; i++ {
		_ = scanInt(sc)
		k := scanInt(sc)
		g.edges[i] = make([]int, k)
		for j := 0; j < k; j++ {
			// 0-indexedに統一した方がバグが少ない
			g.edges[i][j] = scanInt(sc) - 1
		}
	}
	return g
}

func resolve(g *graph) {
	ds := make([]int, g.size)
	fs := make([]int, g.size)
	time := 0
	for i := 0; i < len(ds); i++ {
		if ds[i] == 0 {
			time++
			time = _dfs(ds, fs, time, g, i)
		}
	}
	// print answer
	for i := 0; i < len(ds); i++ {
		fmt.Println(i+1, ds[i], fs[i])
	}
	return
}

func _dfs(ds, fs []int, time int, g *graph, cur int) int {
	ds[cur] = time
	for _, idx := range g.edges[cur] {
		if ds[idx] > 0 {
			continue
		}
		time++
		time = _dfs(ds, fs, time, g, idx)
	}
	time++
	fs[cur] = time
	return time
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
