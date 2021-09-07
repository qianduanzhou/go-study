package slice

import "fmt"

//影响之前的slice
func foo(s []int) {
	for index := range s {
		s[index] += 1
	}
}

//不影响之前的slice
func copyFoo(s []int) {
	slice := s[0 : len(s)-1 : len(s)-1]
	copyS := append(slice, s[len(s)-1])
	// var copyS []int = make([]int, len(s))
	copy(copyS, s)
	fmt.Println("copyS", copyS)
	for index := range copyS {
		copyS[index] += 1
	}
}

func Slice() {
	/*slice
	 */
	my_slice := make([]int, 3, 5)
	// 输出slice
	fmt.Println(my_slice)        // 输出：[0 0 0]
	fmt.Println(len(my_slice))   // 输出slice长度 3
	fmt.Println(cap(my_slice))   // 输出底层数组长度 5
	print("my_slice：", my_slice) //全部输出 [slice长度/底层数组长度]slice指针指向的底层数组的元素
	fmt.Println("————————————————")

	fmt.Println("————slice初始化————")
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
	fmt.Println("————————————————")

	fmt.Println("————snil slice————")
	//nil slice
	var nil_slice []int
	fmt.Printf("nil_slice：%v\n", nil_slice)
	print("nil_slice：", nil_slice)
	fmt.Println("————————————————")

	fmt.Println("————空slice————")
	// 空slice 使用make创建
	empty_slice := make([]int, 0)
	// 直接创建
	// empty_slice := []int{}
	fmt.Printf("empty_slice：%v\n", empty_slice)
	print("empty_slice：", empty_slice)
	fmt.Println("————————————————")

	fmt.Println("————切片————")
	//对slice进行切片切片
	my_slice2 := []int{11, 22, 33, 44, 55}
	// 生成新的slice，从第二个元素取，切取的长度为2
	new_slice2 := my_slice2[1:3]
	fmt.Printf("new_slice2：%v\n", new_slice2)
	//还可以控制新的slice的容量
	my_slice3 := []int{11, 22, 33, 44, 55}
	// 从第二个元素取，切取的长度为2，容量也为2
	new_slice3 := my_slice3[1:3:3]
	fmt.Printf("new_slice3：%v\n", new_slice3)
	fmt.Println("————————————————")

	fmt.Println("————copy函数————")
	//copy函数
	s1 := []int{11, 22, 33}
	s2 := make([]int, 5)
	s3 := make([]int, 2)
	num := copy(s2, s1)
	copy(s3, s1)
	fmt.Println(num) // 3
	fmt.Println(s2)  // [11,22,33,0,0]
	fmt.Println(s3)  // [11,22]
	fmt.Println("————————————————")

	fmt.Println("————copy byte————")
	h := []byte("Hello")
	fmt.Println(h)
	hNum := copy(h, "World")
	fmt.Println(hNum)
	fmt.Println(h)         // 输出[87 111 114 108 100 32]
	fmt.Println(string(h)) //输出"World"
	fmt.Println("————————————————")

	fmt.Println("————append()函数————")
	my_slice4 := []int{11, 22, 33, 44, 55}
	new_slice4 := my_slice4[1:3]
	new_slice5 := my_slice4[1:4]
	// append()追加一个元素2323，返回新的slice
	app_slice := append(new_slice4, 2323)
	my_slice4[1] = 222
	app_slice[2] = 333
	fmt.Println("new_slice4", new_slice4)
	fmt.Println("app_slice", app_slice)
	fmt.Println("my_slice4", my_slice4)
	fmt.Println("new_slice5", new_slice5)

	s5 := []byte("Hello")
	s6 := append(s5, "World"...)
	fmt.Println(string(s6)) // 输出：HelloWorld
	fmt.Println("————————————————")

	fmt.Println("————append 扩容————")

	my_slice6 := []int{11, 22, 33, 44, 55}
	new_slice6 := append(my_slice6, 66)

	new_slice6[3] = 444 // 修改旧的底层数组
	my_slice6[1] = 222
	fmt.Println(my_slice6)
	fmt.Println(new_slice6)

	fmt.Println(len(my_slice6), ":", cap(my_slice6))   // 5:5
	fmt.Println(len(new_slice6), ":", cap(new_slice6)) // 6:10
	fmt.Println("————————————————")

	fmt.Println("————合并slice————")
	s7 := []int{1, 2}
	s8 := []int{3, 4}

	s9 := append(s7, s8...)

	fmt.Println(s9) // [1 2 3 4]
	fmt.Println("————————————————")

	fmt.Println("————slice遍历迭代————")
	s10 := []int{11, 22, 33, 44}
	for index, value := range s10 {
		println("index:", index, " , ", "value", value)
	}
	fmt.Println("————————————————")

	fmt.Println("————传递slice给函数————")
	s11 := []int{11, 22, 33, 44}
	s12 := []int{11, 22, 33, 44}
	foo(s11)
	copyFoo(s12)
	fmt.Println(s11[1]) // 输出：23
	fmt.Println(s12[1]) // 扩容了，所以不影响 输出：22
	fmt.Println("————————————————")

}
