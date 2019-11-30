package main

import (
	"fmt"
	"time"
)

type Notification struct {
	Message string
}

func main() {
	notify := make(chan Notification)
	go func() {
		fmt.Println("sleep ... ")
		time.Sleep(3 * time.Second)
		close(notify)
	}()
	select {
	//case notification, ok:= <- notify:
	//	fmt.Println(notification, ok)
	case <-notify:
		fmt.Println("closed .")
	}

	fmt.Println("done")
}
