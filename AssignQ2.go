package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the number of terms in the Fibonacci series: ")
	fmt.Scan(&n)

	fmt.Println("Fibonacci Series:")
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
