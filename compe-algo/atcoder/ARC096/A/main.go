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
	fmt.Println(resolve())
}
func resolve() int {
	a, b, c, x, y := scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
	var price int

	var base int
	if a+b > 2*c {
		base = 2 * c
	} else {
		base = a + b
	}
	minNum, leftNum := min(x, y), max(x, y)-min(x, y)
	price += minNum * base
	var left int
	if x > y {
		left = a
	} else {
		left = b
	}
	if left > 2*c {
		left = 2 * c
	}
	price += leftNum * left
	return price
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
