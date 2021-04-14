package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//https://tour.golang.org/welcome/4
func main() {
	fmt.Println("hello world!")
	fmt.Println("The time is", time.Now())
	fmt.Println("rand is", rand.Intn(10))
	fmt.Println("math.Squrt is", math.Sqrt(4))

}
