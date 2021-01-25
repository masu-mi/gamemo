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

func parseProblem() []string {
	n := scanInt(sc)
	bs := make([]string, n)
	for i := 0; i < n; i++ {
		bs[i] = scanString(sc)
	}
	return bs
}

func resolve(bs []string) int {
	return resolvedNumber(bs, len(bs)-1)
}

func resolvedNumber(bs []string, idx int) int {
	if idx == 0 {
		if bs[idx] == "OR" {
			return 3
		}
		return 1
	}
	if bs[idx] == "OR" {
		return 1<<uint(idx+1) + resolvedNumber(bs, idx-1)
	}
	return resolvedNumber(bs, idx-1)
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
