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
	sc.Split(bufio.ScanLines) // bufio.ScanLines
	return sc
}

func main() {
	sc = initScanner(os.Stdin)
	parseProblem()
}

func parseProblem() {
	num := scanInt(sc)
	ps := make([]string, num)
	for i := 0; i < num; i++ {
		ps[i] = scanString(sc)
	}
	text := scanString(sc)
	a := newPMA(ps)
	fmt.Printf("%s\n", text)
	for p := range a.searchPatterns(text) {
		fmt.Printf("%s\n", resultToString(p))
	}
}

func resultToString(r result) string {
	return fmt.Sprintf("%s%s", strings.Repeat(" ", r.start), r.pattern)
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
