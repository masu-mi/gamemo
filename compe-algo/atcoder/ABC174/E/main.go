package main

import (
	"bufio"
	"container/heap"
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

type pair struct {
	size float64
	num  int
}

func (p pair) segSize() float64 {
	return p.size / float64(p.num)
}

func parseProblem() int {
	n, k := scanInt(sc), scanInt(sc)
	ps := make([]pair, n)
	for i := 0; i < n; i++ {
		ps[i] = pair{
			size: float64(scanInt(sc)),
			num:  1,
		}
	}
	ops := 0
	if len(ps) == 1 {
		ps[0].num += k
		return int(math.Ceil(ps[0].segSize()))
	}
	segs := segments(ps)
	heap.Init(&segs)
	for ops < k {
		h := heap.Pop(&segs).(pair)
		s := heap.Pop(&segs).(pair)

		requiredSegNum := int(math.Ceil(h.size / s.segSize()))
		reqOps := requiredSegNum - h.num
		curOps := min(k-ops, reqOps)
		h.num += curOps
		ops += curOps
		heap.Push(&segs, h)
		heap.Push(&segs, s)
	}
	h := heap.Pop(&segs).(pair)
	return int(math.Ceil(h.segSize()))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type segments []pair

func (h *segments) Len() int {
	return len(*h)
}

func (h *segments) Less(i, j int) bool {
	items := *h
	return items[i].segSize() > items[j].segSize()
}

func (h *segments) Swap(i, j int) {
	items := *h
	items[i], items[j] = items[j], items[i]
}

func (h *segments) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *segments) Pop() interface{} {
	items := *h
	l := items[len(items)-1]
	*h = items[0 : len(items)-1]
	return l
}

func (h *segments) top() pair {
	return (*h)[h.Len()-1]
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
