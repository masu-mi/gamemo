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
	fmt.Println(resolve())
}

func resolve() int {
	n := scanInt(sc)
	minV := math.MaxInt64
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			div := n / i
			v := fVal(i, div)
			if minV > v {
				minV = v
			}
		}
	}
	return minV
}

func fVal(a, b int) int {
	aL, bL := len(strconv.Itoa(a)), len(strconv.Itoa(b))
	if aL < bL {
		return bL
	}
	return aL
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
