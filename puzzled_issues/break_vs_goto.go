package main

import (
	"fmt"
)

var t int = 100

func main() {
	if t > 10 {
		goto A
	} else {
		goto B
	}
A:
	{
		fmt.Println("A")
	}

B:
	{
		fmt.Println("B")
	}
	input := make(chan int)
	go func() {
		input <- 2
	}()
LOOP:
	for {
		select {
		case <-input:
			println("input")
			break LOOP
			//goto LOOP //fatal error: all goroutines are asleep - deadlock!
		}
	}

	fmt.Println("end")
}
