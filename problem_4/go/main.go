// Project Euler : Problem 4
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
)

func main() {
	max := 100001

	for i := 999; i >= 100; i-- {
		if max >= i*999 {
			break
		}

		for j := 999; j >= i; j-- {
			number := i * j
			if max < number && isPalindrome(number) {
				max = number
			}
		}
	}

	fmt.Println("The max palindrome for the product of two 3 digit numbers is :", max)
}

func isPalindrome(number int) bool {
	original := number
	var reversed int

	for number > 0 {
		remainder := number % 10
		reversed *= 10
		reversed += remainder
		number /= 10
	}

	return reversed == original
}
