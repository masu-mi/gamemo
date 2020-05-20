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

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func gcd(a, b int) int {
	for b > 0 {
		t := a / b
		a, b = b, a-t*b
	}
	return a
}

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (map[direct]int, int) {
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	zero := 0
	groups := map[direct]int{}
	for i := 0; i < n; i++ {
		a, b := scanInt(sc), scanInt(sc)
		if a == 0 && b == 0 {
			zero++
			continue
		} else if a == 0 {
			groups[direct{0, 1}]++
			continue
		} else if b == 0 {
			groups[direct{1, 0}]++
			continue
		}
		d := gcd(abs(a), abs(b))
		if a < 0 {
			a, b = -a, -b
		}
		groups[direct{a / d, b / d}]++
	}
	return groups, zero
}

type direct struct {
	a, b int
}

func (d direct) pair() direct {
	if d.a == 0 && d.b == 1 {
		return direct{a: 1, b: 0}
	} else if d.a == 1 && d.b == 0 {
		return direct{a: 0, b: 1}
	}
	return direct{a: d.b, b: -d.a}
}

const modulo = 1000000007

func resolve(groups map[direct]int, zero int) int {
	pairs := map[direct]direct{}
	for d := range groups {
		p := d.pair()
		_, ok := groups[p]
		if ok {
			pairs[d] = p
			pairs[p] = d
		}
	}
	counted := map[direct]struct{}{}
	result := 1
	for d := range groups {
		if _, ok := counted[d]; ok {
			continue
		}
		var groupNum int
		p, ok := pairs[d]
		if ok {
			num := groups[d]
			baseNum := moduloPow(2, num, modulo)
			num = groups[p]
			pairNum := moduloPow(2, num, modulo)

			groupNum = moduloSub(moduloAdd(baseNum, pairNum, modulo), 1, modulo)
			counted[d] = mark
			counted[p] = mark
		} else {
			num := groups[d]
			groupNum = moduloPow(2, num, modulo)
			counted[d] = mark
		}
		result = moduloMul(result, groupNum, modulo)
	}
	result = moduloAdd(
		result,
		zero-1,
		modulo,
	)
	return result
}

func moduloAdd(a, b, modulo int) int {
	result := a%modulo + b%modulo
	if result < 0 {
		result += modulo
	}
	return result % modulo
}

func moduloSub(a, b, modulo int) int {
	result := a%modulo - b%modulo
	if result < 0 {
		result += modulo
	}
	return result % modulo
}

func moduloMul(a, b, modulo int) int {
	return a % modulo * b % modulo
}

func moduloPow(a, b, modulo int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % modulo
		}
		a = a * a % modulo
		b >>= 1
	}
	return res
}

var mark = struct{}{}

// snip-scan-funcs
