package main

import "fmt"
import "reflect"

func main() {
	first := "fisrt"
	fmt.Printf("vim-go %v %v  %v %v %v \n", reflect.TypeOf(first[0]), reflect.TypeOf("a我"[0]), reflect.TypeOf([]rune("b你")[0]), string("你q我"[0]), string([]rune("我h你")[0]))
}
