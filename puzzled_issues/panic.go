package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("defer in main", e)
		}
	}()
	slice := make([]byte, 2, 5)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("defer in main go", e)
			}
		}()
		fmt.Println(slice[:8])

	}()
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("defer in goroutine of main", e)
			}
		}()
		inGoroutine(slice)
	}()
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("defer in inNoRecoverGoroutine of main", e)
			}
		}()
		inNoRecoverGoroutine(slice)
	}()
	go func() {
		inRecoverGoroutine(slice)
	}()
	go func() { // 在goroutine外recover无效
		panic("panic in goroutine")
	}()
	outOfMain(slice)
	fmt.Println("main is alive")
	time.Sleep(time.Hour)
}
func outOfMain(slice []byte) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("defer in outOfMain", e)
			}
		}()
		fmt.Println(slice[:8])

	}()
}
func inGoroutine(slice []byte) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("defer in inGoroutine", e)
		}
	}()
	fmt.Println(slice[:8])
}
func inRecoverGoroutine(slice []byte) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("defer in inRecoverGoroutine", e)
		}
	}()
	fmt.Println(slice[:8])
}

func inNoRecoverGoroutine(slice []byte) {
	fmt.Println(slice[:8])
}
