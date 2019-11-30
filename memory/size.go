package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var intValue int
	var uint8Value uint8
	var uint16Value uint16
	var uint32Value uint32
	var uint64Value uint64
	var int8Value int8
	var int16Value int16
	var int32Value int32
	var int64Value int64

	var float32Value float32
	var float64Value float64
	var boolValue bool
	var ptrValue uintptr
	var complex64Value complex64
	var complex128Value complex128
	var strValue string
	var byteValue byte
	var runeValue rune
	structValue := struct {
		FieldA float32
		FieldB string
	}{0, ""}
	mapValue := map[int]int{}
	var sliceValue []int
	var intPtrValue *int
	var chanValue chan int
	var funcValue func()

	fmt.Println("funcValue = Size:", unsafe.Sizeof(funcValue))     //size:  8
	fmt.Println("chanValue = Size:", unsafe.Sizeof(chanValue))     //size:  8
	fmt.Println("intPtrValue = Size:", unsafe.Sizeof(intPtrValue)) //size:  8
	fmt.Println("sliceValue = Size:", unsafe.Sizeof(sliceValue))   //size:  24
	//type slice struct {
	//	array unsafe.Pointer
	//	len   int
	//	cap   int
	//}
	fmt.Println("mapValue = Size:", unsafe.Sizeof(mapValue))                //size:  8
	fmt.Println("structValue = Size:", unsafe.Sizeof(structValue))          //size:  24
	fmt.Println("strValue = Size:", unsafe.Sizeof(strValue), len(strValue)) //intValue = Size: 16, string is the set of all strings of 8-bit bytes, conventionally but not
	// necessarily representing UTF-8-encoded text. A string may be empty, but
	// not nil. Values of string type are immutable.
	//type stringStruct struct {
	//	str unsafe.Pointer
	//	len int
	//}
	fmt.Println("byteValue = Size:", unsafe.Sizeof(byteValue))             //intValue = Size: 1
	fmt.Println("runeValue = Size:", unsafe.Sizeof(runeValue))             //intValue = Size: 4
	fmt.Println("boolValue = Size:", unsafe.Sizeof(boolValue))             //intValue = Size: 1
	fmt.Println("ptrValue = Size:", unsafe.Sizeof(ptrValue))               //intValue = Size: 8
	fmt.Println("complex64Value = Size:", unsafe.Sizeof(complex64Value))   //intValue = Size: 8
	fmt.Println("complex128Value = Size:", unsafe.Sizeof(complex128Value)) //intValue = Size: 16
	fmt.Println("intValue = Size:", unsafe.Sizeof(intValue))               //intValue = Size: 8
	fmt.Println("uint8Value = Size:", unsafe.Sizeof(uint8Value))           //uint8Value = Size: 1
	fmt.Println("uint16Value = Size:", unsafe.Sizeof(uint16Value))         //uint16Value = Size: 2
	fmt.Println("uint32Value = Size:", unsafe.Sizeof(uint32Value))         //uint32Value = Size: 4
	fmt.Println("uint64Value = Size:", unsafe.Sizeof(uint64Value))         // uint64Value = Size: 8

	fmt.Println("int8Value = Size:", unsafe.Sizeof(int8Value))   //int8Value = Size: 1
	fmt.Println("int16Value = Size:", unsafe.Sizeof(int16Value)) //int16Value = Size: 2
	fmt.Println("int32Value = Size:", unsafe.Sizeof(int32Value)) //int32Value = Size: 4
	fmt.Println("int64Value = Size:", unsafe.Sizeof(int64Value)) //int64Value = Size: 8

	fmt.Println("float32Value = Size:", unsafe.Sizeof(float32Value)) //float32Value = Size: 4
	fmt.Println("float64Value = Size:", unsafe.Sizeof(float64Value)) //float64Value = Size: 8

}
