package main

import "fmt"

type Node struct {
	i int
	v string
}

func romanToInt(s string) int {
	n := 0
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
	sl := len(s)
	for i := 0; i < sl; {
		for j := ml - 1; j >= 0; j-- {
			l := len(m[j].v)
			if i+l <= sl && m[j].v == s[i:i+l] {
				i += l
				n += m[j].i
				break
			}
		}

	}
	return n
}
func main() {
	fmt.Println(romanToInt("MCMXCIV") == 1994)
	fmt.Println(romanToInt("III") == 3)
}
