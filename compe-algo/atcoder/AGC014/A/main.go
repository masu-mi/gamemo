package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func parseProblem() (int, int, int) {
	a, b, c := scanInt(sc), scanInt(sc), scanInt(sc)
	return a, b, c
}

func resolve(a, b, c int) int {
	return exchangeNum([]int{a, b, c})
}

func allSame(v []int) bool {
	sort.Sort(sort.IntSlice(v))
	return v[0] == v[len(v)-1]
}

func all(v []int, pred func(v int) bool) bool {
	for _, i := range v {
		if !pred(i) {
			return false
		}
	}
	return true
}

func next(v []int) (n []int) {
	n = make([]int, len(v))
	n[0] = (v[0] + v[len(v)-1]) >> 1
	for i := 1; i < len(v); i++ {
		n[i] = (v[i] + v[i-1]) >> 1
	}
	return n
}

func exchangeNum(v []int) int {
	num := 0
	for all(v, func(v int) bool { return v%2 == 0 }) {
		if allSame(v) {
			if v[0]%2 == 0 {
				return -1
			} else {
				return num
			}
		}
		v = next(v)
		num++
	}
	return num
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
