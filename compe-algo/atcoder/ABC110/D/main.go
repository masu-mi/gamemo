package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println(parseProblem(os.Stdin))
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
	n, m := scanInt(sc), scanInt(sc)
	t := primeFactories(m)
	mod := int(1e9) + 7
	num := 1
	moduloCombiInit(mod)
	for _, ex := range t {
		num = moduloMul(
			num,
			moduloCombi(n+ex-1, n-1, mod),
			mod,
		)
	}
	return num
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

func moduloDiv(a, b, modulo int) int {
	return a % moduloInv(b, modulo) % modulo
}

func moduloInv(a, modulo int) int {
	b := modulo
	u, v := 1, 0
	for b > 0 {
		t := a / b
		a, b = b, a-t*b
		u, v = v, u-t*v
	}
	u %= modulo
	if u < 0 {
		u += modulo
	}
	return u
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

const length = 510000

var (
	_fac  = map[int][]int{}
	_finv = map[int][]int{}
	_inv  = map[int][]int{}
)

func moduloCombiInit(modulo int) {
	fac := make([]int, length)
	finv := make([]int, length)
	inv := make([]int, length)

	defer func() {
		_fac[modulo] = fac
		_finv[modulo] = finv
		_inv[modulo] = inv
	}()

	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < length; i++ {
		fac[i] = fac[i-1] * i % modulo
		inv[i] = modulo - inv[modulo%i]*(modulo/i)%modulo
		finv[i] = finv[i-1] * inv[i] % modulo
	}
}

func moduloCombi(n, k, modulo int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return _fac[modulo][n] * (_finv[modulo][k] * _finv[modulo][n-k] % modulo) % modulo
}

func moduloLog(a, b, modulo int) int {

	a %= modulo
	b %= modulo
	m := int(math.Sqrt(float64(modulo)))

	values := map[int]int{}
	val := 1
	for i := 0; i < m+2; i++ {
		if _, ok := values[val]; !ok {
			values[val] = i
		}
		val = moduloMul(val, a, modulo)
	}

	compound := moduloInv(moduloPow(a, m, modulo), modulo)
	val = b
	for i := 0; i < m+2; i++ {
		if l, ok := values[val]; ok {
			return (i*m%modulo + l) % modulo
		}
		val = moduloMul(val, compound, modulo)
	}
	return -1
}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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

func primeFactories(n int) (terms map[int]int) {
	terms = map[int]int{}
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
