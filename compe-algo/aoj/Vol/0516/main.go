package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	for true {
		n, k := scanInt(sc), scanInt(sc)
		if n == 0 && k == 0 {
			break
		}
		sum := make([]int, n+2)
		for i := 2; i <= n+1; i++ {
			sum[i] = sum[i-1] + scanInt(sc)
		}
		max := math.MinInt32
		for i := 1; i+k <= n+1; i++ {
			if v := sum[i+k] - sum[i]; max < v {
				max = v
			}
		}
		fmt.Println(max)
	}
	return
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
