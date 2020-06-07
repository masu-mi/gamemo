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
	resolve(parseProblem())
}

func parseProblem() (int, int) {
	return scanInt(sc), scanInt(sc)
}

func resolve(n, q int) {
	_, z := nextZetaIntSlice(sc, n, 0, nil)
	for i := 0; i < q; i++ {
		l, r := scanInt(sc), scanInt(sc)
		fmt.Println(z[r] - z[l])
	}
	return
}

func nextIntSlice(sc *bufio.Scanner, n int) (a []int) {
	// 0-indexed
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
	}
	return a
}

func nextZetaIntSlice(sc *bufio.Scanner, n int, init int, op func(i, j int) int) (a, z []int) {
	a = make([]int, n)
	z = make([]int, n+1)
	r := &reducer{init: init, op: op}
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
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

func zetaTransform(a []int) (z []int) {
	z = make([]int, len(a)+1)
	(&reducer{}).reduce(&z, a)
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
