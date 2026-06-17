package main

import "fmt"

type Server struct {
	Name     string
	CPUCores uint32
	IsActive bool
}

type Light int

const (
	Red Light = iota
	Yellow
	Green
)

func (l Light) String() string {
	switch l {
	case Red:
		return "STOP"
	case Yellow:
		return "SLOW"
	default:
		return "GO"
	}
}

func main() {
	// Day 2 - Variables, Mutability, and Types
	// var name string = "Mark"
	// var age int32 = 123
	// var heightMeters float32 = 3.5
	// var isDevops bool = true

	// fmt.Println(name, age, heightMeters, isDevops)

	// Day 3 - FizzBuzz
	// iter := 20

	// for i := 1; i <= iter; i++ {
	// 	var output string = strconv.Itoa(i)

	// 	if i%3 == 0 {
	// 		output = "Fizz"
	// 	}
	// 	if i%5 == 0 {
	// 		output = "Buzz"
	// 	}
	// 	if i%3 == 0 && i%5 == 0 {
	// 		output = "FizzBuzz"
	// 	}

	// 	fmt.Println(output)
	// }

	// Day 4
	// for i := 2; i <= 20; i++ {
	// 	if isPrime(i) {
	// 		fmt.Println(i)
	// 	}
	// }

	// Day 5
	// var v []int
	// for i := 1; i <= 10; i++ {
	// 	v = append(v, i)
	// }

	// sum := 0
	// for i := range v {
	// 	sum += i
	// }

	// var avg float64 = float64(sum) / float64(len(v))

	// var evens []int
	// for i := range v {
	// 	if i%2 == 0 {
	// 		evens = append(evens, i)
	// 	}
	// }
	// fmt.Printf("sum=%d, avg=%.2f, evens=%v", sum, avg, evens)

	// Day 6
	// s1 := "testing testing"
	// s2 := "testing"
	// fmt.Println(longest(s1, s2))

	// Day 7
	// servers := []Server{
	// 	{Name: "web-01", CPUCores: 4, IsActive: true},
	// 	{Name: "web-02", CPUCores: 4, IsActive: false},
	// 	{Name: "backend-01", CPUCores: 16, IsActive: true},
	// }
	// numActive := 0
	// for _, s := range servers {
	// 	fmt.Println(s.Name, "has", s.CPUCores, "cores")

	// 	if s.IsActive {
	// 		numActive++
	// 	}
	// }
	// fmt.Println(numActive, "servers active")

	// Day 7
	lights := []Light{Red, Yellow, Green}
	for _, l := range lights {
		next := nextTrafficLight(l)

		fmt.Println("Current:", l, "; Next:", next)
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func longest(s1, s2 string) int {
	if len(s1) > len(s2) {
		return len(s1)
	}
	return len(s2)
}

func nextTrafficLight(l Light) Light {
	switch l {
	case Red:
		return Green
	case Green:
		return Yellow
	default:
		return Red
	}
}
