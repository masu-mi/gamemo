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
	claims := map[int]map[int]int{}
	n := scanInt(sc)
	for i := 0; i < n; i++ {
		a := scanInt(sc)
		m := map[int]int{}
		for j := 0; j < a; j++ {
			x, y := scanInt(sc)-1, scanInt(sc)
			m[x] = y
		}
		claims[i] = m
	}
	maxNum := 0
	for i := 0; i < 1<<uint(n); i++ {
		num := onesCount(uint64(i))
		if maxNum > num {
			continue
		}
		if acceptable(claims, n, i) {
			maxNum = num
		}
	}
	return maxNum
}

func acceptable(claims map[int]map[int]int, num, state int) bool {
	for i := 0; i < num; i++ {
		if state>>uint(i)&1 != 0 {
			for k, v := range claims[i] {
				if state>>uint(k)&1 != v {
					return false
				}
			}
		}
	}
	return true
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/bits_util.go] with goone.

func onesCount(x uint64) (num int) {
	const m0 = 0x5555555555555555
	const m1 = 0x3333333333333333
	const m2 = 0x0f0f0f0f0f0f0f0f

	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}

func ntz(bits uint64) (num int) {

	return onesCount(bits&(-bits) - 1)
}

func nlz(bits uint64) (num int) {

	bits = bits | (bits >> 1)
	bits = bits | (bits >> 2)
	bits = bits | (bits >> 4)
	bits = bits | (bits >> 8)
	bits = bits | (bits >> 16)
	bits = bits | (bits >> 32)
	return onesCount(^bits)
}

func refBit(i uint64, b uint) int {
	return int((i >> b) & 1)
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
