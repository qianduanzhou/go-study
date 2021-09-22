package main

import (
	"fmt"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	/*无缓冲通道*/
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
	close(ch)
	/*有缓冲通道*/
	ch0 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch0 <- 10
	ret := <-ch0
	fmt.Println("接收成功", ret)
	fmt.Println("发送成功")
	close(ch0)

	/*在循环中使用通道*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		fmt.Println("first")
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		fmt.Println("second")
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println("third")
	}
	fmt.Println(3)
}
