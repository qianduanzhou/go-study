/**
包注释 每个包前最好都要有注释
*/
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
	// message := fmt.Sprint(randomFormat()) //单元测试TestHelloName不通过的例子
	return message, nil //nil代表没有错误
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
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
