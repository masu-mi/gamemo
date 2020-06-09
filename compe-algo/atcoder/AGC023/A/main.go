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
	return scanInt(sc)
}

func resolve(n int) int {
	_, z := nextComulative(sc, n, 0, nil)
	sums := map[int]int{}
	num := 0
	for i := 1; i < n+1; i++ {
		sums[z[i]]++
		num += sums[z[i]] - 1
	}
	// ranges include only 0 are acceptable
	num += sums[0]
	return num
}

func getComulative(a []int) (z []int) {
	z = make([]int, len(a)+1)
	(&reducer{}).reduce(&z, a)
	return
}

func nextComulative(sc *bufio.Scanner, n int, init int, op func(v, acc int) int) (a, z []int) {
	a = make([]int, n)
	z = make([]int, n+1)
	r := &reducer{init: init, op: op}
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
		r.update(&z, a, i)
	}
	return a, z
}

type reducer struct {
	init, tmp int
	op        func(v, acc int) int
}

func (r *reducer) update(z *[]int, a []int, idx int) {
	if r.op != nil {
		r.tmp = r.op(r.tmp, a[idx])
	} else {
		r.tmp += a[idx]
	}
	(*z)[idx+1] = r.tmp
}

func (r *reducer) reduce(z *[]int, a []int) {
	for i := 0; i < len(a); i++ {
		r.update(z, a, i)
	}
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
