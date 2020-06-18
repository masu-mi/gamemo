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
	q := scanInt(sc)
	table := newDebugNodeSlice(200001)
	for i := 0; i < n; i++ {
		table.l[i].parent = &table.l[i+1]
		table.l[i+1].left = &table.l[i]
		table.l[i].update()
		table.l[i+1].update()
	}
	var root *SplayTreeNode = &table.l[n]
	for i := 0; i < n; i++ {
		table.l[i].value = scanInt(sc)
		table.l[i].update()
	}
	table.l[n].update()
	for i := 0; i < q; i++ {
		switch scanInt(sc) {
		case 0:
			root = shift(root, scanInt(sc), scanInt(sc))
		case 1:
			l := scanInt(sc)
			r := scanInt(sc)
			tmp, rRoot := root.Split(r + 1)
			lRoot, cRoot := tmp.Split(l)
			fmt.Println(cRoot.minimum)
			root = mergeSplayNode(mergeSplayNode(lRoot, cRoot), rRoot)
		case 2:
			root = root.Get(scanInt(sc))
			root.value = scanInt(sc)
			root.update()
		}
	}
}

func shift(root *SplayTreeNode, l int, r int) *SplayTreeNode {
	root, node := removeSplayTree(root, r)
	return splayTreeInsert(root, l, node)
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
