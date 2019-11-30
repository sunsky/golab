package main

import (
	"fmt"
	"strconv"
)

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	nl := len(needle)
	if nl == 0 {
		return 0
	}
	hl := len(haystack)
	if hl == 0 {
		return -1
	}
	nm := map[byte]int{}
	for i := 0; i < nl; i++ {
		nm[needle[i]] = i
	}
	fmt.Printf("%+#v %v %v ", nm, haystack, needle)
	for i := 0; i < hl; {
		j := 0
		tmp := i
		for ; j < nl && tmp < hl; j++ {
			if haystack[tmp] != needle[j] {
				break
			}
			tmp++
		}
		fmt.Printf("i %v, j %v  ", i, j)
		if j == nl {
			return i
		} else if i+nl < hl {
			hit, exists := nm[haystack[i+nl]]
			fmt.Printf("alpha %v hit %v exst %v ", string(haystack[i+nl]), hit, exists)
			if exists {
				i += nl - hit
			} else {
				i += nl + 1
			}
		} else {
			return -1
		}
	}
	return -1
}
func main() {
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("abcdeabc", "abcab"))
	fmt.Println(strStr("abcdeabc", "abcabe"))
	fmt.Println(strStr("abcebc", "abc"))
	fmt.Println(strStr("nnabcd e aebc", "abc"))
	fmt.Println(strStr("mississippi", "issi"))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Printf("\n %#+v, %#+v ", strconv.Itoa(9), string(9))
}
