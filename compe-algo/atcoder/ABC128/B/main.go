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
	resolve(parseProblem())
}

func parseProblem() int {
	return scanInt(sc)
}

type restaurant struct {
	id, p int
	s     string
}
type restaurants []restaurant

func (r restaurants) Len() int {
	return len(r)
}

func (r restaurants) Less(i, j int) bool {
	l := ([]restaurant)(r)
	if l[i].s == l[j].s {
		return l[i].p > l[j].p
	}
	return l[i].s < l[j].s
}

func (r restaurants) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func resolve(n int) {
	ts := []restaurant{}
	for i := 1; i <= n; i++ {
		r := restaurant{
			id: i,
			s:  scanString(sc),
			p:  scanInt(sc),
		}
		ts = append(ts, r)
	}
	sort.Sort(restaurants(ts))
	for i := range ts {
		fmt.Println(ts[i].id)
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
