package main

import "fmt"

func main() {
	s := "xx"
	var a *string = &s
	var b *string
	b = a

	fmt.Println(a, *a, b, *b)
	b = nil
	fmt.Println(a, *a, b, nil)
	c := "yy"
	b = &c
	fmt.Println(a, *a, b, *b)

	f := func(s *string) {
		*s += "--"
	}
	f(b)
	fmt.Println(a, *a, b, *b)

}
