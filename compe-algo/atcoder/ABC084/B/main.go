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

func inNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseProblem() {
	a, b := scanInt(sc), scanInt(sc)
	s := scanString(sc)
	for i := 0; i < a; i++ {
		if !inNum(s[i]) {
			fmt.Println("No")
			return
		}
	}
	if s[a] != '-' {
		fmt.Println("No")
		return
	}
	for i := 0; i < b; i++ {
		if !inNum(s[a+i+1]) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
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
