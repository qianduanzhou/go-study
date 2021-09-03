/*golang让机器去处理格式化，使用gofmt格式化
 */
package main

import "fmt"

func main() {
	type T struct {
		name  string // name of the object
		value int    // its value
	}
	var x = 5
	var y = 10
	fmt.Println(x<<8 + y<<16)
}
