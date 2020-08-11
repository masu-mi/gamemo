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
	fmt.Println(resolve(parseProblem()))
}

func parseProblem() (int, []int, []int) {
	n := scanInt(sc)
	as := make([]int, n)
	bs := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = scanInt(sc)
		bs[i] = scanInt(sc)
	}
	return n, as, bs
}

func findMedian(as []int) int {
	sort.Sort(sort.IntSlice(as))
	l := len(as)
	if l%2 == 0 {
		return (as[l/2] + as[l/2-1]) / 2
	}
	return as[l/2]
}

func resolve(n int, as, bs []int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += abs(as[i] - bs[i])
	}
	return minTime(as) + sum + minTime(bs)
}

func minTime(as []int) int {
	sort.Sort(sort.IntSlice(as))
	l := len(as)
	medA, medB := as[l/2], as[l/2-1]
	candA, candB := 0, 0
	for i := 0; i < len(as); i++ {
		candA += abs(medA - as[i])
		candB += abs(medB - as[i])
	}
	return min(candA, candB)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
