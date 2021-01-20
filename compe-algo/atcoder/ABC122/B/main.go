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
	acceptable := map[byte]struct{}{'A': {}, 'C': {}, 'G': {}, 'T': {}}
	maxLen, l := 0, 0
	for i := 0; i < len(str); i++ {
		if _, ok := acceptable[str[i]]; ok {
			l++
		} else {
			if maxLen < l {
				maxLen = l
			}
			l = 0
		}
	}
	if maxLen < l {
		maxLen = l
	}
	return maxLen
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
