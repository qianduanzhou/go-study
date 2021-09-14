package main

import "fmt"

/* 闭包 */
func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

/* 外部引用函数参数局部变量 */
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

/*递归数字阶乘*/
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

/*斐波那契数列(Fibonacci)*/
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
	/* 闭包 */
	f := test()
	f()
	/* 外部引用函数参数局部变量 */
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

	// 返回2个函数类型的返回值
	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))

	/*递归数字阶乘*/
	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	/*斐波那契数列(Fibonacci)*/
	var j int
	for j = 0; i < 10; j++ {
		fmt.Printf("%d\n", fibonaci(j))
	}
}
