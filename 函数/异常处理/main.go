package main

import "fmt"

//panic
func panicTest() {
	defer func() { //recover只能和defer配合使用
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
		fmt.Println("panicTest2")
	}()
	panic("panic error!")
	fmt.Println("panicTest1") //panic会终止之后执行的语句，除非用recover捕获
}

func test() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
}
func test2() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()

	panic("test panic")
}

func test3() {
	test4()
}
func test4() int {
	fmt.Println("0")
	return 2
}
func deferTest() {
	defer fmt.Println(test4())

	fmt.Println("1")
}
func main() {
	// panicTest()

	//向已关闭的通道发送数据会引发panic
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// var ch chan int = make(chan int, 10)
	// close(ch)
	// ch <- 1

	//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
	// test()

	//捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
	// test2()

	deferTest()
}
