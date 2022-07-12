package main

// normal recursion

func fibNormalRecursion(i int) int {
	if i <= 2 {
		return 1
	}
	return fibNormalRecursion(i-1) + fibNormalRecursion(i-2)
}
