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

func parseProblem() {
	x1, y1, x2, y2 := scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
	base := vec{x: x1, y: y1}
	dv := vec{x: x2 - x1, y: y2 - y1}
	orDv := orthogonal(dv)
	p3 := add(add(base, dv), orDv)
	p4 := add(base, orDv)
	fmt.Println(p3.x, p3.y, p4.x, p4.y)
}

type vec struct{ x, y int }

func add(v1, v2 vec) vec {
	return vec{x: v1.x + v2.x, y: v1.y + v2.y}
}

func orthogonal(v vec) vec {
	return vec{x: -v.y, y: v.x}
}

func resolve(n int) int {
	return n
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
