// Project Euler : Problem 3
// What is the largest prime factor of the number 600851475143?

package main

import (
  "fmt"
  "math"
)

func main() {
  number := 600851475143

  for number%2 == 0 {
    fmt.Print(2)
    number /= 2
  }

  for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
    for number%i == 0 {
      number /= i
    }
  }

  fmt.Println(number)
}
