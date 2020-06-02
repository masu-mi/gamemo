package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) int {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	return scanInt(sc)
}

const (
	perf = "Perfect"
	def  = "Deficient"
	abu  = "Abundant"
)

func resolve(n int) {
	ed := enumDeci(n)
	sum := 0
	for i := 0; i < len(ed); i++ {
		sum += ed[i]
	}
	if sum == n {
		fmt.Println(perf)
	} else if sum > n {
		fmt.Println(abu)
	} else {
		fmt.Println(def)
	}
	return
}

func enumDeci(n int) (r []int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			r = append(r, i)
			if v := n / i; v != i {
				r = append(r, v)
			}
		}
	}
	sort.Sort(sort.IntSlice(r))
	return r[0 : len(r)-1]
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
