package gocom

import "bufio"

func getComulative(a []int) (z []int) {
	z = make([]int, len(a)+1)
	(&reducer{}).reduce(&z, a)
	return
}

func nextComulative(sc *bufio.Scanner, n int, init int, op func(v, acc int) int) (a, z []int) {
	a = make([]int, n)
	z = make([]int, n+1)
	r := &reducer{init: init, op: op}
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
		r.update(&z, a, i)
	}
	return a, z
}

type reducer struct {
	// TODO multi dimension
	init, tmp int
	op        func(v, acc int) int
}

func (r *reducer) update(z *[]int, a []int, idx int) {
	if r.op != nil {
		r.tmp = r.op(r.tmp, a[idx])
	} else {
		r.tmp += a[idx]
	}
	(*z)[idx+1] = r.tmp
}

func (r *reducer) reduce(z *[]int, a []int) {
	for i := 0; i < len(a); i++ {
		r.update(z, a, i)
	}
}
