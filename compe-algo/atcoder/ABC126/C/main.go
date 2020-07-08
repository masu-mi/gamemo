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
	resolve()
}

func resolve() {
	n, k := scanInt(sc), scanInt(sc)
	rate := 0.0
	for i := 1; i <= n; i++ {
		rate += pWin(i, k)
	}
	rate /= float64(n)
	fmt.Printf("%.10f\n", rate)
}

func pWin(n, k int) float64 {
	cur := n
	rate := 1.0
	for cur < k {
		cur <<= 1
		rate /= 2.0
	}
	return rate
}
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
