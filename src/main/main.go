package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func say() {
	fmt.Println("say")
}

func hasReturn() int {
	return 100
}

func add(x int, y int) int {
	return x + y
}

//https://tour.golang.org/basics/5 短路
func add3(x, y, z int) int {
	return x + y + z
}

//https://tour.golang.org/basics/6 多结果
func swapMulti(p0, p1 string) (string, string) {
	return p1, p0
}

//https://tour.golang.org/welcome/4
func main() {
	fmt.Println("hello world!")
	fmt.Println("The time is", time.Now())
	fmt.Println("rand is", rand.Intn(10))
	fmt.Println("math.Squrt is", math.Sqrt(4))
	fmt.Println("math.pi is not access,please use math.Pi", math.Pi)
	//https://tour.golang.org/basics/4
	fmt.Println("print some,", hasReturn(), add(1000, 1234))
	say()
	fmt.Println("短路", add3(1, 2, 3))
	a, b := swapMulti("hello", "world")
	fmt.Println("多结果", a, b)
}
