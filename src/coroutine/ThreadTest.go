package coroutine

import (
	"fmt"
	"time"
)

func say() {
	fmt.Println("hi")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("i know")
}

func send2chanel(data string, c chan string) {
	c <- data
}

//https://tour.golang.org/concurrency/2
func main() {

	for i := 0; i < 10; i++ {
		go say()
	}
	c := make(chan string)
	go send2chanel("a", c)
	go send2chanel("b", c)
	a, b := <-c, <-c
	fmt.Println(a, b)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3
	fmt.Println(<-ch, <-ch)

	//select https://tour.golang.org/concurrency/5

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("tail")
}
