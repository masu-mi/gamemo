package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) ([][]int, []int, int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, m, c := scanInt(sc), scanInt(sc), scanInt(sc)
	b := nextMatrix(1, m, sc)
	mx := nextMatrix(n, m, sc)
	return mx, b[0], c
}

func nextMatrix(n, m int, sc *bufio.Scanner) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = nextInt(sc)
		}
	}
	return
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

func resolve(mx [][]int, b []int, c int) (num int) {
	for i := 0; i < len(mx); i++ {
		if mul(i, mx, b)+c > 0 {
			num++
		}
	}
	return num
}

func mul(i int, mx [][]int, v []int) int {
	sum := 0
	for j := 0; j < len(mx[i]); j++ {
		sum += mx[i][j] * v[j]
	}
	return sum
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
