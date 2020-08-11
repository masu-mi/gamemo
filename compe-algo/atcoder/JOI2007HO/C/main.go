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

type pos struct{ x, y int }

func (p pos) inv() pos {
	return pos{x: -p.x, y: -p.y}
}

func (p pos) orth() pos {
	return pos{
		x: -p.y,
		y: p.x,
	}
}

func prod(x, y pos) int {
	return x.x*y.x + x.y*y.y
}

func add(x, y pos) pos {
	return pos{
		x: x.x + y.x,
		y: x.y + y.y,
	}
}

func parseProblem() (int, []pos) {
	n := scanInt(sc)
	ps := make([]pos, n)
	for i := 0; i < n; i++ {
		ps[i] = pos{x: scanInt(sc), y: scanInt(sc)}
	}
	return n, ps
}

func resolve(n int, ps []pos) int {
	exists := map[pos]struct{}{}
	for _, p := range ps {
		exists[p] = struct{}{}
	}
	maxArea := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			p1, p2 := ps[i], ps[j]
			v := add(p2, p1.inv())
			area := prod(v, v)
			if maxArea > area {
				continue
			}
			// v(i,j) => p_3, p_4, p_1(i), p_2(j)
			p3 := add(p1, v.orth())
			if _, ok := exists[p3]; !ok {
				continue
			}
			p4 := add(p3, v)
			if _, ok := exists[p4]; !ok {
				continue
			}
			maxArea = area
		}
	}
	return maxArea
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
