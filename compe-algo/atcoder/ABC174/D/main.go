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

func parseProblem() (int, string) {
	n := scanInt(sc)
	s := scanString(sc)
	return n, s
}

func resolve(n int, cs string) int {
	num := 0
	li, ri := 0, n
	for ri-li > 1 {
		for cs[li] == 'R' && li+1 < ri {
			li++
		}
		if li+1 == ri {
			break
		}
		for cs[ri-1] == 'W' && li+1 < ri {
			ri--
		}
		if li+1 == ri {
			break
		}
		num++
		li++
		ri--
	}
	return num
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
