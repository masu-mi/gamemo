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

func parseProblem() (int, int, []int, []clothes) {
	d, n := scanInt(sc), scanInt(sc)
	ds := make([]int, d)
	for i := range loop0(d) {
		ds[i] = scanInt(sc)
	}
	cs := make([]clothes, n)
	for i := range loop0(n) {
		cs[i] = clothes{
			l: scanInt(sc),
			h: scanInt(sc),
			c: scanInt(sc),
		}
	}
	return d, n, ds, cs
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/simple_loop.go] with goone.

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

type clothes struct {
	l, h, c int
}

func (c clothes) accept(t int) bool {
	return c.l <= t && t <= c.h
}

func resolve(d, n int, ds []int, cs []clothes) int {
	dp := make([]map[clothes]int, d)
	for i := range loop0(d) {
		dp[i] = map[clothes]int{}
	}
	for _, c := range cs {
		if c.accept(ds[0]) {
			dp[0][c] = 0
		}
	}

	for i := 1; i < d; i++ {
		for _, c := range cs {
			if !c.accept(ds[i]) {
				continue
			}
			for pc, score := range dp[i-1] {
				v := abs(pc.c-c.c) + score
				if cur := dp[i][c]; cur <= v {
					dp[i][c] = v
				}
			}
		}
	}
	max := 0
	for _, score := range dp[d-1] {
		if max < score {
			max = score
		}
	}
	return max
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
