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
	if resolve(parseProblem()) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func parseProblem() (int, []task) {
	n := scanInt(sc)
	result := make([]task, n)
	for i := 0; i < n; i++ {
		result[i] = task{
			take:  scanInt(sc),
			limit: scanInt(sc),
		}
	}
	return n, result
}

type task struct{ take, limit int }
type tasks []task

func (t tasks) Len() int {
	return len(t)
}

func (t tasks) Less(i, j int) bool {
	return t[i].limit < t[j].limit
}

func (t tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func resolve(n int, ts []task) bool {
	sort.Sort(tasks(ts))
	sum := 0
	for _, t := range ts {
		sum += t.take
		if t.limit < sum {
			return false
		}
	}
	return true
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
