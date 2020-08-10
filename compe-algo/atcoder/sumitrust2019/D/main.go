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
	n, s := scanInt(sc), scanString(sc)
	return n, s
}

func resolve(n int, str string) int {
	num := 0
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if accept(str, i, j, k) {
					num++
				}
			}
		}
	}
	return num
}

func accept(str string, i, j, k int) bool {
	ans := []byte{
		'0' + byte(i),
		'0' + byte(j),
		'0' + byte(k),
	}
	ansIdx := 0
	for idx := 0; idx < len(str); idx++ {
		if ansIdx == len(ans) {
			return true
		}
		if ans[ansIdx] == str[idx] {
			ansIdx++
		}
	}
	return ansIdx == len(ans)
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
