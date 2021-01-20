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
	fmt.Println(resolve())
}

func resolve() int {
	n := scanInt(sc)
	ss := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s := []byte(scanString(sc))

		sort.Sort(byteSlice(s))
		ss = append(ss, string(s))
	}
	sort.Sort(sort.StringSlice(ss))
	counts := map[string]int{}
	var total int
	for i := 0; i < len(ss); i++ {
		total += counts[ss[i]]
		counts[ss[i]]++
	}
	return total
}

type byteSlice []byte

func (r byteSlice) Len() int           { return len(r) }
func (r byteSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byteSlice) Less(i, j int) bool { return r[i] < r[j] }

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
