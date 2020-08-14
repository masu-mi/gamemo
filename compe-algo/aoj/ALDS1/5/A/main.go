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
	n := scanInt(sc)
	as := nextIntSlice(sc, n)
	q := scanInt(sc)
	answers := make([]bool, q)
	for i := 0; i < q; i++ {
		m := scanInt(sc)
		answers[i] = resolve(as, m)
	}
	for i := 0; i < len(answers); i++ {
		if answers[i] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func resolve(as []int, m int) bool {
	sum := 0
	for i := 0; i < len(as); i++ {
		sum += as[i]
	}
	if sum < m {
		return false
	} else if sum == m {
		return true
	}
	for j := 0; j < 1<<uint(len(as)); j++ {
		if match(as, j, m) {
			return true
		}
	}
	return false
}

func match(as []int, j, m int) bool {
	sum := 0
	for i := 0; i < len(as); i++ {
		if j&(1<<uint(i)) != 0 {
			sum += as[i]
		}
		if sum == m {
			return true
		}
	}
	return false
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
