package main

import "fmt"

func longestPalindrome(s string) string {

	fmt.Println(s)
	longest := []int{0, 1}
	l := len(s)
	if l == 0 {
		return ""
	}
	//if i+1 < l && s[i+1] == s[i] {
	for i, _ := range s {
		_s := i
		_e := i
		if _e+1 < l && s[_e+1] == s[i] {
			for _e+1 < l && s[_e+1] == s[i] {
				_e++
			}
			_e--
		}
		fmt.Println("repeat ", i+1)
		//		} else {
		//			_e = i
		//}
	I:
		for {
			if _e >= l || _s < 0 || s[_e] != s[_s] {
				if _e-_s > longest[1]-longest[0] {
					fmt.Println(longest)
					longest = []int{_s + 1, _e}
				}
				fmt.Println("beak at ", i, _s, _e)
				break I
			}
			_e++
			_s--
		}

		//		fmt.Println("alph ", string(s[i]))
	}

	return s[longest[0]:longest[1]]
}
func main() {
	fmt.Println(
		"abc"[0:1],
		longestPalindrome("vimimgo"),
		longestPalindrome("vimiximiy"),
		longestPalindrome("ooxxoo-ooxx9oo"),
		longestPalindrome("go"),
		longestPalindrome(""),
		longestPalindrome("bb"),
		longestPalindrome("bbb"),
	)

}
