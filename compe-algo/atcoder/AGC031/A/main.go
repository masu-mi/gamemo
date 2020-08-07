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

const modulo = 1e9 + 7

func resolve(n int, str string) int {
	nums := map[byte]int{}
	for i := 0; i < len(str); i++ {
		nums[str[i]]++
	}
	num := 1
	for _, v := range nums {
		num = (num * (v + 1)) % modulo
	}
	return (num - 1) % modulo
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
