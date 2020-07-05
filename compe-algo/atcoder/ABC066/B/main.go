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
	fmt.Println(resolve(scanString(sc)))
}

func resolve(str string) int {
	// O(n^2)
	for i := 2; i < len(str); i += 2 {
		if isEvenStr(str[0 : len(str)-i]) {
			return len(str) - i
		}
	}
	return 0
}

func isEvenStr(str string) bool {
	l := len(str)
	if l%2 == 1 {
		return false
	}
	// O(n)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l/2+i] {
			return false
		}
	}
	return true
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
