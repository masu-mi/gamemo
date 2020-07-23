package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	fmt.Println(resolve())
}

func resolve() int {
	str := scanString(sc)
	lastPos := map[byte]int{}
	maxSectionLen := map[byte]int{}
	for i := range str {
		b := str[i]
		var l int
		if _, ok := lastPos[b]; ok {
			l = i - lastPos[b] - 1
		} else {
			l = i
		}
		updateWithMax(maxSectionLen, b, l)
		lastPos[b] = i
	}
	for b, p := range lastPos {
		l := len(str) - p - 1
		updateWithMax(maxSectionLen, b, l)
	}
	return selectMin(maxSectionLen)
}

func updateWithMax(m map[byte]int, i byte, v int) bool {
	update := m[i] <= v
	if update {
		m[i] = v
	}
	return update
}

func selectMin(m map[byte]int) int {
	min := math.MaxInt32
	for _, secLen := range m {
		if min > secLen {
			min = secLen
		}
	}
	return min
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
