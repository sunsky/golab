package main

import "fmt"

type ErrDetailed struct {
	code    int
	message string
}

func (e ErrDetailed) Error() string {
	return fmt.Sprintf("err detailed is %s , %s \n", e.code, e.message)
}
func trigger(x int) error {
	//func trigger(x int) *ErrDetailed {
	if x != 1 {
		return &ErrDetailed{code: 100, message: "!=1"}
	}
	return nil
}

func main() {
	var e error
	//var e *ErrDetailed
	e = trigger(0)
	if e != nil {
		fmt.Printf("!=nil, e: %#v\n", e)
	}
	e = trigger(1)
	if e != nil {
		fmt.Printf("!=nil at 2th, e: %#v \n", e)
	}

	//fmt.Println("nil==nil", nil == nil)
}
