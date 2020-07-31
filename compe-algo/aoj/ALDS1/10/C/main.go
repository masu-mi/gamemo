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
	resolve()
}

func resolve() {
	n := scanInt(sc)
	results := []int{}
	for i := 0; i < n; i++ {
		results = append(results, resolveInstance())
	}
	for _, v := range results {
		fmt.Println(v)
	}
	return
}

func resolveInstance() int {
	x, y := scanString(sc), scanString(sc)
	lx, ly := len(x), len(y)
	dp := make([][]int, lx+1)
	dp[0] = make([]int, ly+1)
	for i := 1; i <= lx; i++ {
		dp[i] = make([]int, ly+1)
		dp[i][0] = dp[i-1][0]
		for j := 1; j <= ly; j++ {
			dp[i][j] = dp[i-1][j] // I forgot this case
			if dp[i][j] < dp[i][j-1] {
				dp[i][j] = dp[i][j-1]
			}
			if x[i-1] == y[j-1] {
				cL := dp[i-1][j-1] + 1
				if dp[i][j] < cL {
					dp[i][j] = cL
				}
			}
		}
	}
	return dp[lx][ly]
}

const (
	remove = 0 + iota
	match
)

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
