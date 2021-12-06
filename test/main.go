package main

import (
	"errors"
	"fmt"
)

type father struct {
	name string
}
type child struct {
	name string
	father
}

func (f *father) test() {
	fmt.Println("father-test", f.name)
}
func (c *child) test() {
	fmt.Println("child-test", c.name)
}
func main() {
	// var f father
	// f.name = "zhou"
	// f.test()

	c := child{
		"li",
		father{
			"zhou",
		},
	}
	c.test()

	err1 := errors.New("<QuerySeter> no row found")
	err2 := errors.New("<QuerySeter> no row found")
	fmt.Println("err比较", errors.Is(err1, err2))
}
