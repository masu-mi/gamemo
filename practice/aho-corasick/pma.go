package main

import (
	"sort"
)

type pma struct {
	edges        []map[byte]int
	failureEdges []int
	patterns     map[int][]string
}

func newPMA(patterns []string) *pma {
	a := &pma{patterns: map[int][]string{}}
	a.addNode()

	sort.Sort(sort.StringSlice(patterns))
	for _, p := range patterns {
		a.addPatternEdgeds(p)
	}
	a.buildFailureEdges()
	return a
}

func (a *pma) addPatternEdgeds(pattern string) {
	cur := a.rootNode()
	for i := 0; i < len(pattern); i++ {
		a.addNode()
		next := a.lastNode()
		a.edges[cur][pattern[i]] = next
		cur = next
	}
	a.patterns[cur] = append(a.patterns[cur], pattern)
}

func (a *pma) buildFailureEdges() {
	// implement with BFS
	a.failureEdges = make([]int, a.nodeNum())
	a.failureEdges[a.rootNode()] = a.rootNode()
	q := []int{a.rootNode()}
	for len(q) > 0 {
		var cur int
		cur, q = q[0], q[1:]
		for token, next := range a.edges[cur] {
			a.failureEdges[next] = a.findFailure(cur, next, token)
			q = append(q, next)
		}
	}
}

func (a *pma) findFailure(cur, next int, token byte) int {
	root := a.rootNode()
	for true {
		if cur == root {
			return root
		}
		curF := a.failureEdges[cur]
		if candidate, ok := a.edges[curF][token]; ok {
			a.failureEdges[next] = candidate
			for _, p := range a.patterns[candidate] {
				a.patterns[next] = append(a.patterns[next], p)
			}
			return candidate
		}
		cur = curF
	}
	return root
}

func (a *pma) addNode()      { a.edges = append(a.edges, map[byte]int{}) }
func (a *pma) nodeNum() int  { return len(a.edges) }
func (a *pma) rootNode() int { return 0 }
func (a *pma) lastNode() int { return a.nodeNum() - 1 }

func (a *pma) nextNode(node int, token byte) int {
	if nextNode, ok := a.edges[node][token]; ok {
		return nextNode
	}
	for true {
		node = a.failureEdges[node]
		if nextNode, ok := a.edges[node][token]; ok {
			return nextNode
		}
		isRootNode := node == a.failureEdges[node]
		if isRootNode {
			break
		}
	}
	return node
}

func (a *pma) searchPatterns(input string) chan result {
	ch := make(chan result)
	go func() {
		iter := newIterator(a, input)
		for iter.active() {
			iter.next()
			for _, p := range iter.foundPatterns() {
				ch <- p
			}
		}
		close(ch)
	}()
	return ch
}

type result struct {
	start   int
	pattern string
}

type iterator struct {
	*pma
	state int

	input string
	index int
}

func newIterator(a *pma, input string) *iterator {
	return &iterator{pma: a, state: a.rootNode(), input: input, index: 0}
}

func (iter *iterator) active() bool {
	return iter.index < len(iter.input)
}

func (iter *iterator) next() (halt bool) {
	iter.state = iter.pma.nextNode(iter.state, iter.input[iter.index])
	iter.index++
	return !iter.active()
}

func (iter *iterator) foundPatterns() []result {
	var r []result
	for _, p := range iter.pma.patterns[iter.state] {
		r = append(r, result{
			start:   iter.index - len(p),
			pattern: p,
		})
	}
	return r
}
