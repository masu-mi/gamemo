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

type state struct{ pasta, num int }

func parseProblem() int {
	n, k := scanInt(sc), scanInt(sc)
	constraint := map[int]int{}
	for _ = range loop0(k) {
		constraint[scanInt(sc)] = scanInt(sc)
	}
	dp := make([]map[state]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = map[state]int{}
	}
	for _, s := range []state{state{1, 1}, state{2, 1}, state{3, 1}} {
		if p, ok := constraint[1]; ok && s.pasta != p {
			continue
		}
		dp[1][s] = 1
	}
	for i := 2; i <= n; i++ {
		for _, s := range []state{state{1, 1}, state{2, 1}, state{3, 1}} {
			if p, ok := constraint[i]; ok && s.pasta != p {
				continue
			}
			for k, v := range dp[i-1] {
				next := s
				if k.pasta == next.pasta {
					next.num += k.num
				}
				if next.num >= 3 {
					continue
				}
				dp[i][next] = (dp[i][next] + v) % 10000
			}
		}
	}
	sum := 0
	for _, v := range dp[n] {
		sum += v
	}
	return sum % 10000
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
