package sieve_test

import "cirello.io/sieve"

func Example() {
	cache := sieve.New[string](3)
	cache.Access("A")
	cache.Access("B")
	cache.Access("C")
	cache.Access("D")
	cache.Show()

	// Output:
	// D (Visited: false) -> C (Visited: false) -> B (Visited: false)
}
