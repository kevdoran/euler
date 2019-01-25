/*
 * Solution to Project Euler Problem 3 https://projecteuler.net/problem=3
 *
 *   The prime factors of 13195 are 5, 7, 13 and 29.
 *   What is the largest prime factor of the number 600851475143 ?
 *
 * Run using:
 *   go run 003.go
 */

package main

import "fmt"

func main() {
	fmt.Println(solve(600851475143))
}


func solve(target int) int {
	return findLargestPrimeDivisor(target)
}


// Find divisors < sqrt(n) in ascending order, for each dividend, check if it is prime.
// The first prime dividend is the largest divisor. If no dividends are prime, try divisors in descending order.
func findLargestPrimeDivisor(n int) int {
	var smallFactors []int
	for i := 2; i * i <= n; i++ {
		smallFactor := i
		if n % smallFactor == 0 {
			largeFactor := n / i
			if isPrime(largeFactor) {
				return largeFactor
			} else {
				smallFactors = append(smallFactors, smallFactor)
			}
		}
	}
	for i := len(smallFactors)-1 ; i >= 0 ; i-- {
		if isPrime(smallFactors[i]) {
			return smallFactors[i]
		}
	}
	return n  // no prime factors, n is prime
}


// This optimized isPrime algorithm is based on the
// principal that all primes > 3 are of the form (6k ± 1).
//   All integers can be expressed as (6k + i) for some integer k
//   and for i = −1, 0, 1, 2, 3, or 4, and for i = 0, 2, 3, 4 the
//   result is divisible by 2 or 3.
// Lastly, we only have to search up to sqrt(n) before we start checking duplicate factors.
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n % 2 == 0 || n % 3 == 0 {
		return false
	}
	for i := 5 ; i * i <= n ; i += 6 {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
	}
	return true
}