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
	n, m := scanInt(sc), scanInt(sc)
	g := nextLinkedList(n, m, 1, sc)
	max := 0
	for i := 1; i < 1<<uint(n); i++ {
		ok := true
	check:
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) == 0 {
				continue
			}
			for k := j + 1; k < n; k++ {
				if i&(1<<uint(k)) == 0 {
					continue
				}
				if !g.exists(j, k) {
					ok = false
					break check
				}
			}
		}
		if ok {
			if v := onesCount(uint64(i)); max < v {
				max = v
			}
		}
	}
	return max
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/bits_util.go] with goone.

func onesCount(x uint64) (num int) {
	const m0 = 0x5555555555555555
	const m1 = 0x3333333333333333
	const m2 = 0x0f0f0f0f0f0f0f0f

	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}

func ntz(bits uint64) (num int) {

	return onesCount(bits&(-bits) - 1)
}

func nlz(bits uint64) (num int) {

	bits = bits | (bits >> 1)
	bits = bits | (bits >> 2)
	bits = bits | (bits >> 4)
	bits = bits | (bits >> 8)
	bits = bits | (bits >> 16)
	bits = bits | (bits >> 32)
	return onesCount(^bits)
}

func refBit(i uint64, b uint) int {
	return int((i >> b) & 1)
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/linkedlist.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/set.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/none.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/next.go] with goone.

type basicLinkedList struct {
	size, deg int
	edges     []intSet
}

func newLinkedList(size int) *basicLinkedList {
	ll := &basicLinkedList{size: size, edges: make([]intSet, size)}
	for i := 0; i < size; i++ {
		ll.edges[i] = newIntSet()
	}
	return ll
}

func (ll *basicLinkedList) addEdge(a, b int) {
	ll.addDirectedEdge(a, b)
	ll.addDirectedEdge(b, a)
}

func (ll *basicLinkedList) addDirectedEdge(a, b int) {
	if ll.edges[a].add(b) {
		ll.deg++
	}
}

func (ll *basicLinkedList) exists(a, b int) bool {
	return ll.edges[a].doesContain(b)
}

func nextLinkedList(n, m, offset int, sc *bufio.Scanner) *basicLinkedList {
	ll := newLinkedList(n)
	for i := 0; i < m; i++ {
		x, y := nextInt(sc), nextInt(sc)

		x -= offset
		y -= offset
		ll.addEdge(x, y)
	}
	return ll
}

func nextDirectedLinkedList(n, m, offset int, sc *bufio.Scanner) *basicLinkedList {
	ll := newLinkedList(n)
	for i := 0; i < m; i++ {
		x, y := nextInt(sc), nextInt(sc)

		x -= offset
		y -= offset
		ll.addDirectedEdge(x, y)
	}
	return ll
}

type intSet map[int]none

func newIntSet() intSet {
	return map[int]none{}
}

func (s intSet) add(i int) (added bool) {
	_, ok := s[i]
	added = !ok
	s[i] = mark
	return
}

func (s intSet) remove(i int) (removed bool) {
	_, removed = s[i]
	delete(s, i)
	return
}

func (s intSet) doesContain(i int) bool {
	_, ok := s[i]
	return ok
}

func (s intSet) size() int {
	return len(s)
}

func (s intSet) members() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for k := range s {
			ch <- k
		}
	}()
	return ch
}

type none struct{}

var mark none

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

func resolve(n int) int {
	return n
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
