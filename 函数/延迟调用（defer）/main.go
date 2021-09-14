package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func test() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制
	x += 10
	y += 100
	println("x =", x, "y =", y)
}

var lock sync.Mutex

func testLock() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}

func bar() (i int) {

	i = 0
	defer func() {
		fmt.Println(i)
	}()

	return 2
}

func main() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}

	//defer碰上闭包
	var whatever2 [5]struct{}
	for i := range whatever2 {
		defer func() { fmt.Println(i) }()
	}

	ts := []Test{{"a"}, {"b"}, {"c"}}
	defer ts[0].Close()
	defer ts[1].Close()

	for _, t := range ts {
		t2 := t
		defer t2.Close()
	}

	/*延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。*/
	test()

	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			testLock()
		}
		elapsed := time.Since(t1)
		fmt.Println("test elapsed: ", elapsed)
	}()
	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			testdefer()
		}
		elapsed := time.Since(t1)
		fmt.Println("testdefer elapsed: ", elapsed)
	}()

	//defer 与 closure
	foo(2, 0)
	//defer 与 return
	fmt.Println("bar", bar())
}
