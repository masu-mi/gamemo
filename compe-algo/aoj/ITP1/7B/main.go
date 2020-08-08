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
	ans := []int{}
	for true {
		n, x := scanInt(sc), scanInt(sc)
		if n == 0 && x == 0 {
			break
		}
		ans = append(ans, _resolve(n, x, 0, 1, 0))
	}
	for _, a := range ans {
		fmt.Println(a)
	}
	return
}

func _resolve(n, x, sum, idx, selectedNum int) int {
	if selectedNum == 3 {
		if sum == x {
			return 1
		}
		return 0
	}
	ans := 0
	for i := idx; i <= n; i++ {
		if sum+i > x {
			continue
		}
		ans += _resolve(n, x, sum+i, i+1, selectedNum+1)
	}
	return ans
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
