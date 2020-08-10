package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	fmt.Println(resolve(parseProblem()))
}

func parseProblem() (a, b, ab, x, y int) {
	return scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
}

func resolve(a, b, ab, x, y int) int {
	minCost := math.MaxInt64
	for i := 0; i <= max(x, y); i++ {
		v := cost(a, b, ab, 2*i, x, y)
		if v < minCost {
			minCost = v
		}
	}
	return minCost
}

func cost(a, b, ab, n, x, y int) int {
	return ab*n + max(0, x-n/2)*a + max(0, y-n/2)*b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
