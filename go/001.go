/*
 * Solution to Project Euler Problem 1 https://projecteuler.net/problem=1
 *
 *   If we list all the natural numbers below 10 that are multiples of 3 or 5,
 *   we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 *   Find the sum of all the multiples of 3 or 5 below 1000.
 *
 * Run using:
 *   go run 001.go
 */

package main

import "fmt"

func main() {
	fmt.Println(solve(1000))
	//fmt.Println(solveOptimally(1000))
}

func solve(target int) int {
	sum := 0
	for i := 0; i < target; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			sum = sum + i
		}
	}
	return sum
}

func solveOptimally(target int) int {
	return sumDivisibleBy(1000, 3) + sumDivisibleBy(1000, 5) - sumDivisibleBy(1000, 15)
}

func sumDivisibleBy(target, n int) int {
	p := target / n
	return n * ( p * (p+1)) / 2
}
