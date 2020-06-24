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

func parseProblem() (int, int, int) {
	n, a, b := scanInt(sc), scanInt(sc), scanInt(sc)
	return n, a, b
}

func resolve(n, a, b int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if s := digitSum(i); a <= s && s <= b {
			sum += i
		}
	}
	return sum
}

func digitSum(n int) (sum int) {
	for cur := n; cur > 0; cur /= 10 {
		sum += cur % 10
	}
	return sum
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
