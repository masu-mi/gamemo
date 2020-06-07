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

func parseProblem() (n, q int) {
	return scanInt(sc), scanInt(sc)
}

func resolve(n, q int) {
	uf := newUnifonFind(n)
	for i := 0; i < q; i++ {
		switch scanInt(sc) {
		case 0:
			uf.union(scanInt(sc), scanInt(sc))
		case 1:
			if uf.same(scanInt(sc), scanInt(sc)) {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	}
	return
}

type unionfind struct {
	card      int
	connected bool

	parent []int
	rank   []int
	childs []int
}

func newUnifonFind(card int) *unionfind {
	uf := &unionfind{
		card:      card,
		parent:    make([]int, card),
		rank:      make([]int, card),
		childs:    make([]int, card),
		connected: card == 1,
	}
	for i := 0; i < card; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (u *unionfind) find(x int) int {
	p := u.parent[x]
	if p == x {
		return x
	}
	r := u.find(p)
	u.parent[x] = r
	return r
}

func (u *unionfind) same(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u *unionfind) union(x, y int) {
	xR, yR := u.find(x), u.find(y)
	if xR == yR {
		return
	}
	if rankX, rankY := u.rank[xR], u.rank[yR]; rankX < rankY {
		u.parent[xR] = yR
		u.childs[yR] += u.childs[xR] + 1
		u.connected = u.card == u.childs[yR]+1
	} else {
		u.parent[yR] = xR
		u.childs[xR] += u.childs[yR] + 1
		u.connected = u.card == u.childs[xR]+1
		if rankX == rankY {
			u.rank[xR]++
		}
	}
}

func (u *unionfind) size(x int) int {
	return u.childs[u.find(x)] + 1
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
