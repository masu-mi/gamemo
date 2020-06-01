package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) int {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	return n
}

func resolve(n int) int {
	// if n < 2 {
	// 	return 0
	// }
	terms := map[int]int{}
	num := n
	for f := 2; f*f < n; f++ {
		if num%f != 0 {
			continue
		}
		ex := 0
		for num%f == 0 {
			num /= f
			ex++
		}
		terms[f] = ex
	}
	if num > 1 {
		terms[num] = 1
	}
	result := 0
	for k, iN := range terms {
		if iN == 0 {
			continue
		}
		for j := 1; j <= iN; j++ {
			d := pow(k, j)
			if n%d != 0 {
				break
			}
			n /= d
			result++
		}
	}
	return result
}

func pow(n, i int) (res int) {
	res = 1
	for i > 0 {
		if i&1 == 1 {
			res *= n
		}
		i >>= 1
		n *= n
	}
	return res
}

func numP(n, p int) (i int) {
	for n > 0 {
		if n%p != 0 {
			break
		}
		i++
		n = n / p
	}
	return i
}

func eratosthenes(n int) []bool {
	t := make([]bool, n)
	for i := 2; i < len(t); i++ {
		t[i] = true
	}
	for i := 2; i < len(t); i++ {
		if !t[i] {
			continue
		}
		for j := 2; j*i < len(t); j++ {
			t[j*i] = false
		}
	}
	return t
}

func primes(n int) chan int {
	ch := make(chan int)
	go func() {
		t := eratosthenes(n)
		for i := 2; i < len(t); i++ {
			if !t[i] {
				continue
			}
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
