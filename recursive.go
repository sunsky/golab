package main

import (
	"fmt"
)

func letterCasePermutation(S string) []string {
	//s := []byte(S)
	//println("S ", S, cap(s), len(S))
	result := []string{}
	prefix := []string{}
	if len(S) == 0 {
		return result
	} else if S[0] < 'A' {
		prefix = append(prefix, string(S[0]))
	} else if S[0] >= 'a' {
		prefix = append(prefix, string(S[0]), string(S[0]-32))
	} else {
		prefix = append(prefix, string(S[0]), string(S[0]+32))
	}
	suffix := letterCasePermutation(S[1:])
	//fmt.Println("suff", suffix, "prefix", prefix)
	if len(suffix) == 0 {
		suffix = []string{""}
	}
	for _, pre := range prefix {
		for _, suf := range suffix {
			result = append(result, pre+suf)
		}
	}
	return result
}
func print(s string) string {
	if len(s) == 0 {
		return ""
	}
	println(s)
	print(s[1:])
	return s
}

func main() {
	println(map[int]int{1: 1, 1: 2})
	S := "abcd"
	//print(S)
	fmt.Printf("%#v\n", letterCasePermutation(S))
}
