package calc

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}

func FibonacciMemorized(n int) int {
	cache := map[int]int{}
	return fibonacci(n, cache)
}

func fibonacci(n int, c map[int]int) int {
	if v, ok := c[n]; ok {
		return v
	}

	return fibonacci(n-2, c) + fibonacci(n-1, c)
}
