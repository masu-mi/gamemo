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

func parseProblem() (int, int) {
	return scanInt(sc), scanInt(sc)
}

func resolve(l, r int) int {
	df := r - l
	if df+1 >= 2019 {
		return 0
	}
	min := 2019
	for i := 0; i < df; i++ {
		for j := i + 1; j <= df; j++ {
			v := (l + i) * (l + j) % 2019
			if min > v {
				min = v
			}
			if v == 0 {
				return 0
			}
		}
	}
	return min
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
