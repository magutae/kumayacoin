package main

import "fmt"

func countToTen(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Printf("sending %d\n", i)
	}
	close(c)
}

func receive(c <-chan int) {
	for {
		a, ok := <-c
		if !ok {
			fmt.Println("we are done.")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

func main() {
	c := make(chan int)
	go countToTen(c)
	receive(c)
}
