package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		//可以看到，如果用 range 的方式去遍历一个切片，拿到的 Value 其实是切片里面的值拷贝。所以每次打印 Value 的地址都不变
		fmt.Printf("value = %d , value-addr = %X , slice-addr = %X\n", value, &value, &slice[index])
		//&slice[index]才是slice成员的真实地址（16进制）, 可见每个地址间隔是8 Byte(字节)
	}
}
