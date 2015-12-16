package main

import (
	"fmt"
)

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3 + <-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func(id string){
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(id, i)
		}
		close(out)
	}(s)
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func(){
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

// go run -race main.go
// vs
// go run main.go
