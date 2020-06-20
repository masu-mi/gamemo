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
	parseProblem()
}

func parseProblem() {
	n := scanInt(sc)
	orders := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		orders[scanInt(sc)] = i
	}
	buf := &strings.Builder{}
	buf.WriteString(fmt.Sprintf("%d", orders[1]))
	for i := 2; i <= n; i++ {
		buf.WriteString(" ")
		buf.WriteString(fmt.Sprintf("%d", orders[i]))
	}
	fmt.Println(buf.String())
	return
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
