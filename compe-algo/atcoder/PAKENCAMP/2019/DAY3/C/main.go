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

func parseProblem() (int, int, [][]int) {
	n, m := scanInt(sc), scanInt(sc)
	tb := make([][]int, n)
	for i := 0; i < n; i++ {
		tb[i] = make([]int, m)
		for j := 0; j < m; j++ {
			tb[i][j] = scanInt(sc)
		}
	}
	return n, m, tb
}

func resolve(n, m int, tb [][]int) int {
	result := 0
	for seq := range cols(m-1, 2) {
		sum := 0
		for i := 0; i < n; i++ {
			sum += max(tb[i][seq[0]], tb[i][seq[1]])
		}
		if result < sum {
			result = sum
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func cols(last, length int) chan []int {
	ch := make(chan []int)
	go func() {
		_cols(last, length, 0, nil, func(ans []int) {
			ch <- ans
		})
		close(ch)
	}()
	return ch
}

func _cols(last, length, start int, ans []int, cb func(ans []int)) {
	if len(ans) == length {
		result := make([]int, len(ans))
		copy(result, ans)
		cb(result)
		return
	}
	if last < start {
		return
	}
	for i := start; i <= last; i++ {
		_cols(last, length, i+1, append(ans, i), cb)
	}
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
