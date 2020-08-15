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
	tbl := make([][]int, n)
	for i := 0; i < n; i++ {
		tbl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			tbl[i][j] = scanInt(sc)
		}
	}
	max := 0
	for i := 0; i < 1<<uint(n); i++ {
		if v := countNum(n, m, tbl, i); max < v {
			max = v
		}
	}
	return max
}

func countNum(n, m int, tbl [][]int, idx int) (num int) {
	for j := 0; j < m; j++ {
		v := [2]int{}
		for i := 0; i < n; i++ {
			val := tbl[i][j]
			if (idx>>uint(i))&1 == 0 {
				v[val]++
				continue
			}
			switch val {
			case 0:
				v[1]++
			case 1:
				v[0]++
			}
		}
		num += max(v[0], v[1])
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
