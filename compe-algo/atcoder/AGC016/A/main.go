package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	secSizes := findSectionSizeForByte(str)
	m := math.MaxInt64
	for _, s := range secSizes {
		if s < m {
			m = s
		}
	}
	return m
}

func findSectionSizeForByte(str string) map[byte]int {
	// section is byte sequence which don't involved the target byte.
	maxSec := map[byte]int{}
	lastPos := map[byte]int{}
	for i := 0; i < len(str); i++ {
		b := str[i]
		if lp, ok := lastPos[b]; !ok {
			maxSec[str[i]] = i
		} else {
			ls := i - lp - 1
			if maxSec[b] < ls {
				maxSec[b] = ls
			}
		}
		lastPos[str[i]] = i
	}
	for b, lp := range lastPos {
		ls := len(str) - lp - 1
		if ls > maxSec[b] {
			maxSec[b] = ls
		}
	}
	return maxSec
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
