package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	s := Student{
		Name: "aaa",
	}
	var b interface{} = s //复制值
	var c = b.(Student)
	c.Name = "bbb"
	fmt.Printf("%#p %#p %#p %+#v \n", &s, &b, &c, b.(Student).Name) //c0000101f0 c000010200 c000010220 "aaa"
	var d interface{} = &s                                          //复制指针
	var e = d.(*Student)
	e.Name = "bbb"
	fmt.Printf("%#p %#p %#p %+#v  \n", &s, d, e, d.(*Student).Name) //c0000101f0 c0000101f0 c0000101f0 "bbb"

}
