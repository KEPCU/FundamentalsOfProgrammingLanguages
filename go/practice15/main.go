package main

import (
  "fmt"
)

func main() {
  fmt.Println("--EXERCISE 1 (Factorial)")
  print("input number : ")
  var number int 
	fmt.Scanln(&number)
  fmt.Println(Factorial(number))

  fmt.Println("--EXERCISE 2 (IsPrime)")
  print("input number : ")
	fmt.Scanln(&number)
  fmt.Println(IsPrime(number))

  fmt.Println("--EXERCISE 3 (Divisors)")
  print("input number : ")
	fmt.Scanln(&number)
  Divisors(number)

  fmt.Println("--EXERCISE 4 (Euler)")
  fmt.Println(Euler())
}

//--EXERCISE 1
func Factorial(number int) int {
	if number < 2 {
		return 1
	}
	return number * Factorial(number - 1)
}

//--EXERCISE 2
func IsPrime(number int) string {
  var count = 0
  for i := 2; i < number; i++ {
    if number%i == 0 {
		  count++
  	}

    if count > 1 {
		  return "no prime"
  	}
  }
	return "prime"
}

//--EXERCISE 3
func Divisors(number int) {
  for i := 1; i <= number; i++ {
    if number%i == 0 {
		  fmt.Println(i)
  	}
  }
}

//--EXERCISE 4
func Euler() float64 {
  var result = 1.0
  var next float64
  var current float64
  for i := 1; i < i+1; i++ {
    current = 1.0 / float64(Factorial(i))
    next = 1.0 / float64(Factorial(i+1))
    result += current
    if current - next < 0.01 {
      fmt.Println("diference ",current - next)
      return result + next
    }
  }
  return result
}

