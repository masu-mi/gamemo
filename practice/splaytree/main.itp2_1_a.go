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
	n := scanInt(sc)
	nodes := make([]SplayTreeNode, 200001)
	for i := 0; i < n; i++ {
		nodes[i].parent = &nodes[i+1]
		nodes[i+1].left = &nodes[i]
		nodes[i].update()
		nodes[i+1].update()
	}
	var root *SplayTreeNode = &nodes[n]
	idx := 0
	for i := 0; i < n; i++ {
		switch scanInt(sc) {
		case 0:
			root = root.Get(idx)
			root.value = scanInt(sc)
			idx++
		case 1:
			root = root.Get(scanInt(sc))
			fmt.Println(root.value)
		case 2:
			idx--
		}
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
