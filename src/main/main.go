package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
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

func swapMulti2(p0, p1 string) (x, y string) {
	x = p1
	y = p0
	return
}

//https://tour.golang.org/basics/8
var c, python, java bool

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
	c, d := swapMulti2("hello", "world")
	fmt.Println("多结果，命名返回", c, d)
	var vari int
	fmt.Println("var is", vari, c, python, java)
	var var1, var2, var3 = 'a', true, 10
	fmt.Println("var1-3", var1, var2, var3)
	k := 3
	fmt.Println("k is", k)
	var floatV float32 = float32(k)
	fmt.Println("floatV is", floatV)
	const notChangeable = "哈哈"
	fmt.Println("const is ", notChangeable)
	sum := 0
	for x := 0; x <= 10; x++ {
		sum += x
	}
	fmt.Println("value is", sum)
	//do while
	for sum < 1000 {
		sum += sum
	}
	// while
	for sum < 1000 {
		sum += sum
	}
	//eq while true
	//for{}

	if sum > 0 {
		//nop
	}
	if sum := 0; sum >= 0 {
		//
	} else {
		//
	}

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//https://tour.golang.org/flowcontrol/12
	defer fmt.Println("world")

	fmt.Println("hello")

	for x := 0; x < 10; x++ {
		defer fmt.Println(x)
	}
}
