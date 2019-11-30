package main

import "fmt"

//import "reflect"

func isValid(s string) bool {
	//mEnd := map[rune]rune{')': '(', ']': '[', '}': '{'}
	m := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	tmp := []rune{}
	for _, v := range s {
		//fmt.Println(reflect.TypeOf(s[i]), reflect.TypeOf(v))
		if _, ok := m[v]; ok {
			tmp = append(tmp, v)
		} else {
			ti := len(tmp) - 1
			if ti < 0 {
				return false
			}
			if tt, ok := m[tmp[ti]]; !ok || tt != v {
				return false
			}
			tmp = tmp[:ti]
		}
	}
	fmt.Printf("%#v \n", tmp)
	return len(tmp) == 0
}
func main() {
	fmt.Println("vim-go", isValid("()") == true, isValid("()[]{}") == true, isValid("{[]}") == true, isValid("]") == false, isValid("[") == false)
}
