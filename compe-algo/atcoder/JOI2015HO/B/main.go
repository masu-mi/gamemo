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
	initDP(n)
	_rec(as, n, 0, n)
	return dp[0][n]
}

func _rec(as []int, n, l, r int) int {
	if v := dp[l%n][r%n]; v > 0 {
		return v
	}
	turn := n - (r+n-l)%n
	if (l+1)%n == r%n {
		switch turn % 2 {
		case 0:
			return as[l%n]
		case 1:
			return 0
		default:
			panic(-1)
		}
	}
	if l%n == r%n {
		maxVal := 0
		for i := 0; i < n; i++ {
			v := _rec(as, n, (l+i+1)%n, (l+i)%n) + as[(l+i)%n]
			if maxVal < v {
				maxVal = v
			}
		}
		dp[l][r] = maxVal
		return dp[l][r]
	}
	switch turn % 2 {
	case 0:
		dp[l][r] = max(
			_rec(as, n, (l+1)%n, r%n)+as[l%n],
			_rec(as, n, l%n, (r-1+n)%n)+as[(r-1+n)%n],
		)
	case 1:
		if as[l%n] < as[(r-1+n)%n] {
			dp[l][r] = _rec(as, n, l%n, (r-1+n)%n)
		} else {
			dp[l][r] = _rec(as, n, (l+1)%n, r%n)
		}
	default:
		panic(-1)
	}
	return dp[l][r]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func initDP(n int) {
	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
}

var dp [][]int

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
