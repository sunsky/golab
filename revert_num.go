package main

import "fmt"
import "math"

func reverse(x int) int {
	fmt.Println("input ", x)
	a := []int{}
	t := x
	for t != 0 {
		a = append(a, t%10)
		t = t / 10
	}
	res := 0
	for i, v := range a {
		for j := len(a) - i - 1; j > 0; j-- {
			v *= 10
		}
		fmt.Println(v)

		res += v
		if math.Abs(float64(res)) > math.Pow(2, 31) {
			return 0
		}
	}
	//fmt.Println(a)
	return res
}
func main() {
	fmt.Println(reverse(123450))
	fmt.Println(reverse(-123450))
	fmt.Println(reverse(1534236469))
}
