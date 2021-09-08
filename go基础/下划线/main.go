package main

//引入的时候使用_代表只引入init函数，使用其他导出的函数会报错
import _ "main/test"

func main() {
	// test.Say()
	arr := [3]int{1, 2, 3}
	//_在代码中代表一个占位符，表示不需要使用该值
	for _, value := range arr {
		println(value)
	}
}
