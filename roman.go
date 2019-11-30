package main

import "fmt"

type Node struct {
	i int
	v string
}

func intToRoman(num int) string {
	n := num
	m := []Node{
		Node{1, "I"},
		Node{4, "IV"},
		Node{5, "V"},
		Node{9, "IX"},
		Node{10, "X"},
		Node{40, "XL"},
		Node{50, "L"},
		Node{90, "XC"},
		Node{100, "C"},
		Node{400, "CD"},
		Node{500, "D"},
		Node{900, "CM"},
		Node{1000, "M"},
	}
	ml := len(m)
	roma := ""
	for n > 0 {
		for j := ml - 1; j >= 0; j-- {
			if n >= m[j].i {
				n -= m[j].i
				fmt.Println(j, m[j], n)
				roma += m[j].v
				break
			}
		}

	}
	return roma
}
func main() {
	fmt.Println(intToRoman(20))
}
