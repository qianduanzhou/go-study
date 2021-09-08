package main

import "fmt"

//全局变量
var global string = "global"

func foo() (int, string) {
	return 10, "Q1mi"
}

func main() {
	fmt.Println(global)
	//局部变量，同名的全局变量和局部变量可以重复声明
	global := "inside"
	fmt.Println(global)

	//匿名变量 _
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	//声明多个变量
	var (
		a string
		b int
	)
	a = "test"
	b = 1
	fmt.Println(a, b)

	//初始化多个变量
	c, d := 1, 2
	fmt.Println(c, d)

	//常量声明
	const e int = 5
	// e = 6
	fmt.Println("e:", e)

	//多个常量一起声明
	const (
		f = 1
		g = 2
		h //与上一行的值一样
		i
	)
	fmt.Println(f, g, h, i)

	//iota
	const (
		a1 = iota
		a2
		a3
	)
	fmt.Println(a1, a2, a3)

	//多个iota定义一行
	const (
		aa, bb = iota + 1, iota + 2 //1,2
		cc, dd                      //2,3
		ee, ff                      //3,4
	)
	fmt.Println(aa, bb, cc, dd, ee, ff)
}
