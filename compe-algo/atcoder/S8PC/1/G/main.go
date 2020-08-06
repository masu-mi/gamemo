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
	resolve(parseProblem())
}

func parseProblem() (int, int, *weightedLinkedList) {
	n, m := scanInt(sc), scanInt(sc)
	g := nextWeightedLinkedList(n, m, 1, sc)
	return n, m, g
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/weigthedlinkedlist.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/next.go] with goone.

type weightedLinkedList struct {
	size, deg int

	edges []map[int]pair
}

type pair struct{ w, l int }

func newWeightedLinkedList(size int) *weightedLinkedList {
	ll := &weightedLinkedList{size: size, edges: make([]map[int]pair, size)}
	for i := 0; i < size; i++ {
		ll.edges[i] = map[int]pair{}
	}
	return ll
}

func (ll *weightedLinkedList) addEdge(a, b, w, l int) {
	ll.addDirectedEdge(a, b, w, l)
	ll.addDirectedEdge(b, a, w, l)
}

func (ll *weightedLinkedList) addDirectedEdge(a, b, w, l int) {
	if _, ok := ll.edges[a][b]; !ok {
		ll.edges[a][b] = pair{w: w, l: l}
		ll.deg++
	}
}

func (ll *weightedLinkedList) info(a, b int) (pair, bool) {
	p, ok := ll.edges[a][b]
	return p, ok
}

func (ll *weightedLinkedList) weight(a, b int) (int, bool) {
	p, ok := ll.info(a, b)
	return p.w, ok
}

func (ll *weightedLinkedList) limit(a, b int) (int, bool) {
	p, ok := ll.info(a, b)
	return p.l, ok
}

func nextWeightedLinkedList(n, m, offset int, sc *bufio.Scanner) *weightedLinkedList {
	ll := newWeightedLinkedList(n)
	for i := 0; i < m; i++ {
		x, y, w, l := nextInt(sc), nextInt(sc), nextInt(sc), nextInt(sc)
		x -= offset
		y -= offset
		ll.addEdge(x, y, w, l)
	}
	return ll
}

func nextDirectedWeightedLinkedList(n, m, offset int, sc *bufio.Scanner) *weightedLinkedList {
	ll := newWeightedLinkedList(n)
	for i := 0; i < m; i++ {
		x, y, w, l := nextInt(sc), nextInt(sc), nextInt(sc), nextInt(sc)
		x -= offset
		y -= offset
		ll.addDirectedEdge(x, y, w, l)
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

func resolve(n, m int, g *weightedLinkedList) {
	itemNum := uint(n)
	numOfSubset := 1 << itemNum
	// init
	dp := make([][]pair, numOfSubset)
	for i := 0; i < numOfSubset; i++ {
		dp[i] = make([]pair, n)
		for j := 0; j < n; j++ {
			dp[i][j] = pair{w: math.MaxInt64, l: 0}
		}
	}
	dp[0][0] = pair{w: 0, l: 0}
	for i := 0; i < n; i++ {
		if info, ok := g.info(0, i); ok {
			dp[1<<uint(i)][i] = pair{w: dp[0][0].w + info.w, l: 1}
		}
	}
	for i := 1; i < numOfSubset; i++ {
		for j := 0; j < n; j++ {
			addingItem := 1 << uint(j)
			if i&addingItem != 0 {
				continue
			}
			for k := range dp[i] {
				if dp[i][k].l == 0 {
					continue
				}
				info, ok := g.info(k, j)
				cand := dp[i][k].w + info.w
				if !ok || cand > info.l {
					continue
				}
				next := i | 1<<uint(j)
				if base := dp[next][j].w; base > cand {
					dp[next][j] = pair{w: cand, l: dp[i][k].l}
				} else if base == cand {
					dp[next][j].l += dp[i][k].l
				}
			}
		}
	}
	if dp[numOfSubset-1][0].l == 0 {
		fmt.Println("IMPOSSIBLE")
		return
	}
	fmt.Println(dp[numOfSubset-1][0].w, dp[numOfSubset-1][0].l)
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
