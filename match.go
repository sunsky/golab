package main

import "fmt"

func isMatch(s string, p string) bool {
	//为空
	l := len(s)
	lp := len(p)
	if l == 0 && lp == 0 {
		return true
	}
	if len(lp) == 0 {
		return false
	}
	j := 0
	dot := 0
	for i, v := range s {
		if '*' == p[j] {
			if dot == 1 {
				z:=j+1
				for z<lp {
					if p[z]=='*' || p[z]=='.'{
						break
				}
			} else {
				if v != p[j-1] {
					return false
				}
			}
		}

		if '.' == p[j] {
			dot = 1
		} else if p[j] == v {
			dot = 0
		} else if j+1 < lp && p[j+1] == '*' {
			dot = 0
			j++
		} else {
			return false
		}
		j++

	}
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
