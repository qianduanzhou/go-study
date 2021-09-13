package main

import (
	"fmt"
	"math"
)

/* 函数定义 */
func define(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
	n := x + y
	return n, fmt.Sprintf(s, n)
}

/* 函数作为参数传递 */
func test(fn func() int) int {
	return fn()
}

// 定义函数类型。
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

/* 引用传递 */
func swap(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}

/*可选参数*/
func testFunc(a int, arr ...int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func testFuncSlice(a int, arr ...int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

/*直接返回*/
func add(a, b int) (sum, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

/*普通返回*/
func add2(a, b int) (int, int) {
	var sum int = a + b
	var avg int = (a + b) / 2
	return sum, avg
}
func main() {
	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(s1, s2)

	/* 引用传递 */
	var a, b int = 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	/*可选参数*/
	testFunc(1, 2, 3, 4)
	//slice可选参数
	var slice []int = []int{1, 2, 3}
	testFuncSlice(1, slice...)

	/*直接返回*/
	sum, avg := add(2, 2)
	fmt.Println(sum, avg)
	//只使用其中一个
	sum2, _ := add2(2, 2)
	fmt.Println(sum2)

	/* 匿名函数 */
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))

	/* 赋值给变量做为结构字段或者在channel里传送 */
	// --- function variable ---
	fn := func() { println("Hello, World!") }
	fn()

	// --- function collection ---
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	// --- function as field ---
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello, World!" },
	}
	println(d.fn())

	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
}
