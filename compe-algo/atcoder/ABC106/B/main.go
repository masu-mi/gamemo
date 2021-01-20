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
	fmt.Println(resolve())
}

func resolve() int {
	n := scanInt(sc)
	var num int
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			continue
		}
		if strangeNumber(i) {
			num++
		}
	}
	return num
}

func strangeNumber(i int) bool {
	var num int
	for j := 1; j*j <= i; j++ {
		if i%j == 0 {
			dv := i / j
			if dv == j {
				num++
			} else {
				num += 2
			}
		}
	}
	return num == 8
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
