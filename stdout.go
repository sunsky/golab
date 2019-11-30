package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\r hah %d ", i)
		os.Stdout.Sync()
	}
	fmt.Println("\r\nFinish!")
}
