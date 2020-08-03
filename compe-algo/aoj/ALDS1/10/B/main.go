package main

import (
	"bufio"
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

type matrixSize struct{ row, col int }

func calcNum(i, j matrixSize) int {
	return i.row * j.col * i.col
}

func parseProblem() (int, []matrixSize) {
	n := scanInt(sc)
	sizes := make([]matrixSize, n)
	for i := 0; i < n; i++ {
		sizes[i] = matrixSize{row: scanInt(sc), col: scanInt(sc)}
	}
	return n, sizes
}

var dp [][]int

func initDP(n int) {
	dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt64
		}
	}
}

func resolve(n int, sizes []matrixSize) int {
	// TODO 復習
	initDP(n)
	_rec(sizes, 0, n)
	return dp[0][n]
}

func _rec(sizes []matrixSize, l, r int) int {
	if dp[l][r] != math.MaxInt64 {
		return dp[l][r]
	}
	if r-l <= 1 {
		dp[l][r] = 0
		return 0
	}
	num := math.MaxInt64
	for i := l + 1; i < r; i++ {
		v := _rec(sizes, l, i) + _rec(sizes, i, r) + calcNum(matrixSize{sizes[l].row, sizes[i].row}, matrixSize{sizes[i].row, sizes[r-1].col})
		if v < num {
			num = v
		}
	}
	dp[l][r] = num
	return dp[l][r]
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
