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
	fmt.Println(resolve(parseProblem()))
}

func parseProblem() (a, b, ab, x, y int) {
	return scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
}

func resolve(a, b, ab, x, y int) int {
	minCost := 0
	if 2*ab < a+b {
		minNum := min(x, y)
		x -= minNum
		y -= minNum
		minCost += ab * 2 * minNum
	}
	if x > 0 {
		if ab*2 < a {
			minCost += ab * 2 * x
		} else {
			minCost += a * x
		}
	}
	if y > 0 {
		if ab*2 < b {
			minCost += ab * 2 * y
		} else {
			minCost += b * y
		}
	}
	return minCost
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
