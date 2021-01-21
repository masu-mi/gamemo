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
	ll := nextLinkedList(n, m, 1, sc)
	dfs(ll, 0, n)
	return num
}

var num int
var visited map[int]bool = map[int]bool{}

func dfs(ll *basicLinkedList, start, n int) bool {
	visited[start] = true
	defer func() {
		delete(visited, start)
	}()
	if len(visited) == n {
		num++
		return false
	}
	for nextNode := range ll.edges[start].members() {
		if _, ok := visited[nextNode]; ok {
			continue
		}
		dfs(ll, nextNode, n)
	}
	return false
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/none.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/next.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/linkedlist.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/set.go] with goone.

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
