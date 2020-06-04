package lca

type tree struct {
	size   int
	parent []int
	childs [][]int
}

func newTree(size int) *tree {
	return &tree{size: size, parent: make([]int, size), childs: make([][]int, size)}
}

func (t *tree) addEdge(p, c int) {
	t.parent[c] = p
	t.childs[p] = append(t.childs[p], c)
}

func (t *tree) isRoot(i int) bool {
	return t.parent[i] == i
}

func (t *tree) lca(x, y int) int {
	visited := newIntSet()
	{
		i := x
		visited.add(i)
		for !t.isRoot(i) {
			visited.add(i)
			i = t.parent[i]
		}
	}
	{
		i := y
		if visited.doesContain(i) {
			return i
		}
		for !t.isRoot(i) {
			if visited.doesContain(i) {
				return i
			}
			i = t.parent[i]
		}
	}
	return -1
}

type intSet map[int]none

func newIntSet() intSet {
	return map[int]none{}
}

func (s intSet) add(i int) (added bool) {
	_, ok := s[i]
	added = !ok
	s[i] = mark
	return
}

func (s intSet) remove(i int) (removed bool) {
	_, removed = s[i]
	delete(s, i)
	return
}

func (s intSet) doesContain(i int) bool {
	_, ok := s[i]
	return ok
}

func (s intSet) size() int {
	return len(s)
}

func (s intSet) members() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for k := range s {
			ch <- k
		}
	}()
	return ch
}

type none struct{}

var mark none
