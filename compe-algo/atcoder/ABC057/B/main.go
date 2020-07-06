package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	n := scanInt(sc)
	m := scanInt(sc)
	ps := make([]pos, n)
	cs := make([]pos, m)
	for i := range loop0(n) {
		ps[i] = pos{x: scanInt(sc), y: scanInt(sc)}
	}
	for i := range loop0(m) {
		cs[i] = pos{x: scanInt(sc), y: scanInt(sc)}
	}
	for _, j := range ps {
		min := math.MaxInt64
		minIdx := -1
		for i, c := range cs {
			v := dist(c, j)
			if min > v {
				min = v
				minIdx = i + 1
			}
		}
		fmt.Println(minIdx)
	}
	return
}

func dist(x, y pos) int {
	return abs(x.x-y.x) + abs(x.y-y.y)
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

type pos struct{ x, y int }

func loop0(n int) chan int {
	return loop(0, n-1, 1)
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
