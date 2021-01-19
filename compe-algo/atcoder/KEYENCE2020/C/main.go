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
		var v int
		if s < 1000000000 {
			v = s + 1
		} else {
			v = 1
		}
		fmt.Printf("%d", v)
		for i := 1; i < n; i++ {
			fmt.Printf(" %d", v)
		}
		fmt.Println()
		return
	}
	if s <= n {
		fmt.Printf(" %d", s)
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
	fmt.Printf(" %d", s)
	for i := 1; i < n; i++ {
		if i < k {
			fmt.Printf(" %d", s)
		} else {
			fmt.Printf(" %d", 1)
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
