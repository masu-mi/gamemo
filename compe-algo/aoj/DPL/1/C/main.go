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
	for i := 0; i < n; i++ {
		items[i] = item{
			v: scanInt(sc),
			w: scanInt(sc),
		}
	}
	dp := make([]int, c+1)
	for i := 1; i <= n; i++ {
		for w := 1; w <= c; w++ {
			it := items[i-1]
			if cw := w - it.w; 0 <= cw {
				if cv := dp[cw] + it.v; dp[w] <= cv {
					dp[w] = dp[cw] + it.v
				}
			}
		}
	}
	return dp[c]
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
