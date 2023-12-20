package main

import "fmt"

// function fn(i):
//     print(i)
//     fn(i + 1)
//     return

func recursionNoBaseCase(n int) {
	print(n)
	recursionNoBaseCase(n + 1)
	return
}

func recursionBaseCase(n int) {
	if n > 10 {
		return
	}

	print(n)
	recursionBaseCase(n + 1)
	fmt.Printf("End of call where i = %d\n", n)
	return
}

func fibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}

	oneBack := fibonacciRecursion(n - 1)
	twoBack := fibonacciRecursion(n - 2)

	return oneBack + twoBack
}
