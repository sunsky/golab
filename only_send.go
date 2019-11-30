package main

import "time"
import "fmt"
import "runtime"

func gen(done chan int, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				break
			}
		}
	}()
	return out
}

func main() {
	done := make(chan int)
	out := gen(done, 2, 3)
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	defer close(done)

	// Set up the pipeline.

	for n := range out {
		fmt.Println(n)              // 2
		time.Sleep(2 * time.Second) // done thing, 可能异常中断接收
		if true {                   // if err != nil
			break
		}
	}

}
