package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
//大写字母开头的函数可以在外部使用
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.

	// var message string = fmt.Sprintf("Hi, %v. Welcome!", name)
	// message := fmt.Sprintf("Hi, %v. Welcome!", name) // %v是一个占位符，name将会取代他

	message := fmt.Sprintf(randomFormat(), name) // %v是一个占位符，name将会取代他

	return message, nil //nil代表没有错误
}

// init sets initial values for variables used in the function.
//小写字母开头的函数只能在内部使用
//init是模块初始化会执行的函数（类似生命周期）
func init() {
	//设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	rand.Seed(time.Now().UnixNano())
}
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
