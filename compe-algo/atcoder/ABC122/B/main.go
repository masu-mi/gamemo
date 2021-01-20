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
	str := scanString(sc)
	s := newByteSet("ACGT")
	var maxLen, l int
	for i := 0; i < len(str); i++ {
		if s.doesContain(str[i]) {
			l++
		} else {
			if maxLen < l {
				maxLen = l
			}
			l = 0
		}
	}
	if maxLen < l {
		maxLen = l
	}
	return maxLen
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
