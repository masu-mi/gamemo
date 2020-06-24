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
	n, d := scanInt(sc), scanInt(sc)
	ds := make([][]int, n)
	for i := range loop0(n) {
		ds[i] = make([]int, d)
		for j := range loop0(d) {
			ds[i][j] = scanInt(sc)
		}
	}
	num := 0
	t := squreNums(16000)
	for i := range loop0(n) {
		for j := i + 1; j < n; j++ {
			sum := 0
			for k := 0; k < d; k++ {
				df := ds[i][k] - ds[j][k]
				sum += df * df
			}
			if t[sum] {
				num++
			}
		}
	}
	return num
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

func squreNums(max int) []bool {
	t := make([]bool, max+1)
	for i := 1; i*i <= max; i++ {
		t[i*i] = true
	}
	return t
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
