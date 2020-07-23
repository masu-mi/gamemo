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
	n, p := scanInt(sc), scanInt(sc)
	as := nextIntSlice(sc, n)
	nums := make([]int, 2)
	for _, v := range as {
		nums[v%2]++
	}
	patternNum := 0
	for i := 0; i <= nums[1]; i++ {
		if i%2 != p {
			continue
		}
		v := 1
		for j := 1; j <= i; j++ {
			v *= nums[1] - j + 1
			v /= j
		}
		patternNum += v
	}
	return exp(2, nums[0]) * patternNum
}

func exp(b, e int) int {
	r := 1
	p := b
	c := e
	for c > 0 {
		if c&1 == 1 {
			r *= p
		}
		p = p * p
		c >>= 1
	}
	return r
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
