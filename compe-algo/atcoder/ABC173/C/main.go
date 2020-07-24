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

func parseProblem() (h, w, k int, g [][]int) {
	h, w, k = scanInt(sc), scanInt(sc), scanInt(sc)
	g = nextGrid(h, w)
	return
}

const (
	white = 0 + iota
	black
	red
)

func nextGrid(h, w int) [][]int {
	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
		str := scanString(sc)
		for j := 0; j < w; j++ {
			switch str[j] {
			case '.':
				grid[i][j] = white
			case '#':
				grid[i][j] = black
			}
		}
	}
	return grid
}

func resolve(h, w, k int, g [][]int) int {
	num := 0
	for i := 0; i < 1<<uint(h); i++ {
		for j := 0; j < 1<<uint(w); j++ {
			if fullfill(g, uint(i), uint(j), k) {
				num++
			}
		}
	}
	return num
}

func fullfill(g [][]int, xFilter, yFilter uint, k int) bool {
	h := len(g)
	if h == 0 {
		return k == 0
	}
	w := len(g[0])
	num := 0
	for i := range loop0(h) {
		if xFilter&(1<<uint(i)) != 0 {
			continue
		}
		for j := range loop0(w) {
			if yFilter&(1<<uint(j)) != 0 {
				continue
			}
			if g[i][j] == black {
				num++
			}
			if num > k {
				return false
			}
		}
	}
	return num == k
}

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
