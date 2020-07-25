package main

import (
	"bufio"
	"bytes"
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
	if resolve(parseProblem()) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func parseProblem() (int, int, []string) {
	h, w, g := loadGrid(sc)
	return h, w, g
}

const (
	wallByte = '#'
)

func loadGrid(sc *bufio.Scanner) (h, w int, grid []string) {
	h, w = nextInt(sc), nextInt(sc)
	grid = make([]string, h+2)
	wall := createWall(w)
	grid[0] = wall
	for i := 1; i <= h; i++ {
		sc.Scan()
		buf := bytes.NewBuffer([]byte{})
		buf.Write([]byte{wallByte})
		buf.WriteString(sc.Text())
		buf.Write([]byte{wallByte})
		grid[i] = buf.String()
	}
	grid[h+1] = wall
	return h, w, grid
}

func createWall(w int) string {
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < w+2; i++ {
		buf.Write([]byte{wallByte})
	}
	return buf.String()
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

type pos struct{ x, y int }

func resolve(h, w int, g []string) bool {
	s := find(g, 's')
	t := find(g, 'g')
	visited := newBooleanGrid(h, w)
	visited[s.x][s.y] = true
	q := []pos{s}
	var cur pos
	for len(q) > 0 {
		cur, q = q[0], q[1:]
		for _, df := range []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			next := pos{cur.x + df.x, cur.y + df.y}
			if visited[next.x][next.y] {
				continue
			}
			visited[next.x][next.y] = true
			if g[next.x][next.y] == '#' {
				continue
			}
			if next == t {
				return true
			}
			q = append(q, next)
		}
	}
	return false
}

func newBooleanGrid(h, w int) (r [][]bool) {
	r = make([][]bool, h+2)
	for i := range r {
		r[i] = make([]bool, w+2)
	}
	return r
}

func find(g []string, c byte) pos {
	for i := 1; i < len(g); i++ {
		for j := 1; j < len(g[i]); j++ {
			if g[i][j] == c {
				return pos{i, j}
			}
		}
	}
	return pos{0, 0}
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
