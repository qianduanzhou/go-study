package main

import (
	"fmt"
)

//全局声明数组 只能用var
var arr [5]int = [5]int{}
var arr2 = [...]int{1, 2, 3}
var arr3 = [...]int{3: 5}

func main() {
	fmt.Println("arr", arr)
	fmt.Println("arr2", arr2)
	fmt.Println("arr3", arr3)
	//局部声明数组
	arr := [5]int{1, 2, 3, 4, 5}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println("arr", arr)
	fmt.Println("d", d)

	//多维数组
	mArr := [...][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println("mArr", mArr)

	//值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	test(a)
	fmt.Println(a)

	//多维数组遍历：
	for i, v := range mArr {
		fmt.Println(i, v)
		for i2, v2 := range v {
			fmt.Println(i2, v2)
		}
	}

	//数组拷贝和传参
	arr[1] = 5
	fmt.Println("arr", arr)
	//传参修改，用指针去修改
	printArr(&arr)
	fmt.Println("arr", arr)
}

//赋值
func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

//赋值
func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
