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

func parseProblem() int {
	return scanInt(sc)
}

func resolve(n int) int {
	if n < 600 {
		return 8
	} else if n < 800 {
		return 7
	} else if n < 1000 {
		return 6
	} else if n < 1200 {
		return 5
	} else if n < 1400 {
		return 4
	} else if n < 1600 {
		return 3
	} else if n < 1800 {
		return 2
	} else if n < 2000 {
		return 1
	}
	panic(-1)
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
