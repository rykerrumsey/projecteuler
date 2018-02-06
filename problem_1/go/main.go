// Project Euler : Problem 1
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func main() {

  var multiple5 []int
  var multiple3 []int

  // get all the multiples of 5 under 1000
  for i := 999; i > 0; i-- {
    if i%5 == 0 {
      multiple5 = append(multiple5, i)
    }
  }

  // get all the multiples of 3 under 1000
  for i := 999; i > 0; i-- {
    if i%3 == 0 {
      multiple3 = append(multiple3, i)
    }
  }

  // filter out the multiples of 3 and 5
  newMultiple5 := Filter(multiple5, divisibleBy3)

  // add the multiples of 5
  total5 := getSumSlice(newMultiple5)

  // add the multiples of 3
  total3 := getSumSlice(multiple3)

  fmt.Printf("The total sum of all the multiples is : %v\n", total5 + total3)
}

func getSumSlice(values []int) int {
  var temp int
  for _, v := range values {
    temp += v
  }
  return temp
}

func Filter(vs []int, f func(int) bool) []int {
    vsf := make([]int, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func divisibleBy3(number int) bool {
  if number%3 == 0 {
    return false
  } else {
    return true
  }
}
