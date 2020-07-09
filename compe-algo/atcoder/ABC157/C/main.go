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
	resolve()
}

func resolve() {
	n, m := scanInt(sc), scanInt(sc)
	sToC := map[int]int{}
	for i := 0; i < m; i++ {
		s := scanInt(sc)
		c := scanInt(sc)
		if v, ok := sToC[s]; ok && v != c {
			fmt.Println(-1)
			return
		}
		sToC[s] = c
	}
	for i := 1; i <= n; i++ {
		if c, ok := sToC[i]; ok {
			if i == 1 && n != 1 && c == 0 {
				fmt.Println(-1)
				return
			}
			fmt.Printf("%d", c)
		} else if n != 1 && i == 1 {
			fmt.Printf("%d", 1)
		} else {
			fmt.Printf("%d", 0)
		}
	}
	fmt.Println()
	return
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
