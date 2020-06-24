package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func parseProblem() (int, int, []string) {
	h := scanInt(sc)
	w := scanInt(sc)
	strs := make([]string, h)
	for i := range loop0(h) {
		strs[i] = scanString(sc)
	}
	return h, w, strs
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
func resolve(h, w int, strs []string) {
	b := &strings.Builder{}
	wall(b, w+2)
	b.WriteByte('\n')
	for i := range loop0(h) {
		b.WriteByte('#')
		b.WriteString(strs[i])
		b.WriteString("#\n")
	}
	wall(b, w+2)
	fmt.Println(b.String())
}

func wall(b *strings.Builder, l int) {
	for _ = range loop0(l) {
		b.WriteByte('#')
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
