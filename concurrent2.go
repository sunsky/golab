package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("index: %d,goroutine Num: %d\n", 0, runtime.NumGoroutine())
	fmt.Println("done!")
}
