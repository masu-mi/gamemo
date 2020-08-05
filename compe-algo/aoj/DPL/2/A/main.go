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
	fmt.Println(resolve(parseProblem()))
}

func parseProblem() (int, int, *weightedLinkedList) {
	n, m := scanInt(sc), scanInt(sc)
	g := nextDirectedWeightedLinkedList(n, m, 0, sc)
	return n, m, g
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/weigthedlinkedlist.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/next.go] with goone.

type weightedLinkedList struct {
	size, deg int

	edges []map[int]int
}

func newWeightedLinkedList(size int) *weightedLinkedList {
	ll := &weightedLinkedList{size: size, edges: make([]map[int]int, size)}
	for i := 0; i < size; i++ {
		ll.edges[i] = map[int]int{}
	}
	return ll
}

func (ll *weightedLinkedList) addEdge(a, b, w int) {
	ll.addDirectedEdge(a, b, w)
	ll.addDirectedEdge(b, a, w)
}

func (ll *weightedLinkedList) addDirectedEdge(a, b, w int) {
	if _, ok := ll.edges[a][b]; !ok {
		ll.edges[a][b] = w
		ll.deg++
	}
}

func (ll *weightedLinkedList) weight(a, b int) (int, bool) {
	w, ok := ll.edges[a][b]
	return w, ok
}

func nextWeightedLinkedList(n, m, offset int, sc *bufio.Scanner) *weightedLinkedList {
	ll := newWeightedLinkedList(n)
	for i := 0; i < m; i++ {
		x, y, w := nextInt(sc), nextInt(sc), nextInt(sc)
		x -= offset
		y -= offset
		ll.addEdge(x, y, w)
	}
	return ll
}

func nextDirectedWeightedLinkedList(n, m, offset int, sc *bufio.Scanner) *weightedLinkedList {
	ll := newWeightedLinkedList(n)
	for i := 0; i < m; i++ {
		x, y, w := nextInt(sc), nextInt(sc), nextInt(sc)
		x -= offset
		y -= offset
		ll.addDirectedEdge(x, y, w)
	}
	return ll
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

func resolve(n, m int, g *weightedLinkedList) int {
	dp := make([][]int, 1<<uint(n))
	for i := 0; i < 1<<uint(n); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		if c, ok := g.edges[0][i]; ok {
			dp[1<<uint(i)][i] = c
		}
	}
	// give dp
	for i := 1; i < 1<<uint(n); i++ {
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) != 0 {
				// if j in subset(i), not target
				continue
			}
			next := i | 1<<uint(j)
			for k := 0; k < n; k++ {
				if i&(1<<uint(k)) == 0 {
					continue
				}
				if dp[i][k] == math.MaxInt64 {
					continue
				}
				if ec, ok := g.edges[k][j]; ok {
					cost := dp[i][k] + ec
					if dp[i|1<<uint(j)][j] > cost {
						dp[next][j] = cost
					}
				}
			}
		}
	}
	v := dp[1<<uint(n)-1][0]
	if v == math.MaxInt64 {
		return -1
	}
	return v
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
