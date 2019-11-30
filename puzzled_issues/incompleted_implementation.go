package main

import "fmt"

type Talkable interface {
	TalkEnglish(string)
	TalkChinese(string)
}

type Student1 struct {
	Talkable
	Name string
	Age  int
}

func (s *Student1) TalkChinese(s1 string) {
	fmt.Printf("我是 %s, 今年%d岁,%s", s.Name, s.Age, s1)
}
func main() {
	a := Student1{Name: "aaa", Age: 12}
	var b Talkable = &a
	talk, ok := b.(Talkable)
	fmt.Printf("%#+v %#+v %v \n", b, talk, ok) //main.Student1{Talkable:main.Talkable(nil), Name:"aaa", Age:12} main.Student1{Talkable:main.Talkable(nil), Name:"aaa", Age:12} true
	//a.TalkEnglish("nice to meet you\n") //panic: runtime error: invalid memory address or nil pointer dereference
	a.TalkChinese("哈哈\n") //我是 aaa, 今年12岁,哈哈

	/*
		interface 时 golang 编程中使用得非常频繁的特性，我们需要明白它的底层结构，以及一些编译和运行时的特殊之处，能帮我们避免一些不必要的麻烦：

		interface 很类似void*，但在值类型的变量和 interface 类型变量相互赋值时，会发生数据的复制。
		将某个类型的指针赋值给 interface，interface 的值永远不可能是 nil；
		interface 可以嵌入结构体，帮类型快速实现接口，但是注意如果调用未实现的方法则会 panic；
	*/
}
