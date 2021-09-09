package main

import "fmt"

func main() {
	//指针
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
	fmt.Println(*b)                    // 10

	//指针传值示例：
	c := 10
	modify1(c)
	fmt.Println(c) // 10
	modify2(&c)
	fmt.Println(c) // 100

	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}

	//以下代码会引发panic
	// var a2 *int
	// *a2 = 100
	// fmt.Println(*a2)
	// var b2 map[string]int
	// b2["测试"] = 100
	// fmt.Println(b2)

	//new
	a2 := new(int)
	fmt.Println("a2", a2)
	*a2 = 5
	fmt.Println("a2", *a2)
	var a3 *int
	a3 = new(int)
	fmt.Println("a3", *a3)
	*a3 = 5
	fmt.Println("a3", *a3)

	//make
	var b3 map[string]int
	b3 = make(map[string]int, 10)
	b3["测试"] = 100
	fmt.Println("b3", b3)
}

//传值
func modify1(x int) {
	x = 100
}

//传地址
func modify2(x *int) {
	*x = 100
}
