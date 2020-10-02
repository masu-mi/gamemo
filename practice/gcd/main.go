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
	a, b := scanInt(sc), scanInt(sc)
	fmt.Printf("gcd(%d, %d): %d\n", a, b, gcd(a, b))
	fmt.Printf("lcm(%d, %d): %d\n", a, b, lcm(a, b))
}

func gcd(a, b int) int {
	for b > 0 {
		t := a / b // a > b or t = 0
		a, b = b, a-t*b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
