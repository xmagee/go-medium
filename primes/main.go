package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("prime number checker")

	fmt.Println("input a number:")
	var input int
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if isPrime(input) {
		fmt.Printf("%v is prime\n", input)
	} else {
		fmt.Printf("%v is not prime\n", input)
	}
}
