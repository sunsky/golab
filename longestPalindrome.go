package main

import "fmt"

func longestPalindrome(s string) string {
	fmt.Println(s)
	longest := []int{0, 1}
	m := map[rune]int{}
	l := len(s)
	for i, c := range s {
		if v, ok := m[c]; ok {
			_s := v
			_e := i
		I:
			for _e < l && _s > -1 {
				if _e-_s <= 2 && s[_e] != s[_s] {
					if _e-_s > longest[1]-longest[0] {
						longest = []int{_s + 1, _e - 1}
					}

					break I
				}
				_e++
				_s--
			}
		}
		m[c] = i
	}
	fmt.Println(longest)

	return s[longest[0]:longest[1]]
}
func main() {
	fmt.Println(
		longestPalindrome("vimim-go"),
		longestPalindrome("vimimxmim"),
		longestPalindrome("-go"),
	)

}
