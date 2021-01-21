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
	fmt.Println(resolve())
}

func resolve() int {
	n, m := scanInt(sc), scanInt(sc)
	connection := map[int]int{}
	for i := 0; i < m; i++ {
		flg := 1 << uint(i)
		k := scanInt(sc)
		for j := 0; j < k; j++ {
			sw := scanInt(sc)
			connection[sw-1] |= flg
		}
	}
	num := 0
	goalState := 0
	for j := 0; j < m; j++ {
		p := scanInt(sc)
		goalState |= p << uint(j)
	}
	for i := 0; i < 1<<uint(n); i++ {
		state := 0
		for idx := 0; idx < n; idx++ {
			if (1<<uint(idx))&i == 0 {
				continue
			}
			state ^= connection[idx]
		}
		if state == goalState {
			num++
		}
	}
	return num
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
