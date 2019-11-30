package main

import (
	"bytes"
	"fmt"
	"io"
)

func GetReader(id int64) io.Reader {
	var r *bytes.Buffer = nil
	if id > 0 && id < 10000 {
		r.WriteByte(byte(id))
	}
	return r
}
func main() {
	r := GetReader(-2)
	if r == nil {
		fmt.Println("bad reader")
	} else {
		fmt.Printf("valid reader: %+#v \n", r) //valid reader: (*bytes.Buffer)(nil)
	}
	type Student struct {
		Name string
	}
	var b interface{} = nil
	fmt.Println(b == nil) //true    无类型的nil值就是nil
	var a *Student = nil
	var c interface{} = a
	fmt.Println(c == nil) //false   有类型的nil值不是nil
	//当一个带类型的指针赋值给interface 类型时，无论此指针是否为nil，赋值过的 interface 都不为 nil。无论是 iface 还是 eface，都有两个指针，指向数据的是通用指针，还有一个指针用于指定数据类型或方法集；当我们将一个 nil 指针赋值给 interface 时，实际是对 interface 的这两个指针分别赋值，虽言数据指针 data 为 nil，但是类型指针_type 或 tab 并不是 nil，他将指向你的空指针的类型，因此赋值的结果 interface 肯定不是 nil 啦！
}
