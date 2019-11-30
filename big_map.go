package main

import "runtime/debug"

var count = 10_000_000
var dict = make(map[int]int, count)

func addition() {
	for i := 0; i < count; i++ {
		dict[i] = i
	}
}
func clear() {
	for k := range dict {
		delete(dict, k)
	}
}
func main() {
	addition()
	println("delete all map item")
	clear()
	debug.FreeOSMemory()
	println("delete map")
	dict = nil
	debug.FreeOSMemory()

}
