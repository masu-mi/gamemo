package gocom

func permutations(l, offset int) chan []int {
	ch := make(chan []int)
	go func() {
		dfsPermutations(0, offset, make([]bool, l), []int{}, func(perm []int) bool {
			ch <- perm
			return false
		})
		close(ch)
	}()
	return ch
}

func dfsPermutations(pos, off int, used []bool, perm []int, atLeaf func(perm []int) (halt bool)) (halt bool) {
	l := len(used)
	if pos == l {
		p := append(perm[:0:0], perm...) // copy for safe
		return atLeaf(p)
	}

	for i := 0; i < l; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		if dfsPermutations(pos+1, off, used, append(perm, i+off), atLeaf) {
			return true
		}
		used[i] = false
	}
	return false
}
