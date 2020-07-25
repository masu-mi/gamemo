package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

type v struct{ x, y, p int }

type w struct{ tipe, pos int }

type ways []w

func (w ways) Len() int {
	return len(w)
}

func (w ways) Less(i, j int) bool {
	if w[i].tipe == w[j].tipe {
		return w[i].pos < w[j].pos
	}
	return w[i].tipe < w[j].tipe
}

func (w ways) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func parseProblem() int {
	n := scanInt(sc)
	vs := make([]v, n)
	for i := 0; i < n; i++ {
		vs[i] = v{
			x: scanInt(sc),
			y: scanInt(sc),
			p: scanInt(sc),
		}
	}
	var ws []w
	for i := range vs {
		ws = append(ws, w{tipe: 0, pos: vs[i].x})
		ws = append(ws, w{tipe: 1, pos: vs[i].y})
	}
	var candidate []w
	sort.Sort(ways(ws))
	candidate = append(candidate, ws[0])
	for i := 1; i < len(ws); i++ {
		if ws[i] != ws[i-1] {
			candidate = append(candidate, ws[i])
		}
	}
	for i := 0; i <= n; i++ {
		s := minimumScore(vs, candidate, i)
		fmt.Println(s)
	}
	return scanInt(sc)
}

func minimumScore(vs []v, candidate []w, k int) int {
	min := math.MaxInt32
	for mask := range bitCombinationsWithSize(len(vs), k) {
		v := score(vs, candidate, mask)
		if v < min {
			min = v
		}
	}
	return min
}

func score(vs []v, candidate []w, mask uint) int {
	sum := 0
	infra := newInfra(candidate, len(vs), mask)
	for i := range vs {
		sum += _score(vs[i], infra)
	}
	return sum
}

func _score(vi v, infra constructedInfra) int {
}

type constructedInfra struct{ x, y []int }

func newInfra(candidate []w, n int, mask uint) constructedInfra {
	result := constructedInfra{}
	for i := 0; i < len(candidate); i++ {
		if 1<<uint(i)&mask != 0 {
			used := candidate[i]
			switch used.tipe {
			case 0:
				result.x = append(result.x, used.pos)
			case 1:
				result.y = append(result.y, used.pos)
			}
		}
	}
	sort.Sort(sort.IntSlice(result.x))
	sort.Sort(sort.IntSlice(result.y))
	return result
}

func bitCombinations(num int) chan uint {
	ch := make(chan uint)
	go func() {
		defer close(ch)
		for i := 0; i < 1<<uint(num); i++ {
			ch <- uint(i)
		}
	}()
	return ch
}

func bitCombinationsOverSubsets(nums ...int) chan uint {
	ch := make(chan uint)
	s := uint(0)
	for _, v := range nums {
		s |= 1 << uint(v)
	}
	go func() {
		defer close(ch)
		for bit := s; ; bit = (bit - 1) & s {
			ch <- uint(bit)
			if bit == 0 {
				break
			}
		}
	}()
	return ch
}

func bitCombinationsWithSize(num, size int) chan uint {
	ch := make(chan uint)
	bit := uint(1<<uint(size) - 1)
	go func() {
		defer close(ch)
		for ; bit < 1<<uint(num); bit = nextBitCombination(uint(bit)) {
			ch <- bit
		}
	}()
	return ch
}

func nextBitCombination(cur uint) uint {
	x := cur & -cur
	y := cur + x
	return (((cur & ^y) / x) >> 1) | y
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
