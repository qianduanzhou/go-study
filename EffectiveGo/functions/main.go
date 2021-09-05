package main

import "fmt"

//返回值默认赋值0
func test() (i int) {
	// i = 5
	return
}

//defer执行顺序
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() string {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
	return "end"
}

func main() {
	i := test()
	fmt.Println(i)
	//延迟函数按后进先出顺序执行，因此当函数返回时，这段代码将导致打印4 3 2 10
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println(b())
}
