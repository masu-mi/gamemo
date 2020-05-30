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
	n, a, b := scanInt(sc), scanInt(sc), scanInt(sc)
	s := scanString(sc)
	rank, rankB := 0, 0
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'c':
			fmt.Println("No")
			continue
		case 'a':
			if rank < a+b {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
				continue
			}
		case 'b':
			if rank < a+b && rankB < b {
				fmt.Println("Yes")
				rankB++
			} else {
				fmt.Println("No")
				continue
			}
		}
		rank++
	}
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
