package kruskal

import (
	"errors"
	"sort"
)

func findMSTWithKruskal(card int, edges edgeList) (mst edgeList, cost int, err error) {
	sort.Sort(edges)
	uf := newUnionFind(card)
	for _, e := range edges {
		if uf.connectedNodes(e.x, e.y) {
			continue
		}
		uf.union(e.x, e.y)
		mst = append(mst, e)
		cost += e.cost
		if uf.connected() {
			return mst, cost, nil
		}
	}
	return nil, -1, errors.New("No MST")
}

type edge struct {
	x, y, cost int
}
type edgeList []edge

func (e edgeList) Len() int { return len(e) }

func (e edgeList) Less(i, j int) bool { return e[i].cost < e[j].cost }

func (e edgeList) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type unionFind struct {
	maxCard  int
	root     []int
	childNum []int
}

func newUnionFind(num int) *unionFind {
	f := &unionFind{
		maxCard:  1,
		root:     make([]int, num),
		childNum: make([]int, num),
	}
	for i := 0; i < len(f.root); i++ {
		f.root[i] = i
	}
	return f
}

func (uf *unionFind) connected() bool {
	return uf.maxCard == len(uf.root)
}
func (uf *unionFind) connectedNodes(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *unionFind) find(n int) int {
	p := uf.root[n]
	if p == uf.root[p] {
		return p
	}
	uf.childNum[p] = 0
	root := uf.find(p)
	uf.root[n] = root
	uf.childNum[root] = uf.childNum[n] + 1
	return uf.root[n]
}

func (uf *unionFind) union(x, y int) {
	xr, yr := uf.find(x), uf.find(y)
	if uf.childNum[xr] <= uf.childNum[yr] {
		uf.root[xr] = yr
		uf.childNum[yr] += uf.childNum[xr] + 1
		uf.maxCard = max(uf.maxCard, uf.childNum[yr]+1)
	} else {
		uf.root[yr] = xr
		uf.childNum[xr] += uf.childNum[yr] + 1
		uf.maxCard = max(uf.maxCard, uf.childNum[xr]+1)
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
