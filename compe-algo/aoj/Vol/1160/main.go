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
	resolve()
}

func resolve() {
	results := []int{}
	for true {
		w, h := scanInt(sc), scanInt(sc)
		if w == 0 && h == 0 {
			break
		}
		grid := make([][]int, h)
		for i := 0; i < h; i++ {
			grid[i] = make([]int, w)
			for j := 0; j < w; j++ {
				grid[i][j] = scanInt(sc)
			}
		}
		num := 0
		for i := range loop0(h) {
			for j := range loop0(w) {
				start := pos{y: i, x: j}
				if grid[start.y][start.x] == 0 {
					continue
				}
				num++
				_dfs(grid, w, h, start)
			}
		}
		results = append(results, num)
	}
	for _, v := range results {
		fmt.Println(v)
	}
	return
}

func _dfs(grid [][]int, w, h int, start pos) {
	grid[start.y][start.x] = 0
	for _, df := range []pos{
		pos{1, 0}, pos{-1, 0},
		pos{0, -1}, pos{0, 1},
		pos{1, 1}, pos{-1, -1},
		pos{-1, 1}, pos{1, -1},
	} {
		np := addPos(start, df)
		if np.y < 0 || h <= np.y || np.x < 0 || w <= np.x {
			continue
		}
		if grid[np.y][np.x] == 0 {
			continue
		}
		_dfs(grid, w, h, np)
	}
}

type pos struct{ x, y int }

func addPos(p, d pos) pos {
	return pos{
		x: p.x + d.x,
		y: p.y + d.y,
	}
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
