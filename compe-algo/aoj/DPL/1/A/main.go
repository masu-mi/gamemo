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

type coin struct{ v, w int }

func parseProblem() int {
	n, m := scanInt(sc), scanInt(sc)
	coins := make([]coin, m)
	for i := 0; i < m; i++ {
		coins[i] = coin{
			v: scanInt(sc),
			w: 1,
		}
	}
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= m; i++ {
		for v := 0; v <= n; v++ {
			it := coins[i-1]
			cv := v - it.v
			if cv >= 0 {
				cw := dp[cv] + it.w
				if dp[v] > cw {
					dp[v] = cw
				}
			}
		}
	}
	return dp[n]
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
