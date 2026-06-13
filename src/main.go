package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Day 2 - Variables, Mutability, and Types
	// var name string = "Mark"
	// var age int32 = 123
	// var heightMeters float32 = 3.5
	// var isDevops bool = true

	// fmt.Println(name, age, heightMeters, isDevops)

	// Day 3 - FizzBuzz
	iter := 20

	for i := 1; i <= iter; i++ {
		var output string = strconv.Itoa(i)

		if i%3 == 0 {
			output = "Fizz"
		}
		if i%5 == 0 {
			output = "Buzz"
		}
		if i%3 == 0 && i%5 == 0 {
			output = "FizzBuzz"
		}

		fmt.Println(output)
	}
}
