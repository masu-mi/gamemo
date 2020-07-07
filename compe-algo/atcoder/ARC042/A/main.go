package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

type pair struct{ o, v int }

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].o < p[j].o
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func resolve() {
	n, m := scanInt(sc), scanInt(sc)
	list := make([]pair, n)
	for i := 0; i < n; i++ {
		list[i] = pair{i, i}
	}
	for i := 1; i <= m; i++ {
		idx := scanInt(sc)
		list[idx-1].o = -i
	}
	sort.Sort(pairs(list))
	for _, v := range list {
		fmt.Println(v.v + 1)
	}
	return
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
