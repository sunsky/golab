package main

import (
	"fmt"
	"sort"
)

//按插入的顺序
type RawSortedMap struct {
	m    map[string]interface{}
	keys []interface{}
}

func NewRawSortedMap(cap int) RawSortedMap {
	m := make(map[string]interface{}, cap)
	keys := make([]interface{}, cap)
	n := RawSortedMap{m, keys}
	return n
}

func (r *RawSortedMap) Set(k string, v interface{}) {
	//通过k 快速查找
	r.m[k] = len(r.m)
	//通过keys数组快速遍历
	r.keys = append(r.keys, v)
}
func main() {

	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"

	// To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	//按数字大小的顺序
	sort.Ints(keys)
	//按字母的顺序
	//sort.Strings(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
