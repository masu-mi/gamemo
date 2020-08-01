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

func parseProblem() int {
	w, h := scanInt(sc), scanInt(sc)
	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)
		for j := 0; j < w; j++ {
			grid[i][j] = scanInt(sc)
		}
	}
	visited := make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	maxLen := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 0 {
				continue
			}
			ml := _dfs(visited, grid, i, j, 1)
			if maxLen < ml {
				maxLen = ml
			}
		}
	}
	return maxLen
}

type pair struct{ x, y int }

func _dfs(visited [][]bool, grid [][]int, i, j, l int) int {
	visited[i][j] = true
	defer func() { visited[i][j] = false }()
	maxL := l
	for _, df := range []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		ni, nj := i+df.y, j+df.x
		if ni < 0 || len(grid) <= ni || nj < 0 || len(grid[0]) <= nj {
			continue
		}
		if visited[ni][nj] {
			continue
		}
		if grid[ni][nj] == 0 {
			continue
		}
		ll := _dfs(visited, grid, ni, nj, l+1)
		if maxL < ll {
			maxL = ll
		}
	}
	return maxL
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
