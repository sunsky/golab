package main

import "strconv"

type S struct {
	M *int
}

func main() {

	strconv.ParseInt()
	var x S
	var i int
	ref(&i, &x)
}
func ref(y *int, z *S) {
	z.M = y
}
