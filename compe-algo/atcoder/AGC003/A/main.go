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
	parseProblem()
}

func parseProblem() {
	s := scanString(sc)
	ex := map[byte]struct{}{}
	for i := range s {
		ex[s[i]] = struct{}{}
	}
	_, nOk := ex['N']
	_, sOk := ex['S']
	if nOk != sOk {
		fmt.Println("No")
		return
	}
	_, wOk := ex['W']
	_, eOk := ex['E']
	if wOk != eOk {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
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
