package main

import "fmt"

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

}
