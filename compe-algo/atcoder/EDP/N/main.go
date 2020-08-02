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

func parseProblem() (int, []int) {
	n := scanInt(sc)
	as := nextIntSlice(sc, n)
	return n, as
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

type info struct{ cost, size int }

var dp [][]info

func resolve(n int, as []int) int {
	initDP(n)
	return _rec(as, 0, len(as)).cost
}

func _rec(as []int, l, r int) info {
	if dp[l][r].cost != 0 { // no 0-size slime
		return dp[l][r]
	}
	if l == r {
		return info{}
	}
	if l+1 == r {
		dp[l][r] = info{cost: 0, size: as[l]}
		return dp[l][r]
	}
	dp[l][r] = info{cost: math.MaxInt64, size: -1}
	for i := l + 1; i < r; i++ {
		cand := fusion(_rec(as, l, i), _rec(as, i, r))
		if cand.cost < dp[l][r].cost {
			dp[l][r] = cand
		}
	}
	return dp[l][r]
}

func fusion(x, y info) info {
	return info{
		cost: x.size + y.size + x.cost + y.cost,
		size: x.size + y.size,
	}
}

func initDP(n int) {
	dp = make([][]info, n+1)
	for i := range loop0(n + 1) {
		dp[i] = make([]info, n+1)
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
