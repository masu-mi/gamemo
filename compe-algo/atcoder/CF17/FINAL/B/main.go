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
	resolve(parseProblem())
}

func parseProblem() string {
	return scanString(sc)
}

func resolve(s string) {
	nums := map[byte]int{'a': 0, 'b': 0, 'c': 0}
	for i := 0; i < len(s); i++ {
		nums[s[i]]++
	}
	max, min := 0, math.MaxInt32
	for _, v := range nums {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	if max-min <= 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	return
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
