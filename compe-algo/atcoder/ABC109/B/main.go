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
	var last byte
	used := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		w := scanString(sc)
		if _, ok := used[w]; ok {
			fmt.Println("No")
			return
		}
		used[w] = struct{}{}
		if i > 0 {
			if last != w[0] {
				fmt.Println("No")
				return
			}
		}
		last = w[len(w)-1]
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
