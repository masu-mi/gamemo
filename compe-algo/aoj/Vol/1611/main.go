package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000000
	maxBufSize     = 100000000
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

var dp [][]int

func initDP(n int) {
	dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < len(dp[0]); j++ {
			dp[i][j] = -1
		}
	}
}

func resolve() {
	ans := []int{}
	for true {
		n := nextInt(sc)
		if n == 0 {
			break
		}
		as := nextIntSlice(sc, n)
		initDP(n)
		_rec(as, n, 0, n)
		ans = append(ans, dp[0][n])
	}
	for _, v := range ans {
		fmt.Println(v)
	}
	return
}

func _rec(as []int, n, l, r int) int {
	if dp[l][r] > -1 {
		return dp[l][r]
	}
	if r-l == 1 { // prevent to hit daruma
		dp[l][r] = 0
		return dp[l][r]
	}
	if r-l == 2 {
		if abs(as[l]-as[r-1]) <= 1 {
			dp[l][r] = 2
		} else {
			dp[l][r] = 0
		}
		return dp[l][r]
	}
	if abs(as[l]-as[r-1]) <= 1 {
		v := _rec(as, n, l+1, r-1)
		// if we can remove internal blocks
		if (r - 1 - (l + 1)) == v {
			dp[l][r] = v + 2
			return dp[l][r]
		} else {
			dp[l][r] = v
		}
	}
	for i := l + 1; i < r; i++ {
		v := _rec(as, n, l, i) + _rec(as, n, i, r)
		if dp[l][r] < v {
			dp[l][r] = v
		}
	}
	return dp[l][r]
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
