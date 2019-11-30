package main

import "fmt"

func main() {
	//slice是对array的引用或部分引用，值的变化会反映到array和每个slice上
	array := [4]int{10, 20, 30, 40}
	slice := array[0:2]
	//array的cap == len
	fmt.Printf("Before append array = %v, Pointer = %p, len = %d, cap = %d\n", array, &array, len(array), cap(array))
	fmt.Printf("Before append slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	newSlice := append(slice, 50)
	fmt.Printf("After append  slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After append newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After modify slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After modify newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	//array会被append索引修改！！！,array增容后地址没变！！！
	fmt.Printf("After modify array = %v , Pointer = %p \n", array, &array)
	thirdSlice := newSlice[0 : len(newSlice) : cap(newSlice)-1]
	fmt.Printf("%+v %#v %#v %v \n", newSlice, newSlice, thirdSlice, cap(thirdSlice))
}
