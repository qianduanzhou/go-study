package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串转义
	fmt.Println("str := \"c:\\pprof\\main.exe\"")

	/*字符串常见操作
	 */
	str := "test"
	//1.求长度
	fmt.Println(len(str))
	//2.拼接字符串
	fmt.Println(str + "1")
	//3.分割
	fmt.Println(strings.Split(str, "e"))
	//4.判断是否包含
	fmt.Println(strings.Contains(str, "e"))
	//5.前缀/后缀判断
	fmt.Println(strings.HasPrefix(str, "t"))
	fmt.Println(strings.HasSuffix(str, "t"))
	//6.子串出现的位置(找第一个)
	fmt.Println(strings.Index(str, "t"))     //->
	fmt.Println(strings.LastIndex(str, "t")) //<-
	//7.join操作（将slice拼接成一个string）
	slice := []string{"a", "b", "c"}
	fmt.Println(strings.Join(slice, "，"))
}
