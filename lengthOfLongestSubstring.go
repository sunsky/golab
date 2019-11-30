package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	fmt.Println("input ", s)
	l := len(s)
	m := map[byte]int{}
	max := 0
	count := 0
	for i := 0; i < l; i++ {
		if v, ok := m[s[i]]; ok {
			for k, _v := range m {
				if _v <= v {
					println("deleting hit repeated index/value/count ", _v, string(k), count, "\n")
					delete(m, k)
					count--
				}

			}

		}
		count++
		if max < count {
			max = count
		}
		m[s[i]] = i
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("a"))
}
