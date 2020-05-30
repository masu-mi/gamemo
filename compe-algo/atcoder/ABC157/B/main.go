package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) ([][]int, intSet) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	a := nextMatrix(3, 3, sc)
	n := nextInt(sc)
	s := newIntSet()
	for i := 0; i < n; i++ {
		s.add(nextInt(sc))
	}
	return a, s
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

func nextMatrix(n, m int, sc *bufio.Scanner) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = nextInt(sc)
		}
	}
	return
}
func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return int(a)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func resolve(mx [][]int, b intSet) {
	for i := 0; i < 3; i++ {
		okL, okC := true, true
		for j := 0; j < 3; j++ {
			okL = okL && b.doesContain(mx[i][j])
			okC = okC && b.doesContain(mx[j][i])
		}
		if okL || okC {
			fmt.Println("Yes")
			return
		}
		// diagonal-l
		// diagonal-r
	}
	okL, okR := true, true
	for i := 0; i < 3; i++ {
		okL = okL && b.doesContain(mx[i][i])
		okR = okR && b.doesContain(mx[i][2-i])
	}
	if okL || okR {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
	return
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
