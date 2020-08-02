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

func parseProblem() (int, []int) {
	n := scanInt(sc)
	as := nextIntSlice(sc, n)
	return n, as
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return int(a)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextIntSlice(sc *bufio.Scanner, n int) (a []int) {

	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
	}
	return a
}
func resolve(n int, as []int) int {
	dp = make([][]int, len(as))
	for i := 0; i < len(as); i++ {
		dp[i] = make([]int, len(as))
		for j := 0; j < len(as); j++ {
			dp[i][j] = -1
			dp[i][j] = -1
		}
	}
	return _resolve(as, 0, len(as)-1)
}

var dp [][]int

func _resolve(as []int, l, r int) int {
	if dp[l][r] != -1 {
		return dp[l][r]
	}
	if l == r {
		dp[l][r] = as[l]
		return dp[l][r]
	}
	dp[l][r] = as[l] - _resolve(as, l+1, r)
	cand := as[r] - _resolve(as, l, r-1)
	if dp[l][r] < cand {
		dp[l][r] = cand
	}
	return dp[l][r]
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
