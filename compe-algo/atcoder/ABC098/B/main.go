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

func parseProblem() int {
	_ = scanInt(sc)
	str := scanString(sc)
	max := 0
	for i := 0; i < len(str); i++ {
		lsets := map[byte]struct{}{}
		rsets := map[byte]struct{}{}
		num := 0
		for j := 0; j < len(str); j++ {
			if j < i {
				lsets[str[j]] = struct{}{}
			} else {
				if _, ok := lsets[str[j]]; ok {
					if _, ok := rsets[str[j]]; !ok {
						rsets[str[j]] = struct{}{}
						num++
					}
				}
			}
			if max < num {
				max = num
			}
		}
	}
	return max
}

func resolve(n int) int {
	return n
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
