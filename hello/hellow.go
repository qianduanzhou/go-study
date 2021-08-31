package main

import (
	//本地模块（该模块有输出控制台等方法）
	"fmt"
	"log"

	//由于未发布该包，所以需要指向本地（1. go mod edit -replace example.com/greetings=../greetings 2. go mod tidy）
	"example.com/greetings"
	//外部模块（使用go mod tidy命令下载）
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// message := greetings.Hello("Gladys")
	// fmt.Println(message)

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("zhou")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
