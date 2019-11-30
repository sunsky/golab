package main

import (
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.ParseInt(os.Getenv("A"), 10, 64)
	b, _ := strconv.ParseInt(os.Getenv("B"), 10, 64)
	a ^= b
	b ^= a
	a ^= b
	println("A :", a, "B :", b, 1.33)
}
