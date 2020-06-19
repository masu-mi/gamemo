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
	input := []int{0, 10, 20, 20, 30, 50, 60, 100, 1000}
	gen := func(b int) func(v int) bool {
		return func(v int) bool {
			return v <= b
		}
	}
	for _, v := range []int{0, 1, 25, 40, 70, 100, 101, 10000} {
		fmt.Printf("v:%d\n", v)
		id := binarySearch(input, gen(v))
		if id >= 0 && id < len(input) {
			fmt.Printf("l[%d] = %d\n", id, input[id])
		} else {
			fmt.Printf("l[%d] = [out of bound]\n", id)
		}
	}
	return
}

func binarySearch(l []int, pred func(v int) bool) (bound int) {
	ok := -1
	ng := len(l)
	for abs(ok-ng) > 1 {
		mid := (ok + ng) >> 1
		fmt.Println(ok, mid, ng)
		if pred(l[mid]) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
