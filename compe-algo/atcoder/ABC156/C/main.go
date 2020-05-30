package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	parseProblem(os.Stdin)
}

func parseProblem(r io.Reader) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n := scanInt(sc)
	xs := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		xs[i] = scanInt(sc)
		sum += xs[i]
	}
	avgL := sum / n
	costL := 0
	avgR := avgL + 1
	costR := 0
	for _, v := range xs {
		costL += (v - avgL) * (v - avgL)
		costR += (v - avgR) * (v - avgR)
	}
	fmt.Println(min(costL, costR))
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
