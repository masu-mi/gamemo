package gocom

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
