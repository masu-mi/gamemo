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
	resolve(parseProblem())
}

func parseProblem() (n, q int, d [][]int) {
	n = scanInt(sc)
	d = make([][]int, n)
	for i := range loop0(n) {
		d[i] = make([]int, n)
		for j := range loop0(n) {
			d[i][j] = scanInt(sc)
		}
	}
	q = scanInt(sc)
	return
}

func loop0(n int) chan int {
	return loop(0, n-1, 1)
}

func loop1(n int) chan int {
	return loop(1, n, 1)
}

func loop(s, e, d int) chan int {
	ch := make(chan int)
	go func() {
		for i := s; i <= e; i += d {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func resolve(n, q int, d [][]int) {
	// init
	comu := make([][]int, n+1)
	comu[0] = make([]int, n+1)
	for i := range loop1(n) {
		comu[i] = make([]int, n+1)
		for j := range loop1(n) {
			comu[i][j] = d[i-1][j-1] +
				comu[i][j-1] + comu[i-1][j] - comu[i-1][j-1]
		}
	}
	// summarize for size
	sum := map[int]int{}
	for xl := 1; xl < n+1; xl++ {
		for xh := xl; xh < n+1; xh++ {
			for yl := 1; yl < n+1; yl++ {
				for yh := yl; yh < n+1; yh++ {
					// h == l -> 1
					size := (xh + 1 - xl) * (yh + 1 - yl)
					// 1..n -> lower part should get -1
					v := comu[xh][yh] -
						comu[xh][yl-1] -
						comu[xl-1][yh] +
						comu[xl-1][yl-1]
					sum[size] = max(sum[size], v)
				}
			}
		}
	}
	for range loop0(q) {
		s := scanInt(sc)
		m := 0
		for i := 1; i <= s; i++ {
			m = max(m, sum[i])
		}
		fmt.Println(m)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
