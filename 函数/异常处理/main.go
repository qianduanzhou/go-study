package main

import (
	"errors"
	"fmt"
)

//panic
func panicTest() {
	defer func() { //recover只能和defer配合使用
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
		fmt.Println("panicTest2") //recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
	}()
	panic("panic error!")
	fmt.Println("panicTest1") //recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
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

func except() {
	fmt.Println(recover())
}

func test3(x, y int) {
	var z int
	defer func() {
		err := recover()
		fmt.Println("err:", err)

		if err != nil {
			z = 0
		}
	}()
	z = x / y
	fmt.Println("z:", z)
	panic("test3")
}

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
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

	//使用延迟匿名函数或下面这样都是有效的。
	// defer except()
	// panic("err")

	//如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执 。
	// test3(1, 2)

	//标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型。
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}
}
