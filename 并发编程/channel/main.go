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
	ch0 := make(chan int, 2) // 创建一个容量为1的有缓冲区通道
	ch0 <- 10
	ch0 <- 101
	ret := <-ch0
	fmt.Println("接收成功1", ret)
	ret2 := <-ch0
	fmt.Println("接收成功2", ret2)

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
		fmt.Println(i)
	}
	fmt.Println("third")
	/*控制通道读写*/
	ch3 := make(chan int)
	ch4 := make(chan int)
	go counter(ch3)
	go squarer(ch4, ch3)
	printer(ch4)
}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
