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

func parseProblem() int {
	n, m := scanInt(sc), scanInt(sc)
	tb := make([][]int, m)
	for i := 0; i < m; i++ {
		tb[i] = make([]int, n)
		l := scanInt(sc)
		for j := 0; j < l; j++ {
			tb[i][scanInt(sc)-1] = 1
		}
	}
	ps := make([]int, m)
	for i := 0; i < m; i++ {
		ps[i] = scanInt(sc)
	}
	num := 0
	for i := 0; i < 1<<uint(n); i++ {
		if accept(n, m, tb, ps, i) {
			num++
		}
	}
	return num
}

func accept(n, m int, tb [][]int, ps []int, idx int) bool {
	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if idx&(1<<uint(j)) == 0 {
				continue
			}
			sum += tb[i][j]
		}
		if ps[i] != sum%2 {
			return false
		}
	}
	return true
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
