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
	fmt.Println(resolve())
}

func resolve() int {
	a, b, c, x, y := scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
	result := math.MaxInt64
	for i := 0; i <= max(x, y); i++ {
		v := cost(a, b, c, 2*i, x, y)
		if result > v {
			result = v
		}
	}
	return result
}

func cost(a, b, c, n, x, y int) int {
	return c*n + a*max(0, x-(n/2)) + b*max(0, y-(n/2))
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/nums.go] with goone.

func relu(n int) int {
	return max(0, n)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
