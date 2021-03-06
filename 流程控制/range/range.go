package main

import "fmt"

func main() {
	//range 会复制对象。
	a := [3]int{0, 1, 2}

	for i, v := range a { // index、value 都是从复制品中取出。

		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}
		println(i, v)
		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
	}

	fmt.Println(a) // 输出 [100, 101, 102]。

	//改用引用类型，其底层数据不会被复制
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s { // 复制 struct slice { pointer, len, cap }。

		if i == 0 {
			s = s[:3]  // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。
		}

		println(i, v)
	}
	println(s)
	fmt.Println("s", s)
}
