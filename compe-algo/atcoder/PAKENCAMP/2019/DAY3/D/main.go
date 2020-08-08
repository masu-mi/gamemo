package main

import (
	"bufio"
	"bytes"
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

func parseProblem() (int, []string) {
	_, w, g := loadGrid(sc)
	return w, g
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/next.go /Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/grid.go] with goone.

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

const (
	wallByte = '#'
)

func loadGrid(sc *bufio.Scanner) (h, w int, grid []string) {
	h, w = 5, nextInt(sc)
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

func newBooleanGrid(h, w int) (r [][]bool) {
	r = make([][]bool, h+2)
	for i := range r {
		r[i] = make([]bool, w+2)
	}
	return r
}

func getByte(g []string, p pos) byte {
	return g[p.x][p.y]
}

type pos struct{ x, y int }

func findBytes(g []string, str string) map[byte]pos {
	targets := map[byte]struct{}{}
	result := map[byte]pos{}
	for i := range str {
		targets[str[i]] = struct{}{}
	}
	for i := 1; i < len(g); i++ {
		for j := 1; j < len(g[i]); j++ {
			b := g[i][j]
			if _, ok := targets[b]; ok {
				result[b] = pos{i, j}
			}
		}
	}
	return result
}

func resolve(w int, g []string) int {
	colors := []byte{'R', 'B', 'W'}
	dp := make([][]int, w+1)
	for i := range loop0(w + 1) {
		dp[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	for i := 0; i < 3; i++ {
		dp[1][i] = countRequired(g, colors[i], 1)
	}
	for i := 2; i <= w; i++ {
		for color := 0; color < 3; color++ {
			for pc, score := range dp[i-1] {
				if pc == color {
					continue
				}
				v := score + countRequired(g, colors[color], i)
				if v <= dp[i][color] {
					dp[i][color] = v
				}
			}
		}
	}
	min := math.MaxInt64
	for _, score := range dp[w] {
		if score < min {
			min = score
		}
	}
	return min
}

func countRequired(g []string, c byte, idx int) int {
	num := 0
	for j := 1; j <= 5; j++ {
		if g[j][idx] != c {
			num++
		}
	}
	return num
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/simple_loop.go] with goone.

func loop0(n int) chan int {
	return loop(0, n-1, 1)
}

func loop1(n int) chan int {
	return loop(1, n, 1)
}

func loop(s, e, d int) chan int {
	ch := make(chan int)
	go func() {
		for i := s; i <= e; i += d {
			ch <- i
		}
		close(ch)
	}()
	return ch
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
