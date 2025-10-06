package main

// 	Write a function that will accept a positive number n as input. It will
//   print numbers m from 1 to n. If m is divisible by 3, it will
//   print fizz. If it's divisible by 5, it will print bizz. If it's
//   divisible by 15, it will print fizzbizz. Otherwise, it will just
//   print the number.
import "fmt"

func main() {
	fizzbizz(20)
}
func fizzbizz(x int) {
	for i := 1; i <= x; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBizz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Bizz")
		} else {
			fmt.Println(i)
		}
	}
}
