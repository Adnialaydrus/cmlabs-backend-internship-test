package main

import (
	"fmt"
)

func main() {

	for counter := 1; counter <= 100; counter++ {
		switch {
		case counter%3 == 0 && counter%5 == 0:
			fmt.Println("FizzBuzz")
		case counter%3 == 0:
			fmt.Println("Fizz")
		case counter%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(counter)
		}
	}
}
