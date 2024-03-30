package main

import "testing"

func TestIsPrime(t *testing.T) {
	knownPrimes := map[int]bool{
		1: false, 2: true, 3: true,
		4: false, 5: true, 6: false,
		7: true, 8: false, 9: false,
		10: false, 11: true, 12: false,
		13: true, 14: false, 15: false,
	}

	for k, v := range knownPrimes {
		if isPrime(k) != v {
			t.Fatalf("isPrime(%d) != %v", k, v)
		}
	}
}
