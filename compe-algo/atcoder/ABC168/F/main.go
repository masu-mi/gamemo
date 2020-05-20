package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) (vs []tuple, hs []tuple) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m := scanInt(sc), scanInt(sc)
	vs = make([]tuple, n)
	for i := 0; i < n; i++ {
		a, b, c := scanInt(sc), scanInt(sc), scanInt(sc)
		vs[i] = tuple{a, b, c}
	}
	hs = make([]tuple, m)
	for i := 0; i < m; i++ {
		// d, e, f
		d, e, f := scanInt(sc), scanInt(sc), scanInt(sc)
		hs[i] = tuple{e, f, d}
	}
	return vs, hs
}

type compressedIDs struct {
	idx map[int]int
	l   []int
}

func newCompressedIDs() *compressedIDs {
	return &compressedIDs{idx: make(map[int]int)}
}

func (ci *compressedIDs) addItem(item int) {
	ci.l = append(ci.l, item)
}

func (ci *compressedIDs) build() {
	sort.Sort(sort.IntSlice(ci.l))
	uniq(&ci.l)
	for id, name := range ci.l {
		ci.idx[name] = id
	}
}
func uniq(l *[]int) {
	if len(*l) == 0 {
		return
	}
	li := *l
	r := make([]int, 0, len(li))
	pre := li[0]
	r = append(r, li[0])
	for i := 1; i < len(li); i++ {
		if pre == li[i] {
			continue
		}
		r = append(r, li[i])
		pre = li[i]
	}
	*l = r
}

func (ci *compressedIDs) name(id int) int {
	return ci.l[id]
}
func (ci *compressedIDs) id(name int) int {
	return ci.idx[name]
}

type tuple struct {
	s, e, counterPos int
}

func resolve(vs, hs []tuple) {
	// compressID
	vID, hID := newCompressedIDs(), newCompressedIDs()
	vID.addItem(0)
	hID.addItem(0)
	for _, v := range vs {
		vID.addItem(v.s)
		vID.addItem(v.e)
		hID.addItem(v.counterPos)
	}
	for _, h := range hs {
		hID.addItem(h.s)
		hID.addItem(h.e)
		vID.addItem(h.counterPos)
	}
	vID.build()
	hID.build()

	// build grid
	gridVSize := (len(vID.l))*2 + 1
	gridHSize := (len(hID.l))*2 + 1
	grid := make([][]int, gridVSize)
	visited := make([][]bool, gridVSize)
	for i := 0; i < gridVSize; i++ {
		grid[i] = make([]int, gridHSize)
		visited[i] = make([]bool, gridHSize)
	}
	for _, v := range vs {
		hPos := hID.id(v.counterPos)
		for i := vID.id(v.s)*2 + 1; i <= vID.id(v.e)*2+1; i++ {
			grid[i][hPos*2+1] = 1
		}
	}
	for _, h := range hs {
		vPos := vID.id(h.counterPos)
		for i := hID.id(h.s)*2 + 1; i <= hID.id(h.e)*2+1; i++ {
			grid[vPos*2+1][i] = 1
		}
	}
	// for i := range grid {
	// 	fmt.Printf("%v\n", grid[i])
	// }
	sum := 0
	visited[vID.id(0)*2+1][hID.id(0)*2+1] = true
	q := []pos{pos{vID.id(0)*2 + 1, hID.id(0)*2 + 1}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.v == 0 || cur.v == gridVSize-1 {
			fmt.Println("INF")
			return
		} else if cur.h == 0 || cur.h == gridHSize-1 {
			fmt.Println("INF")
			return
		}
		if cur.v%2 == 0 && cur.h%2 == 0 {
			vId := (cur.v - 1) / 2
			hId := (cur.h - 1) / 2
			// TODO RE here out of range
			vLen := vID.name(vId+1) - vID.name(vId)
			hLen := hID.name(hId+1) - hID.name(hId)
			sum += vLen * hLen
		}
		for _, d := range []pos{
			pos{1, 0},
			pos{-1, 0},
			pos{0, -1},
			pos{0, 1},
		} {
			next := pos{cur.v + d.v, cur.h + d.h}
			if visited[next.v][next.h] {
				continue
			}
			if grid[cur.v+d.v][cur.h+d.h] == 1 {
				continue
			}
			visited[next.v][next.h] = true
			q = append(q, next)
		}
	}
	fmt.Println(sum)
	return
}

type pos struct{ v, h int }

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
