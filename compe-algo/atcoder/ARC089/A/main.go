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

func parseProblem() []request {
	n := scanInt(sc)
	cs := make([]request, n)
	for i := 0; i < n; i++ {
		cs[i] = request{
			t: scanInt(sc),
			x: scanInt(sc),
			y: scanInt(sc),
		}
	}
	return cs
}

type request struct {
	t, x, y int
}

func dist(x, y request) int {
	return abs(x.x-y.x) + abs(x.y-y.y)
}

func isReachable(s, t request) bool {
	l := t.t - s.t
	d := dist(s, t)
	if d <= l && d%2 == l%2 {
		return true
	}
	return false
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func resolve(cs []request) {
	cur := request{t: 0, x: 0, y: 0}
	for i := 0; i < len(cs); i++ {
		if !isReachable(cur, cs[i]) {
			fmt.Println("No")
			return
		}
		cur = cs[i]
	}
	fmt.Println("Yes")
	return
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
