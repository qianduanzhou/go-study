package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

// var m = make(map[string]int)
var m = sync.Map{}

// func get(key string) int {
// 	return m[key]
// }
// func set(key string, value int) {
// 	m[key] = value
// }

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}
func main() {
	wg.Add(1)
	go hello()
	fmt.Println("main goroutine done!")
	wg.Wait()

	/*sync.Map*/
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// set(key, n)
			m.Store(key, n)
			value, _ := m.Load(key)
			// fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
