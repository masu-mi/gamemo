package main

type SplayTreeNode struct {
	parent, left, right *SplayTreeNode

	size    int
	value   int // TODO change to interface{}
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
