package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	n, c, k := scanInt(sc), scanInt(sc), scanInt(sc)
	ts := nextIntSlice(sc, n)
	fmt.Println(resolve(n, c, k, ts))
}

func resolve(n, c, k int, ts []int) int {
	sort.Sort(sort.IntSlice(ts))
	var idx, trackNum int

	limit := ts[0] + k
	idx, sum := 0, 1
	for i := 1; i < len(ts); i++ {
		if sum == c {
			sum = 0
			trackNum++
			idx += c
			limit = ts[idx] + k
		} else if limit < ts[i] {
			sum = 0
			trackNum++
			idx = i
			limit = ts[idx] + k
		}
		sum++
	}
	trackNum += (sum + c - 1) / c
	return trackNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
