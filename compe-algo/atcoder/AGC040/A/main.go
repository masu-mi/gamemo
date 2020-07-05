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

func parseProblem() string {
	return scanString(sc)
}

func resolve(str string) int {
	lMax, r := 0, 0
	for i := 0; i < len(str); {
		h, l := str[i], 0
		{
			j := i
			for ; j < len(str); j++ {
				if str[j] != h {
					break
				}
				l++
			}
			i = j
		}
		r += (l - 1) * l >> 1
		switch h {
		case '>':
			r += max(lMax, l)
			lMax = 0
		case '<':
			lMax = l
		}
	}
	r += lMax
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
