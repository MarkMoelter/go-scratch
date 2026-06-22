package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

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

type Rectangle struct{ Width, Height float64 }

func NewSquare(side float64) Rectangle { return Rectangle{Width: side, Height: side} }

func (r *Rectangle) Scale(factor float64) { r.Width *= factor; r.Height *= factor }
func (r Rectangle) Area() float64         { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64    { return 2 * (r.Width + r.Height) }
func (r Rectangle) IsSquare() bool        { return r.Width == r.Height }

type Describable interface{ Describe() string }

type Database struct{ Name, Engine string }

func (s Server) Describe() string   { return "Server: " + s.Name }
func (d Database) Describe() string { return fmt.Sprintf("Database: %s (%s)", d.Name, d.Engine) }

type Shape interface{ Area() float64 }
type Circle struct{ Radius float64 }

func (c Circle) Area() float64 { return math.Pow(c.Radius, 2) * math.Pi }

func main() {
	// Day 1 - Hello World & Toolchain
	// fmt.Println("Hello, world!")

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

	// Day 4 - Functions
	// for i := 2; i <= 20; i++ {
	// 	if isPrime(i) {
	// 		fmt.Println(i)
	// 	}
	// }

	// Day 5 - Collections — Slices, Strings
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

	// Day 6 - Pointers
	// s1 := "testing testing"
	// s2 := "testing"
	// fmt.Println(longest(s1, s2))

	// Day 7 - Structs
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

	// Day 8 - Iota
	// lights := []Light{Red, Yellow, Green}
	// for _, l := range lights {
	// 	next := nextTrafficLight(l)
	// 	fmt.Println("Current:", l, "; Next:", next)
	// }

	// Day 9 - Methods
	// s := NewSquare(4.0)
	// s.Scale(2.0)
	// fmt.Println("area =", s.Area(), "; perimeter = ", s.Perimeter(), "; is a square = ", s.IsSquare())
	// r := Rectangle{Width: 4.0, Height: 3.0}
	// r.Scale(2.0)
	// fmt.Println("area =", r.Area(), "; perimeter = ", r.Perimeter(), "; is a square = ", r.IsSquare())

	// Day 10 - Maps & Iterators
	// hosts := map[string]uint32{
	// 	"web-01":  4,
	// 	"web-02":  4,
	// 	"db-01":   8,
	// 	"bcdr-01": 16,
	// }
	// var maxName string
	// var maxCores uint32
	// for name, cores := range hosts {
	// 	if cores > maxCores {
	// 		maxCores = cores
	// 		maxName = name
	// 	}
	// }
	// fmt.Println(maxName, maxCores)

	// Day 11 - Error Handling
	// result, err := divide(10, 2)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// } else {
	// 	fmt.Println("result =", result)
	// }
	// _, err = divide(10, 0)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// pInt, err := parsePositiveInt("6")
	// if err != nil {
	// 	fmt.Println("error:", err)
	// } else {
	// 	fmt.Println("result =", pInt)
	// }

	// Day 12 - Interfaces
	s := Server{Name: "web-01", CPUCores: 4, IsActive: true}
	d := Database{Name: "db-01", Engine: "SQL Server"}

	items := []Describable{s, d}
	printAll(items)

	r := Rectangle{Height: 5.0, Width: 3.0}
	c := Circle{Radius: 10}

	shapes := []Shape{r, c}
	for _, item := range shapes {
		fmt.Println(item.Area())
	}
}

// Day 4
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

// Day 6
func longest(s1, s2 string) int {
	if len(s1) > len(s2) {
		return len(s1)
	}
	return len(s2)
}

// Day 8
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

// Day 11
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func parsePositiveInt(s string) (uint32, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("not a number")
	}
	if n < 0 {
		return 0, errors.New("negative")
	}
	return uint32(n), nil

}

func printAll(items []Describable) {
	for _, item := range items {
		fmt.Println(item.Describe())
	}
}
