/* test包
 */
package test

import "fmt"

var owner = "zhou"

//getter方法，获取owner字段
func Owner() string {
	return owner
}

//setter方法
func SetOwner(user string) {
	owner = user
	fmt.Println(owner)
}

func Test() {
	fmt.Println("test")
}
