package main

import (
	"bytes"
	"fmt"
	"sync"
)

//复核字面量
// func NewFile(fd int, name string) *File {
// 	if fd < 0 {
// 		return nil
// 	}
// 	// f := File{fd, name, nil, 0}
// 	// return &f
// 	// return &File{fd, name, nil, 0}

// 	//new(File)和&File{}是等价的
// 	return &File{fd: fd, name: name}
// }

//传递一个数组的指针
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

//slice长度改变
func Append(slice, data []byte) []byte {
	l := len(slice)
	fmt.Println("len(slice),cap(slice)", len(slice), cap(slice))
	if l+len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}
func main() {
	//new 方法
	type SyncedBuffer struct {
		lock   sync.Mutex
		buffer bytes.Buffer
	}
	p := new(SyncedBuffer) // type *SyncedBuffer
	var v SyncedBuffer     // type  SyncedBuffer
	fmt.Println(p)
	fmt.Println(v)

	// a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	// s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	// m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	//以下的例子的new和make的不同
	var p2 *[]int = new([]int)      // allocates slice structure; *p == nil; rarely useful
	var v2 []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	// Unnecessarily complex:
	var p3 *[]int = new([]int)
	*p3 = make([]int, 100)

	// Idiomatic:
	v3 := make([]int, 100)
	fmt.Printf("p2：%v\n", p2)
	fmt.Printf("v2：%v\n", v2)
	fmt.Printf("p3：%v\n", p3)
	fmt.Printf("v3：%v\n", v3)
	//array
	array := [...]float64{7.0, 8.5, 9.1}
	// array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	fmt.Printf("array：%v, x：%v\n", array, x)

	/*slice
	 */
	my_slice := make([]int, 3, 5)
	// 输出slice
	fmt.Println(my_slice)        // 输出：[0 0 0]
	fmt.Println(len(my_slice))   // 输出slice长度 3
	fmt.Println(cap(my_slice))   // 输出底层数组长度 5
	print("my_slice：", my_slice) //全部输出 [slice长度/底层数组长度]slice指针指向的底层数组的元素

	//初始化slice
	//make方式
	// 创建一个length和capacity都等于5的slice
	slice := make([]int, 5)
	// length=3,capacity=5的slice
	// slice := make([]int,3,5)
	fmt.Printf("slice：%v\n", slice)

	//还可以直接赋值初始化的方式创建slice：
	// 创建长度和容量都为4的slice，并初始化赋值
	color_slice := []string{"red", "blue", "black", "green"}
	// 创建长度和容量为100的slice，并为第1个赋值为5，100个元素赋值为3
	numSlice := []int{0: 5, 99: 3}
	fmt.Printf("color_slice：%v\n", color_slice)
	fmt.Printf("numSlice：%v\n", numSlice)
	fmt.Printf("numSlice-length：%v\n", numSlice[99])

	s := make([]byte, 10)
	d := []byte{}
	appendSlice := Append(s, d)
	fmt.Printf("appendSlice：%v\n", appendSlice)
}
