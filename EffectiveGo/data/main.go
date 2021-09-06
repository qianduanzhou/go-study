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
	*p3 = make([]int, 100, 100)

	// Idiomatic:
	v3 := make([]int, 100)
	fmt.Printf("p2：%v\n", p2)
	fmt.Printf("v2：%v\n", v2)
	fmt.Printf("p3：%v\n", p3)
	fmt.Printf("v3：%v\n", v3)
}
