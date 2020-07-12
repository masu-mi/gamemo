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
	fmt.Printf("%.7f\n", resolve(parseProblem()))
}

func parseProblem() (int, int) {
	n, k := scanInt(sc), scanInt(sc)
	return n, k
}

func resolve(n, k int) float64 {
	sumOfExpectation := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		sumOfExpectation[i] = sumOfExpectation[i-1] + expectation(scanInt(sc))
	}
	max := 0.0
	for i := 0; i+k < len(sumOfExpectation); i++ {
		v := sumOfExpectation[i+k] - sumOfExpectation[i]
		if max < v {
			max = v
		}
	}
	return max
}

func expectation(n int) float64 {
	exp := (n + 1) * n >> 1
	return float64(exp) / float64(n)
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
