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

func parseProblem() scorer {
	d := scanInt(sc)
	v := 26
	cs := make([]int, v)
	for i := 0; i < v; i++ {
		cs[i] = scanInt(sc)
	}
	ss := make([][]int, d)
	for i := range loop0(d) {
		ss[i] = make([]int, v)
		for j := range loop0(v) {
			ss[i][j] = scanInt(sc)
		}
	}
	return scorer{
		d: d, v: v,
		cs: cs, ss: ss,
	}
}

type scorer struct {
	// 0-indexed
	d, v int
	cs   []int
	ss   [][]int // [<d][<v]
}

func (scr scorer) receiveInput(sc *bufio.Scanner) []int {
	satisfaction := 0
	lasts := make([]int, scr.v)
	results := make([]int, scr.d)
	for day := range loop0(scr.d) {
		tipe := scanInt(sc) - 1
		satisfaction += scr.ss[day][tipe]
		lasts[tipe] = day + 1
		for cType := range loop0(scr.v) {
			satisfaction -= scr.cs[cType] * (day + 1 - lasts[cType])
		}
		results[day] = satisfaction
	}
	return results
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

func resolve(scor scorer) {
	for _, r := range scor.receiveInput(sc) {
		fmt.Println(r)
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
