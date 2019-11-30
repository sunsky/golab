package main

import "testing"

var count int = 100_000

func addition(m map[int]int) map[int]int {
	for i := 0; i < count; i++ {
		m[i] = i
	}
	return m
}
func BenchmarkGrows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		addition(m)
	}
}
func BenchmarkNoGrows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, count)
		addition(m)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		n := make(map[int]int, count)
		addition(n)
		//m := make(map[int]int)
		//addition(m)
	}
}
