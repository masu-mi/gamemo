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

func parseProblem() (n, m, k int, as, bs []int) {
	n, m, k = scanInt(sc), scanInt(sc), scanInt(sc)
	as, bs = nextIntSlice(sc, n), nextIntSlice(sc, m)
	return
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return int(a)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextIntSlice(sc *bufio.Scanner, n int) (a []int) {

	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
	}
	return a
}

func nextIdx(sumA []int, k, start int) int {
	idx := start
	for ; idx >= 0; idx-- {
		if sumA[idx] <= k {
			break
		}
	}
	return idx
}

func resolve(n, m, k int, as, bs []int) int {
	sumA := make([]int, len(as)+1)
	for i := 1; i <= len(as); i++ {
		sumA[i] = sumA[i-1] + as[i-1]
	}
	idx := nextIdx(sumA, k, len(sumA)-1)
	max := idx
	for j := 1; j <= len(bs); j++ {
		k -= bs[j-1]
		if k < 0 {
			break
		}
		idx = nextIdx(sumA, k, idx)
		if max < j+idx {
			max = j + idx
		}
	}
	return max
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
