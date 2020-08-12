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
	parseProblem()
}

type pair struct{ x, y int }

func (p pair) inv() pair {
	return pair{x: -p.x, y: -p.y}
}
func (p pair) String() string {
	return fmt.Sprintf("%d %d", p.x, p.y)
}

func add(x, y pair) pair {
	return pair{x: x.x + y.x, y: x.y + y.y}
}

func sub(x, y pair) pair {
	return add(x, y.inv())
}

func next() pair {
	return pair{
		x: scanInt(sc),
		y: scanInt(sc),
	}
}

func parseProblem() {
	m := scanInt(sc)
	base := next()
	dfs := make([]pair, m-1)
	for i := 0; i < m-1; i++ {
		dfs[i] = sub(next(), base)
	}
	n := scanInt(sc)
	stars := map[pair]struct{}{}
	for i := 0; i < n; i++ {
		stars[next()] = struct{}{}
	}
external:
	for start := range stars {
		for _, df := range dfs {
			other := add(start, df)
			if _, ok := stars[other]; !ok {
				continue external
			}
		}
		fmt.Println(sub(start, base))
		return
	}
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
