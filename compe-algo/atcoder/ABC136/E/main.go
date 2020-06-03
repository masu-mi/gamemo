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

func parseProblem(r io.Reader) (int, int, []int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n, k := scanInt(sc), scanInt(sc)
	as := []int{}
	sum := 0
	for i := 0; i < n; i++ {
		as = append(as, scanInt(sc))
		sum += as[i]
	}
	return sum, k, as
}

func resolve(sum, k int, as []int) int {
	divs := enumDivisors(sum)
	sort.Sort(sort.Reverse(sort.IntSlice(divs)))
	for _, d := range divs {
		rests := make([]int, len(as))
		for i, a := range as {
			rests[i] = a % d
		}
		sort.Sort(sort.IntSlice(rests))
		mSums, pSums := make([]int, len(rests)), make([]int, len(rests))
		mSums[0] = rests[0]
		pSums[0] = d - rests[0]
		for i := 1; i < len(rests); i++ {
			mSums[i] = mSums[i-1] + rests[i]
			pSums[i] = pSums[i-1] + d - rests[i]
		}
		for i := 0; i < len(rests); i++ {
			if mSums[i] == pSums[len(rests)-1]-pSums[i] && mSums[i] <= k {
				return d
			}
		}
	}
	return 1
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
