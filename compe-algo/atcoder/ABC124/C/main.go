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
	str := scanString(sc)
	startWith0, startWith1 := 0, 0
	for i, c := range str {
		switch i % 2 {
		case 0:
			if c == '0' {
				startWith1++
			} else {
				startWith0++
			}
		case 1:
			if c == '1' {
				startWith1++
			} else {
				startWith0++
			}
		}
	}
	if startWith0 < startWith1 {
		return startWith0
	}
	return startWith1
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
