package rbtree

import (
	"fmt"
	"testing"
)

func Test_Insert(t *testing.T) {
	type testCase struct {
		input    []int
		expected *Node
	}
	for idx, test := range []testCase{
		testCase{
			input:    []int{1},
			expected: rootNode(BLACK, 1, nil, nil),
		},
		testCase{
			input:    []int{1, 1},
			expected: rootNode(BLACK, 1, nil, nil),
		},
		testCase{
			input: []int{1, 2},
			expected: rootNode(
				BLACK, 1,
				nil, simpleNode(RED, 2),
			),
		},
		testCase{
			input: []int{2, 1},
			expected: rootNode(
				BLACK, 2,
				simpleNode(RED, 1), nil,
			),
		},
		testCase{
			input: []int{1, 2, 3},
			expected: rootNode(
				BLACK, 2,
				simpleNode(BLACK, 1),
				parentNode(BLACK, 3, nil, nil),
			),
		},
		testCase{
			input: []int{1, 2, 3, 4},
			expected: rootNode(
				BLACK, 2,
				simpleNode(BLACK, 1),
				parentNode(BLACK, 3, nil, simpleNode(RED, 4)),
			),
		},
		testCase{
			input: []int{1, 2, 3, 4, 5},
			expected: rootNode(
				BLACK, 2,
				simpleNode(BLACK, 1),
				parentNode(RED, 4, simpleNode(BLACK, 3), simpleNode(BLACK, 5)),
			),
		},
		testCase{
			input: []int{4, 3, 2, 1},
			expected: rootNode(
				BLACK, 3,
				parentNode(BLACK, 2, simpleNode(RED, 1), nil),
				simpleNode(BLACK, 4),
			),
		},
		testCase{
			input: []int{5, 4, 3, 2, 1},
			expected: rootNode(
				BLACK, 4,
				parentNode(RED, 2, simpleNode(BLACK, 1), simpleNode(BLACK, 3)),
				simpleNode(BLACK, 5),
			),
		},
	} {
		tree := &RBTree{}
		for _, k := range test.input {
			tree.Insert(key(k), k)
			if valid, _ := checkNoBrokenLink(tree.root); !valid {
				t.Errorf("test_id: %d; inserting(%d) break tree's link!!: %s", idx, k, tree.root)
			}
			if !allRanksSame(tree.root) {
				t.Errorf("test_id: %d; inserting(%d) break tree's rank!!: %s", idx, k, tree.root)
			}
		}
		if !tree.root.EqualAsSubTree(test.expected) {
			t.Errorf("test_id: %d; unmatch tree!! %s", idx, tree.root)
		}
	}
}

func Test_recoverBalance(t *testing.T) {
	type testCase struct {
		desc string
		node *Node
	}
	for _, test := range []testCase{
		testCase{
			desc: "root node balanced(its color will be made black)",
			node: rootNode(RED, 1, nil, nil),
		},
		testCase{
			desc: "rotateR(p)",
			node: func() *Node {
				n := parentNode(RED, 0, simpleNode(BLACK, 100), simpleNode(BLACK, 200))
				rootNode(
					BLACK, -100,
					parentNode(
						RED, -200,
						n,
						simpleNode(BLACK, -300),
					),
					simpleNode(BLACK, -400),
				)
				return n
			}(),
		},
		testCase{
			desc: "rotateLR(p)",
			node: func() *Node {
				n := parentNode(RED, 0, simpleNode(BLACK, 100), simpleNode(BLACK, 200))
				rootNode(
					BLACK, -100,
					parentNode(
						RED, -200,
						simpleNode(BLACK, -300),
						n,
					),
					simpleNode(BLACK, -400),
				)
				return n
			}(),
		},
		testCase{
			desc: "rotateRL(p)",
			node: func() *Node {
				n := parentNode(RED, 0, simpleNode(BLACK, 100), simpleNode(BLACK, 200))
				rootNode(
					BLACK, -100,
					simpleNode(BLACK, -200),
					parentNode(
						RED, -300,
						n,
						simpleNode(BLACK, -400),
					),
				)
				return n
			}(),
		},
		testCase{
			desc: "rotateL(p)",
			node: func() *Node {
				n := parentNode(RED, 0, simpleNode(BLACK, 100), simpleNode(BLACK, 200))
				rootNode(
					BLACK, -100,
					simpleNode(BLACK, -200),
					parentNode(
						RED, -300,
						simpleNode(BLACK, -400),
						n,
					),
				)
				return n
			}(),
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			tree := &RBTree{root: findRoot(test.node)}
			tree.recoverBalance(test.node)
			if tree.root.color != BLACK {
				t.Errorf("tree.root.color isn't BLACK!!(\n%s)", tree.root)
			}
			if valid, act := checkNoBrokenLink(tree.root); !valid {
				t.Errorf("tree is broken!!(\nat %s)", act)
			}
		})
	}
}

func Test_Lookup(t *testing.T) {
	tree := &RBTree{}
	tree.Insert(key(10), "found")
	v, err := tree.Lookup(key(10))
	if err != nil {
		t.Errorf("not found!!\n    %s", tree.root)
	}
	if v.(string) != "found" {
		t.Errorf("invalid value was returnd(%v)!!\n    %s", v, tree.root)
	}

	v, err = tree.Lookup(key(11))
	if err == nil {
		t.Errorf("invalid value returned nil!!\n    %s", tree.root)
	}
	if v != nil {
		t.Errorf("invalid value was returnd(%v)!!\n    %s", v, tree.root)
	}
}

func Test_find(t *testing.T) {
	type testCase struct {
		desc               string
		key                int
		top, foundP, found *Node
	}
	for _, test := range []testCase{
		testCase{
			desc:   "returns the node has passed key",
			key:    10,
			top:    rootNode(BLACK, 5, nil, simpleNode(BLACK, 10)),
			foundP: simpleNode(BLACK, 5),
			found:  simpleNode(BLACK, 10),
		},
		testCase{
			desc:   "when found node is root, return nil as its parent",
			key:    5,
			top:    rootNode(BLACK, 5, nil, simpleNode(BLACK, 10)),
			foundP: nil,
			found:  simpleNode(BLACK, 5),
		},
		testCase{
			desc:   "when not found has passed key, found node is nil",
			key:    0,
			top:    rootNode(BLACK, 5, nil, simpleNode(BLACK, 10)),
			foundP: simpleNode(BLACK, 5),
			found:  nil,
		},
		testCase{
			desc:   "when not found has passed key, found parent node is able to be proper place",
			key:    8,
			top:    rootNode(BLACK, 5, nil, simpleNode(BLACK, 10)),
			foundP: simpleNode(BLACK, 10),
			found:  nil,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			tree := &RBTree{root: test.top}
			pl := tree.find(key(test.key))
			actP, actFound := pl.parent, pl.Node()
			assertNode(t, "parent", test.foundP, actP)
			assertNode(t, "", test.found, actFound)
		})
	}
}

func Test_findMax(t *testing.T) {
	type testCase struct {
		desc               string
		top, foundP, found *Node
	}
	for _, test := range []testCase{
		testCase{
			desc: "returns the node has max value and its parent node",
			top: rootNode(
				BLACK, 3,
				parentNode(BLACK, 2, simpleNode(RED, 1), nil),
				simpleNode(BLACK, 4),
			),
			foundP: simpleNode(BLACK, 3),
			found:  simpleNode(BLACK, 4),
		},
		testCase{
			desc:   "when the node has max value is root, returned parent node is nil",
			top:    rootNode(BLACK, 3, nil, nil),
			foundP: nil,
			found:  simpleNode(BLACK, 3),
		},
		testCase{
			desc:   "when the root node is nil, returned node and its parent are nil",
			top:    nil,
			foundP: nil,
			found:  nil,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			actP, actFound := findMax(test.top)
			assertNode(t, "parent", test.foundP, actP)
			assertNode(t, "", test.found, actFound)
		})
	}
}

func Test_findSubstitue(t *testing.T) {
	type testCase struct {
		desc                                 string
		input, expectedParent, expectedFound *Node
	}
	for idx, test := range []testCase{
		testCase{
			desc:           "if target have no child, returns nil as substitute and target as parent",
			input:          rootNode(BLACK, 5, nil, nil),
			expectedParent: simpleNode(BLACK, 5),
			expectedFound:  nil,
		},
		testCase{
			desc:           "if target have no left child, returns right child as substitute and target as parent",
			input:          rootNode(BLACK, 5, nil, simpleNode(RED, 8)),
			expectedParent: simpleNode(BLACK, 5),
			expectedFound:  simpleNode(RED, 8),
		},
		testCase{
			desc:           "if target have left child, returns max node exists under left child as substitute",
			input:          rootNode(BLACK, 5, simpleNode(RED, 3), nil),
			expectedParent: simpleNode(BLACK, 5),
			expectedFound:  simpleNode(RED, 3),
		},
		testCase{
			desc: "if target have left child, returns max node exists under left child as substitute",
			input: rootNode(
				BLACK, 5,
				simpleNode(RED, 3),
				simpleNode(RED, 8),
			),
			expectedParent: simpleNode(BLACK, 5),
			expectedFound:  simpleNode(RED, 3),
		},
		testCase{
			desc: "if target have left child, returns max node exists under left child as substitute",
			input: rootNode(
				BLACK, 5,
				parentNode(RED, 3, simpleNode(BLACK, 2), simpleNode(BLACK, 4)),
				nil,
			),
			expectedParent: simpleNode(RED, 3),
			expectedFound:  simpleNode(BLACK, 4),
		},
		testCase{
			desc: "if target have left child, returns max node exists under left child as substitute",
			input: rootNode(
				BLACK, 6,
				parentNode(
					RED, 3, nil,
					parentNode(
						BLACK, 4, nil,
						simpleNode(RED, 5),
					),
				),
				nil,
			),
			expectedParent: simpleNode(BLACK, 4),
			expectedFound:  simpleNode(RED, 5),
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			actP, actF := findSubstitue(test.input)
			assertNode(t, fmt.Sprintf("parent(idx: %d)", idx), test.expectedParent, actP)
			assertNode(t, fmt.Sprintf("found(idx: %d)", idx), test.expectedFound, actF)
		})
	}
}

type replaceTestCase struct {
	root         *Node
	target       int
	newNode      *Node
	expectedRoot *Node
}

func Test_replace(t *testing.T) {
	for idx, test := range []replaceTestCase{
		replaceTestCase{
			root:         rootNode(BLACK, 5, nil, nil),
			target:       5,
			newNode:      nil,
			expectedRoot: nil,
		},
		replaceTestCase{
			root:         rootNode(BLACK, 5, nil, nil),
			target:       5,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 3, nil, nil),
		},

		replaceTestCase{
			root: rootNode(
				BLACK, 5,
				simpleNode(RED, 2),
				simpleNode(RED, 10),
			),
			target:       5,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 3, nil, nil),
		},

		replaceTestCase{
			root:         rootNode(BLACK, 5, simpleNode(RED, 2), simpleNode(RED, 10)),
			target:       2,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 5, simpleNode(RED, 3), simpleNode(RED, 10)),
		},
		replaceTestCase{
			root:         rootNode(BLACK, 5, simpleNode(RED, 2), simpleNode(RED, 10)),
			target:       10,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 5, simpleNode(RED, 2), simpleNode(RED, 3)),
		},
	} {
		tree := &RBTree{root: test.root}
		pl := tree.find(key(test.target))
		tree.replaceWith(pl, test.newNode)
		if !tree.root.EqualAsSubTree(test.expectedRoot) {
			t.Errorf("TEST CASE(%d) failed! act:%s", idx, tree.root)
		}
	}
}

func Test_update(t *testing.T) {
	for idx, test := range []replaceTestCase{
		replaceTestCase{
			root:         rootNode(BLACK, 5, nil, nil),
			target:       5,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 3, nil, nil),
		},
		replaceTestCase{
			root:         rootNode(BLACK, 5, simpleNode(RED, 2), simpleNode(RED, 10)),
			target:       5,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 3, simpleNode(RED, 2), simpleNode(RED, 10)),
		},
		replaceTestCase{
			root:         rootNode(BLACK, 5, simpleNode(RED, 2), simpleNode(RED, 10)),
			target:       2,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 5, simpleNode(RED, 3), simpleNode(RED, 10)),
		},
		replaceTestCase{
			root:         rootNode(BLACK, 5, simpleNode(RED, 2), simpleNode(RED, 10)),
			target:       10,
			newNode:      simpleNode(BLACK, 3),
			expectedRoot: rootNode(BLACK, 5, simpleNode(RED, 2), simpleNode(RED, 3)),
		},
		replaceTestCase{
			root: rootNode(
				BLACK, 5,
				simpleNode(RED, 2),
				parentNode(
					RED, 10,
					simpleNode(BLACK, 7),
					simpleNode(BLACK, 12),
				),
			),
			target:  10,
			newNode: simpleNode(BLACK, 3),
			expectedRoot: rootNode(
				BLACK, 5,
				simpleNode(RED, 2),
				parentNode(
					RED, 3,
					simpleNode(BLACK, 7),
					simpleNode(BLACK, 12),
				),
			),
		},
	} {
		tree := &RBTree{root: test.root}
		pl := tree.find(key(test.target))
		e := tree.updateValueWith(pl, test.newNode)
		if e != nil {
			t.Errorf("TEST CASE(%d) failed! err:%s", idx, e)
		}
		if !tree.root.EqualAsSubTree(test.expectedRoot) {
			t.Errorf("TEST CASE(%d) failed! act:%s", idx, tree.root)
		}
	}
}

func Test_update_dont_support_with_nil(t *testing.T) {
	tree := &RBTree{root: rootNode(BLACK, 5, nil, nil)}
	pl := tree.find(key(5))
	e := tree.updateValueWith(pl, nil)
	if e == nil {
		t.Errorf("failed!")
	}
}

func Test_recoverRank_Left(t *testing.T) {
	type testCase struct {
		parent   *Node
		t        placeType
		expected *Node
	}
	for idx, test := range []testCase{
		testCase{
			parent: func() *Node {
				return rootNode(
					RED, 10,
					nil,
					simpleNode(
						BLACK, 5,
					),
				)
			}(),
			t:        left,
			expected: rootNode(BLACK, 10, nil, simpleNode(RED, 5)),
		},
		testCase{
			parent: func() *Node {
				p := parentNode(
					RED, 10,
					nil,
					parentNode(
						BLACK, 12,
						simpleNode(RED, 11),
						simpleNode(BLACK, 15),
					),
				)
				rootNode(
					BLACK, 5,
					nil,
					p,
				)
				return p
			}(),
			t: left,
			expected: rootNode(
				BLACK, 5,
				nil,
				parentNode(
					RED, 11,
					simpleNode(BLACK, 10),
					parentNode(BLACK, 12, nil, simpleNode(BLACK, 15)),
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := parentNode(
					RED, 10,
					nil,
					parentNode(
						BLACK, 12,
						simpleNode(BLACK, 11),
						simpleNode(RED, 15),
					),
				)
				rootNode(
					BLACK, 5,
					nil,
					p,
				)
				return p
			}(),
			t: left,
			expected: rootNode(
				BLACK, 5,
				nil,
				parentNode(
					RED, 12,
					parentNode(BLACK, 10, nil, simpleNode(BLACK, 11)),
					simpleNode(BLACK, 15),
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := parentNode(
					RED, 10,
					nil,
					parentNode(
						BLACK, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
				)
				rootNode(
					BLACK, 5,
					nil,
					p,
				)
				return p
			}(),
			t: left,
			expected: rootNode(
				BLACK, 5,
				nil,
				parentNode(
					BLACK, 10,
					nil,
					parentNode(
						RED, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := rootNode(
					RED, 10,
					nil,
					parentNode(
						BLACK, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
				)
				return p
			}(),
			t: left,
			expected: parentNode(
				BLACK, 10,
				nil,
				parentNode(
					RED, 12,
					simpleNode(BLACK, 11),
					simpleNode(BLACK, 15),
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := rootNode(
					RED, 10,
					simpleNode(BLACK, 8),
					parentNode(
						BLACK, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
				)
				return p
			}(),
			t: left,
			expected: parentNode(
				BLACK, 10,
				simpleNode(BLACK, 8),
				parentNode(
					RED, 12,
					simpleNode(BLACK, 11),
					simpleNode(BLACK, 15),
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := rootNode(
					BLACK, 10,
					simpleNode(BLACK, 8),
					parentNode(
						RED, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
				)
				return p
			}(),
			t: left,
			expected: parentNode(
				BLACK, 12,
				parentNode(BLACK, 10,
					simpleNode(BLACK, 8),
					simpleNode(RED, 11),
				),
				simpleNode(BLACK, 15),
			),
		},
	} {
		tree := &RBTree{root: findRoot(test.parent)}
		pl := place{
			t:      test.t,
			tree:   tree,
			parent: test.parent,
		}
		tree.recoverRank(pl)
		act := tree.root
		if !act.EqualAsSubTree(test.expected) {
			t.Errorf("TEST CASE(%d) failed!\n%s", idx, act)
		}
	}
}

func Test_recoverRank_Right(t *testing.T) {
	type testCase struct {
		parent   *Node
		t        placeType
		expected *Node
	}
	for idx, test := range []testCase{
		testCase{
			parent: func() *Node {
				return rootNode(
					RED, 100,
					simpleNode(
						BLACK, 50,
					),
					nil,
				)
			}(),
			t:        right,
			expected: rootNode(BLACK, 100, simpleNode(RED, 50), nil),
		},
		testCase{
			parent: func() *Node {
				p := parentNode(
					RED, 100,
					parentNode(
						BLACK, 12,
						simpleNode(RED, 11),
						simpleNode(BLACK, 15),
					),
					nil,
				)
				rootNode(
					BLACK, 5,
					nil,
					p,
				)
				return p
			}(),
			t: right,
			expected: rootNode(
				BLACK, 5,
				nil,
				parentNode(
					RED, 12,
					parentNode(BLACK, 11, nil, nil),
					parentNode(BLACK, 100, simpleNode(BLACK, 15), nil),
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := parentNode(
					RED, 100,
					parentNode(
						BLACK, 12,
						simpleNode(BLACK, 11),
						simpleNode(RED, 15),
					),
					nil,
				)
				rootNode(
					BLACK, 5,
					nil,
					p,
				)
				return p
			}(),
			t: right,
			expected: rootNode(
				BLACK, 5,
				nil,
				parentNode(
					RED, 15,
					parentNode(BLACK, 12, simpleNode(BLACK, 11), nil),
					simpleNode(BLACK, 100),
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := parentNode(
					RED, 100,
					parentNode(
						BLACK, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
					nil,
				)
				rootNode(
					BLACK, 5,
					nil,
					p,
				)
				return p
			}(),
			t: right,
			expected: rootNode(
				BLACK, 5,
				nil,
				parentNode(
					BLACK, 100,
					parentNode(
						RED, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
					nil,
				),
			),
		},
		testCase{
			parent: func() *Node {
				p := rootNode(
					RED, 100,
					parentNode(
						BLACK, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
					nil,
				)
				return p
			}(),
			t: right,
			expected: parentNode(
				BLACK, 100,
				parentNode(
					RED, 12,
					simpleNode(BLACK, 11),
					simpleNode(BLACK, 15),
				),
				nil,
			),
		},
		testCase{
			parent: func() *Node {
				p := rootNode(
					BLACK, 100,
					parentNode(
						RED, 12,
						simpleNode(BLACK, 11),
						simpleNode(BLACK, 15),
					),
					simpleNode(BLACK, 108),
				)
				return p
			}(),
			t: right,
			expected: parentNode(
				BLACK, 12,
				simpleNode(BLACK, 11),
				parentNode(
					BLACK, 100,
					simpleNode(RED, 15),
					simpleNode(BLACK, 108),
				),
			),
		},
	} {
		tree := &RBTree{root: findRoot(test.parent)}
		pl := place{
			t:      test.t,
			tree:   tree,
			parent: test.parent,
		}
		tree.recoverRank(pl)
		act := tree.root
		if !act.EqualAsSubTree(test.expected) {
			t.Errorf("TEST CASE(%d) failed!\n%s", idx, act)
		}
	}
}

func Test_Delete(t *testing.T) {
	type testCase struct {
		insertOrder []int
		target      int
	}
	for idx, test := range []testCase{
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      1,
		},
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      2,
		},
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      3,
		},
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      4,
		},
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      5,
		},
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      6,
		},
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      7,
		},
		testCase{
			insertOrder: []int{3, 8, 5, 1, 2, 4, 6, 7},
			target:      8,
		},
	} {
		tree := &RBTree{}
		for _, k := range test.insertOrder {
			tree.Insert(key(k), fmt.Sprintf("%d", k))
		}
		tree.Delete(key(test.target))
		if !allRanksSame(tree.root) {
			t.Errorf("test_id: %d; tree's rank broken!!: %s", idx, tree.root)
		}
		for _, k := range test.insertOrder {
			if k != test.target {
				_, e := tree.Lookup(key(k))
				if e != nil {
					t.Errorf("test_id: %d; lost (%d)!: %s", idx, k, tree.root)
				}
			} else {
				_, e := tree.Lookup(key(k))
				if e == nil {
					t.Errorf("test_id: %d; not deleted (%d)!: %s", idx, k, tree.root)
				}
			}
		}
	}
}
