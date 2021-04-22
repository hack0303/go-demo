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

type Vertex struct {
	x int
	y int
}

func printArray(ary []int) {
	for i := 0; i < len(ary); i++ {
		fmt.Println(ary[i])
	}
}

func testFuncParam(x, y int, f func(int, int) int) (v int) {
	v = f(x, y)
	return
}

func returnFunc() func(int) int {
	origin := 10
	return func(x int) int {
		return origin + x
	}
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

	//point
	//https://tour.golang.org/moretypes/1
	t01, t02 := 1024, 2048
	p := &t01
	*p = *p + t02
	*p = *p / 3
	fmt.Println("*p is ", *p)

	fmt.Println(Vertex{1, 2})

	vertex1 := Vertex{2, 4}
	fmt.Println(vertex1.x)

	//参考c，
	vertex2 := Vertex{}
	vertex3 := Vertex{x: 10}
	fmt.Println(vertex2.x)
	fmt.Println(vertex3.x)
	vertex3_ptr := &vertex3
	fmt.Println(vertex3_ptr.x)

	var array1 [10]int
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	fmt.Println(array1)
	array2 := []int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	var array3 []int = array2[0:1]
	fmt.Println(array3)

	fmt.Println([]int{1, 2, 3, 4})
	fmt.Println([]bool{true, false, true, false})
	fmt.Println([]struct {
		x int
		y string
	}{{1, "y1"}, {2, "y2"}})

	fmt.Println(array2[:], array2[1:], array2[:1], array2[1:4])

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	var s0 []int
	fmt.Println(s0, len(s0), cap(s0))
	if s0 == nil {
		fmt.Println("nil!")
	}

	pow := make([]int, 10)
	for i := 0; i < len(pow); i++ {
		pow[i] = i
	}
	//or
	for i := range pow {
		pow[i] = i
	}

	for _, value := range pow {
		fmt.Println(value)
	}

	//map
	aMap := make(map[string]Vertex)

	aMap["heihei"] = Vertex{1, 2}
	fmt.Println(aMap["heihei"])

	//注意每个k,v(ˇ?ˇ) 项后面的 逗号
	var aMap1 = map[string]Vertex{
		"a": {1, 2},
	}
	for _, value := range aMap1 {
		fmt.Println(value.x, value.y)
	}

	aMap2 := make(map[string]int)
	aMap2["a"] = 2
	aMap2["b"] = 3
	fmt.Println("fmt.Println(aMap2[\"b\"])", aMap2["b"])
	delete(aMap2, "b")
	fmt.Println("fmt.Println(aMap2[\"b\"])", aMap2["b"])
	elem, ok := aMap2["b"]
	fmt.Println("fmt.Println(aMap2[\"b\"])", ok, elem)
	//H
	tfunc := func(x int, y int) int {
		return x + y
	}
	fmt.Println("testFuncParam(1,2,tfunc)", testFuncParam(1, 2, tfunc))

	fmt.Println("returnFunc", returnFunc()(20))

	//error
	//https://tour.golang.org/methods/19
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
