package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//map定义
	var m map[string]int = make(map[string]int, 8)
	m["张三"] = 5
	m["李四"] = 10
	m["李四"] = 10
	m["李四2"] = 10
	m["李四3"] = 10
	m["李四4"] = 10
	m["李四5"] = 10

	fmt.Println("m", m)
	//声明时填充
	m2 := map[string]int{
		"张三": 5,
		"李四": 10,
	}
	fmt.Println("m2", m2)

	//判断某个键是否存在
	v, ok := m["张三"]
	fmt.Println(v)
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}

	//map遍历
	for k, v := range m {
		fmt.Println("k,v", k, v)
	}

	//使用delete()函数删除键值对
	delete(m, "李四5")
	fmt.Println("m", m)

	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println("scoreMap", scoreMap)

	slice := make([]string, 0, 200)
	for key := range scoreMap {
		slice = append(slice, key)
	}

	sort.Strings(slice)

	for _, value := range slice {
		fmt.Printf("%v ", scoreMap[value])
	}
	fmt.Println()

	//元素为map的slice
	mapSlice := make([]map[string]string, 3)
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["test"] = "test"
	fmt.Println("mapSlice", mapSlice)

	//值为slice的map
	sliceMap := make(map[string][]int, 10)
	sliceMap["test"] = make([]int, 5)
	sliceMap["test"][1] = 2
	fmt.Println("sliceMap", sliceMap)
	var slice2 []int = sliceMap["test"][1:3:3]
	fmt.Println("slice2", slice2)

}
