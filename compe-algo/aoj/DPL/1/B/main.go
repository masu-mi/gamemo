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

type item struct{ v, w int }

func parseProblem() int {
	n, c := scanInt(sc), scanInt(sc)
	items := make([]item, n)
	for i := range loop0(n) {
		items[i] = item{
			v: scanInt(sc),
			w: scanInt(sc),
		}
	}
	dp := make([][]int, n+1) // max value of j<i, w
	dp[0] = make([]int, c+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, c+1)
		for w := 1; w <= c; w++ {
			dp[i][w] = dp[i-1][w]
			item := items[i-1]
			if pW := w - item.w; pW >= 0 {
				if cV := dp[i-1][pW] + item.v; dp[i][w] < cV {
					dp[i][w] = cV
				}
			}
		}
	}
	return dp[n][c]
}

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
