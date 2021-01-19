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
	resolve(parseProblem())
}

func parseProblem() (int, int, int) {
	n, k, s := scanInt(sc), scanInt(sc), scanInt(sc)
	return n, k, s
}

func resolve(n, k, s int) {
	if k == 0 {
		fmt.Printf("%d", s+1)
		for i := 1; i < n; i++ {
			fmt.Printf(" %d", s+1)
		}
		fmt.Println()
		return
	}
	segmentSize := n - k + 1
	if s < 2*segmentSize {
		fmt.Printf("%d", s)
		for i := 1; i < n; i++ {
			if i < k {
				fmt.Printf(" %d", s)
			} else {
				fmt.Printf(" %d", s+1)
			}
		}
		fmt.Println()
		return
	}
	ordVal := s / segmentSize
	var lVal, rVal int
	{
		df := s % segmentSize
		if df == 0 {
			lVal, rVal = ordVal-1, ordVal+1
		} else {
			lVal, rVal = ordVal, ordVal+df
		}
	}
	fmt.Printf("%d", lVal)
	for i := 1; i < n; i++ {
		if rem := i % segmentSize; rem == 0 {
			fmt.Printf(" %d", lVal)
		} else if rem == segmentSize-1 {
			fmt.Printf(" %d", rVal)
		} else {
			fmt.Printf(" %d", ordVal)
		}
	}
	fmt.Println()
	return
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
