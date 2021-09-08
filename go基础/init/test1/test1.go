package test1

import (
	"fmt"
	"init/test1/test3"
)

func init() {
	fmt.Println("test1-first")
	test3.Test()
}

func init() {
	fmt.Println("test1-second")
	test3.Test()
}
func Test() {

}
