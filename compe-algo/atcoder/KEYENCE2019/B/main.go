package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	resolve(parseProblem())
}

func parseProblem() string {
	return scanString(sc)
}

const (
	Pre = 0 + iota
	Removing
	Post
)

func resolve(str string) {
	const target = "keyence"
	if strings.HasSuffix(str, target) {
		fmt.Println("YES")
		return
	}
	if str[0:len(target)] == target {
		fmt.Println("YES")
		return
	}
	i := 0
	for ; i < len(str); i++ {
		if str[i] != target[i] {
			break
		}
	}
	if strings.HasSuffix(str, target[i:len(target)]) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
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
