package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	parseProblem(os.Stdin)
}

func parseProblem(r io.Reader) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n := scanInt(sc)
	var limit uint64 = 1
	for i := 0; i < 18; i++ {
		limit *= 10
	}
	var exp float64 = 1
	var res uint64 = 1
	over := false
	for i := 0; i < n; i++ {
		c := scanUint(sc)
		if c == 0 {
			fmt.Println("0")
			return
		}
		exp *= float64(c)
		res *= c
		if res > limit || exp > 1e18 {
			over = true
			res = 1
		}
	}
	if over {
		fmt.Println("-1")
		return
	}
	fmt.Println(res)
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanUint(sc *bufio.Scanner) uint64 {
	sc.Scan()
	u, _ := strconv.ParseUint(sc.Text(), 10, 64)
	return u
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
