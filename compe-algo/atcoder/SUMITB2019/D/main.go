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
	_ = scanInt(sc)
	s := scanString(sc)

	lastPos := map[byte]int{}
	for i := 0; i < len(s); i++ {
		lastPos[s[i]] = i
	}

	accessed := map[string]struct{}{}
	var num int
	for i := 0; i+2 < len(s); i++ {
		if _, ok := accessed[string([]byte{s[i]})]; ok {
			continue
		}
		for j := i + 1; j+1 < len(s); j++ {
			if _, ok := accessed[string([]byte{s[i], s[j]})]; ok {
				continue
			}
			num += countVariationNum(j+1, lastPos)
			accessed[string([]byte{s[i], s[j]})] = struct{}{}
		}
		accessed[string([]byte{s[i]})] = struct{}{}
	}
	return num
}

func countVariationNum(k int, lastPos map[byte]int) int {
	num := 0
	for _, p := range lastPos {
		if k <= p {
			num++
		}
	}
	return num
}

func newByteSet(input string) byteSet {
	s := newSet()
	for i := 0; i < len(input); i++ {
		s.add(input[i])
	}
	return s
}

type byteSet map[byte]none

func newSet() byteSet {
	return make(map[byte]none)
}

func (s byteSet) add(item byte) {
	s[item] = mark
}

func (s byteSet) doesContain(item byte) bool {
	_, ok := s[item]
	return ok
}

func (s byteSet) size() int {
	return len(s)
}

var mark none

type none struct{}

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
