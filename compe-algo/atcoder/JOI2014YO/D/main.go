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

func parseProblem() (int, string) {
	n := scanInt(sc)
	str := scanString(sc)
	return n, str
}

func include(set int, name byte) bool {
	m := map[byte]uint{'J': 0, 'O': 1, 'I': 2}
	return set&(1<<m[name]) != 0
}

func resolve(n int, str string) int {
	// O(2^n*day)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 1<<uint(3))
	}
	dp[0][1] = 1
	for i := 1; i < n+1; i++ {
		for j := 1; j < 1<<uint(3); j++ {
			if !include(j, str[i-1]) {
				continue
			}
			for k := 1; k < 1<<uint(3); k++ {
				if k&j != 0 {
					dp[i][j] = (dp[i][j] + dp[i-1][k]) % 10007
				}
			}
		}
	}
	sum := 0
	for _, v := range dp[len(str)] {
		sum = (sum + v) % 10007
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
