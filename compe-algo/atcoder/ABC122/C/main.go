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
	n, q := scanInt(sc), scanInt(sc)
	s := scanString(sc)
	culum := make([]int, n+2)
	sum := 0
	for i := 3; i <= n+1; i++ {
		if s[i-2] == 'C' && s[i-3] == 'A' {
			sum++
		}
		culum[i] = sum
	}
	res := []int{}
	for i := 0; i < q; i++ {
		l, r := scanInt(sc), scanInt(sc)
		res = append(res, culum[r+1]-culum[l+1])
	}
	for _, r := range res {
		fmt.Println(r)
	}
	return
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
