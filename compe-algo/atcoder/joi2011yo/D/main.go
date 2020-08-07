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

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/next.go] with goone.

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
	dp := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		dp[i] = make([]int, 21)
	}
	dp[0][as[0]] = 1
	for i := 1; i < n-1; i++ {
		for j := 0; j <= 20; j++ {
			k := j + as[i]
			if 0 <= k && k <= 20 {
				dp[i][k] += dp[i-1][j]
			}
			k = j - as[i]
			if 0 <= k && k <= 20 {
				dp[i][k] += dp[i-1][j]
			}
		}
	}
	return dp[n-2][as[n-1]]
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/simple_loop.go] with goone.

func loop0(n int) chan int {
	return loop(0, n-1, 1)
}

func loop1(n int) chan int {
	return loop(1, n, 1)
}

func loop(s, e, d int) chan int {
	ch := make(chan int)
	go func() {
		for i := s; i <= e; i += d {
			ch <- i
		}
		close(ch)
	}()
	return ch
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
