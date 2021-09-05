package main

import (
	"fmt"
	"names/test"
)

func main() {
	test.Test()
	fmt.Println("main")

	var user = "liu"
	owner := test.Owner()
	if owner != user {
		test.SetOwner(user)
	}
}
