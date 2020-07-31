package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	parseProblem()
}

type graph struct {
	size  int
	edges [][]int
	nodes []int
}

func newGraph(n int) *graph {
	return &graph{
		size:  n,
		edges: make([][]int, n),
		nodes: make([]int, n),
	}
}

func parseProblem() {
	n, q := scanInt(sc), scanInt(sc)
	g := newGraph(n)
	for i := 0; i < n-1; i++ {
		a, b := scanInt(sc)-1, scanInt(sc)-1
		g.edges[a] = append(g.edges[a], b)
		g.edges[b] = append(g.edges[b], a)
	}
	for _ = range loop0(q) {
		p, x := scanInt(sc)-1, scanInt(sc)
		g.nodes[p] += x
	}
	nums := make([]int, n)
	visited := make([]bool, n)
	_dfs(nums, visited, g, 0, 0)
	b := strings.Builder{}
	for i := range nums {
		b.WriteString(fmt.Sprint(nums[i]))
		b.WriteString(" ")
	}
	fmt.Println(b.String())
	return
}

func _dfs(nums []int, visited []bool, g *graph, idx, cur int) {
	cur += g.nodes[idx]
	visited[idx] = true
	nums[idx] = cur
	for _, next := range g.edges[idx] {
		if visited[next] {
			continue
		}
		_dfs(nums, visited, g, next, cur)
	}
}

func loop0(n int) chan int {
	return loop(0, n-1, 1)
}

func loop1(n int) chan int {
	return loop(1, n, 1)
}

func loop(s, e, d int) chan int {
	ch := make(chan int)
	go func() {
		for i := s; i <= e; i += d {
			ch <- i
		}
		close(ch)
	}()
	return ch
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
