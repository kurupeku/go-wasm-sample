package calc

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciMemorized(n int) int {
	cache := make(map[int]int)
	return fibonacci(n, &cache)
}

func fibonacci(n int, c *map[int]int) int {
	if n < 2 {
		return n
	}

	cv := *c
	if v, ok := cv[n]; ok {
		return v
	}

	cv[n] = fibonacci(n-1, c) + fibonacci(n-2, c)
	return cv[n]
}
