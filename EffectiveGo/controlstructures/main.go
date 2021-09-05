/* 流程控制
 */
package main

import (
	"errors"
	"fmt"
)

func check(val string) error {
	if val == "" {
		return errors.New("empty")
	}
	return nil
}

func main() {
	//简单的if
	var x = 3
	var y = 5
	if x > 0 {
		fmt.Println(y)
	}
	//可设置初始化的if
	if err := check("11"); err != nil {
		fmt.Println(err)
	}

	//for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//使用range控制循环
	oldMap := map[string]int{"a": 1, "b": 2}
	newMap := map[string]int{}
	for key, value := range oldMap {
		newMap[key] = value
	}
	fmt.Println(newMap)
	//可以删除一项，保留key或者value
	for key := range oldMap {
		fmt.Println(key)
	}
	for _, value := range oldMap {
		fmt.Println(value)
	}
	//字符串进行range可以有更多的作用
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
	//最后，Go没有逗号操作符，而且++和--都是语句而不是表达式。因此，如果你想在for中运行多个变量，你应该使用并行赋值(尽管这排除了++和--)。
	a := [5]int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(a)

	/* switch
	switch不需要break来中断，只要匹配到就自动中断
	*/
	//可以像if一样用来判断
	d := 0
	switch {
	default:
		fmt.Println("未匹配")
	case 5 > 4:
		d = 6
	case 5 > 3:
		d = 7
	}
	fmt.Println(d)
	c := '?'
	//也可以像c那样匹配字符或者字符串类型
	switch c {
	default:
		fmt.Println("未匹配")
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("匹配")
	}

	//Type switch 还可以用来发现接口变量的动态类型
	var t interface{} = true
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}
