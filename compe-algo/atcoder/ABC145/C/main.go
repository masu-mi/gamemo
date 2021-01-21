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

func resolve() float64 {
	n := scanInt(sc)
	m := make([]pair, 0, n)
	for i := 0; i < n; i++ {
		m = append(m, pair{x: scanInt(sc), y: scanInt(sc)})
	}
	num := 0
	totalLen := 0.0
	for path := range permutations(n) {
		totalLen += length(m, path)
		num++
	}
	return totalLen / float64(num)
}

func length(m []pair, path []int) float64 {
	l := 0.0
	for i := 1; i < len(path); i++ {
		sp, tp := m[path[i-1]], m[path[i]]
		xd, yd := sp.x-tp.x, sp.y-tp.y
		l += math.Sqrt(float64(xd*xd + yd*yd))
	}
	return l
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/pair.go] with goone.

type pair struct{ x, y int }

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/exhaustive_permutations.go] with goone.

func permutations(l int) chan []int {
	ch := make(chan []int)
	go func() {
		dfsPermutations(0, make([]bool, l), []int{}, func(perm []int) bool {
			p := make([]int, len(perm))
			copy(p, perm)
			ch <- p
			return false
		})
		close(ch)
	}()
	return ch
}

func dfsPermutations(pos int, used []bool, perm []int, atLeaf func(perm []int) (halt bool)) (halt bool) {
	l := len(used)
	if pos == l {
		if atLeaf(perm) {
			return true
		}
	}

	for i := 0; i < l; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		if dfsPermutations(pos+1, used, append(perm, i), atLeaf) {
			return true
		}
		used[i] = false
	}
	return false
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
