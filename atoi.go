package main

import "fmt"

func myAtoi(str string) int {
	//l := len(str)
	isN := 0
	for i, v := range str {
		fmt.Println(i, v, int(v))
		if v == ' ' {
			continue
		}
		switch v ( 
		)
		if v == '-' || v == '+' || v >= '0' || v <= '9' {
			isN = 1
		} else {
			return 0
		}
		if isN > 0 && (v >= '0' || v <= '9') {

		}
	}
	//for i := 0; i < l; i++ {
	//	fmt.Println(int())
	//}
	return 0
}

func main() {
	fmt.Println("vim-go", myAtoi("-12340"))
	fmt.Println("vim-go", myAtoi("12340"))
	fmt.Println("vim-go", myAtoi(" 12340"))
	fmt.Println("vim-go", myAtoi(" -12340"))
	fmt.Println("vim-go", myAtoi(" -12.340"))
	fmt.Println("vim-go", myAtoi(" -12ff"))
	fmt.Println("vim-go", myAtoi("xx 12ff"))
}
