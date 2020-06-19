package main
// packed from [splaytree_debug.go splaytree.go main.rmq.go] with goone.

import (
	"fmt"
	"io"
	"strings"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type debugNodeSlice struct {
	l   []SplayTreeNode
	ids map[*SplayTreeNode]int
}

func newDebugNodeSlice(n int) debugNodeSlice {
	s := debugNodeSlice{
		l:   make([]SplayTreeNode, n),
		ids: map[*SplayTreeNode]int{},
	}
	for i := range s.l {
		s.ids[&(s.l[i])] = i
	}
	return s
}

func (t *debugNodeSlice) display(w io.Writer, root *SplayTreeNode) {
	cur := root
	t._display(w, cur, 0)
}

func (t *debugNodeSlice) _display(w io.Writer, cur *SplayTreeNode, indent int) {
	if cur == nil {
		fmt.Fprintf(w, "[nil]\n")
		return
	}
	id := t.ids[cur]
	fmt.Fprintf(w, "[id: %d, value: %d, size: %d, %p]\n", id, cur.value, cur.size, cur)

	indent += 4
	fmt.Fprintf(w, "%s%d.LEFT ::", strings.Repeat(" ", indent), id)
	t._display(w, cur.left, indent)
	fmt.Fprintf(w, "%s%d.Right::", strings.Repeat(" ", indent), id)
	t._display(w, cur.right, indent)
}

type SplayTreeNode struct {
	parent, left, right *SplayTreeNode

	size    int
	value   int
	minimum int
}

func newSplayTreeNode(v int) *SplayTreeNode {
	return &SplayTreeNode{value: v}
}

func (node *SplayTreeNode) rotate() {
	var pp, p, c *SplayTreeNode
	p = node.parent
	pp = p.parent
	if pp != nil && pp.left == p {
		pp.left = node
	}
	if pp != nil && pp.right == p {
		pp.right = node
	}
	node.parent = pp
	if p.left == node {
		c = node.right
		node.right = p
		p.left = c
	} else {
		c = node.left
		node.left = p
		p.right = c
	}
	p.parent = node
	if c != nil {
		c.parent = p
	}
	p.update()
	node.update()
}

func (node *SplayTreeNode) state() int {
	if node.parent == nil {
		return 0
	}
	if node.parent.left == node {
		return -1
	} else if node.parent.right == node {
		return 1
	}
	return 0
}

func (node *SplayTreeNode) splay() {
	for node.state() != 0 {
		if node.parent.state() == 0 {
			node.rotate()
			return
		}
		if node.state() == node.parent.state() {
			node.parent.rotate()
			node.rotate()
		} else {
			node.rotate()
			node.rotate()
		}
	}
}

func (node *SplayTreeNode) update() {
	node.size = 1
	node.minimum = node.value
	if node.left != nil {
		node.size += node.left.size
		node.minimum = min(node.minimum, node.left.minimum)
	}
	if node.right != nil {
		node.size += node.right.size
		node.minimum = min(node.minimum, node.right.minimum)
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (node *SplayTreeNode) Get(idx int) *SplayTreeNode {
	cur := node
	for true {
		lSize := 0
		if cur.left != nil {
			lSize = cur.left.size
		}
		if idx < lSize {
			cur = cur.left
		}
		if lSize == idx {
			cur.splay()
			return cur
		}
		if idx > lSize {
			cur = cur.right
			idx = idx - lSize - 1
		}
	}
	return nil
}

func mergeSplayNode(l, r *SplayTreeNode) (root *SplayTreeNode) {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	root = l.Get(l.size - 1)
	root.right = r
	r.parent = root
	root.update()
	return root
}

func (node *SplayTreeNode) Split(leftCount int) (left *SplayTreeNode, right *SplayTreeNode) {
	root := node.Get(leftCount)
	lRoot, rRoot := root.left, root

	rRoot.left = nil
	if lRoot != nil {
		lRoot.parent = nil
	}
	rRoot.update()
	return lRoot, rRoot
}

func splayTreeInsert(t *SplayTreeNode, idx int, node *SplayTreeNode) (root *SplayTreeNode) {
	l, r := t.Split(idx)
	return mergeSplayNode(
		mergeSplayNode(l, node),
		r,
	)
}

func removeSplayTree(t *SplayTreeNode, idx int) (root *SplayTreeNode, removed *SplayTreeNode) {
	target := t.Get(idx)
	l, r := target.left, target.right
	if l != nil {
		l.parent = nil
	}
	if r != nil {
		r.parent = nil
	}
	target.left = nil
	target.right = nil
	target.update()
	return mergeSplayNode(l, r), target
}

const (
	initialBufSize = 100000
	maxBufSize     = 1000000
)

var sc *bufio.Scanner

func initScanner(r io.Reader) *bufio.Scanner {
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
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

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
