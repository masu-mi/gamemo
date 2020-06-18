package main

import (
	"fmt"
	"io"
	"strings"
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
