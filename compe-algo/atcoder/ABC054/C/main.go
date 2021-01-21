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
	fmt.Println(resolve())
}

func resolve() int {
	n, m := scanInt(sc), scanInt(sc)
	ps := newSet()
	for i := 0; i < m; i++ {
		x, y := scanInt(sc)-1, scanInt(sc)-1
		ps.add(pair{x, y})
		ps.add(pair{y, x})
	}

	count := 0
	for p := range permutations(n-1, 1) {
		if allExists(ps, p) {
			count++
		}
	}
	return count
}

func allExists(ps pairSet, p []int) bool {
	cur := 0
	for i := 0; i < len(p); i++ {
		next := p[i]
		if !ps.doesContain(pair{cur, next}) {
			return false
		}
		cur = next
	}
	return true
}

type pair struct{ x, y int }

type pairSet map[pair]none

func newSet() pairSet {
	return make(map[pair]none)
}

func (s pairSet) add(item pair) {
	s[item] = mark
}

func (s pairSet) doesContain(item pair) bool {
	_, ok := s[item]
	return ok
}

func (s pairSet) size() int {
	return len(s)
}

func (s pairSet) members() (l []pair) {
	for k := range s {
		l = append(l, k)
	}
	return l
}

var mark none

type none struct{}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/exhaustive_permutations.go] with goone.

func permutations(l, offset int) chan []int {
	ch := make(chan []int)
	go func() {
		dfsPermutations(0, offset, make([]bool, l), []int{}, func(perm []int) bool {
			ch <- perm
			return false
		})
		close(ch)
	}()
	return ch
}

func dfsPermutations(pos, off int, used []bool, perm []int, atLeaf func(perm []int) (halt bool)) (halt bool) {
	l := len(used)
	if pos == l {
		p := append(perm[:0:0], perm...)
		return atLeaf(p)
	}

	for i := 0; i < l; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		if dfsPermutations(pos+1, off, used, append(perm, i+off), atLeaf) {
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
