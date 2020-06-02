package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(resolve(parseProblem(os.Stdin)))
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

func resolve(n int) (int, int) {
	ds := enumDivisors(n)
	max := len(ds) - 1

	cand := map[int]struct{}{}
	for _, d := range ds {
		cand[d] = struct{}{}
	}
	delete(cand, n)
	sort.Sort(sort.Reverse(sort.IntSlice(ds)))
	min := 0
	for _, k := range ds {
		if _, ok := cand[k]; !ok {
			continue
		}
		if len(cand) == 0 {
			break
		}
		for _, kk := range enumDivisors(k) {
			delete(cand, kk)
		}
		min++
	}
	return min, max
}

func pow(n, i int) (res int) {
	res = 1
	for i > 0 {
		if i&1 == 1 {
			res *= n
		}
		i >>= 1
		n *= n
	}
	return res
}

func enumDivisors(n int) (r []int) {
	max := int(math.Sqrt(float64(n))) + 1
	for i := 1; i < max; i++ {
		if n%i == 0 {
			r = append(r, i)
			if n/i != i {
				r = append(r, n/i)
			}
		}
	}
	sort.Sort(sort.IntSlice(r))
	return r
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
