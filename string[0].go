package main

import (
	"fmt"
	"reflect"
)

func main() {
	asci := "abc"
	ch := "我是中国人"
	mix := "i am 中国人"
	fmt.Println(
		reflect.TypeOf(asci),
		reflect.TypeOf(ch),
		reflect.TypeOf(mix),
	)
	fmt.Println(
		reflect.TypeOf(asci[0]),
		reflect.TypeOf(ch[0]),
		reflect.TypeOf(mix[0]),
	)
	for i, v := range asci {
		fmt.Println(reflect.TypeOf(asci[i]), reflect.TypeOf(v), string(v), string(asci[i]), string(asci[:3]))
		break
	}
	for i, v := range ch {
		fmt.Println(reflect.TypeOf(asci[i]), reflect.TypeOf(v), string(v), string(ch[i]), string(ch[:3]))
		break
	}
	for i, v := range mix {
		fmt.Println(reflect.TypeOf(asci[i]), reflect.TypeOf(v), string(v), string(mix[i]), string(mix[:3]))
		break
	}
	// 总结 ：
	// for会把 string[i]进行强制转换为 rune,  如果我们只需要byte，就多此一举了
	// 如果我们的string是非英文的语言，那是默认按rune处理是非常稳妥的方式， 不会有乱码。

	s := "截取中文"
	//试试这样能不能截取?
	//fmt.Println(s[:3])
	fmt.Println(string([]rune(s)[:3]))

}
