package gocom

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
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
