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
	n := scanInt(sc)
	as := nextIntSlice(sc, n)
	exNums := make([]int, 3)
	for _, v := range as {
		ex := exOf(v, 2)
		exNums[min(2, ex)]++
	}
	if exNums[0]+exNums[1]%2 > exNums[2]+1 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func exOf(n, p int) (ex int) {
	c := n
	for c%p == 0 {
		ex++
		c /= p
	}
	return ex
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return int(a)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextIntSlice(sc *bufio.Scanner, n int) (a []int) {

	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
	}
	return a
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
