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
	parseProblem()
}

func parseProblem() {
	a, b := scanString(sc), scanString(sc)
	switch lessThan(a, b) {
	case -1:
		fmt.Println("LESS")
	case 0:
		fmt.Println("EQUAL")
	case 1:
		fmt.Println("GREATER")
	}
}

func lessThan(a, b string) int {
	if la, lb := len(a), len(b); la < lb { // no 0 at head
		return -1
	} else if la > lb {
		return 1
	} else {
		for i := 0; i < la; i++ {
			if a[i] < b[i] {
				return -1
			} else if a[i] > b[i] {
				return 1
			}
		}
		return 0
	}
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
