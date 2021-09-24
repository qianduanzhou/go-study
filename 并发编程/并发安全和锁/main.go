package main

import (
	"fmt"
	"sync"
	"time"
)

/*
有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。
类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。
*/
var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func add2() {
	lock.Lock() // 加锁
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	lock.Unlock() // 解锁
	wg.Done()
}

func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)
	/*
		上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，
		导致最后的结果与期待的不符。
	*/

	/*互斥锁
	互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
	Go语言中使用sync包的Mutex类型来实现互斥锁。
	使用互斥锁来修复上面代码的问题：
	使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
	当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。
	*/
	wg.Add(2)
	go add2()
	go add2()
	wg.Wait()
	fmt.Println(x)

	/*读写互斥锁
	互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。
	读写锁在Go语言中使用sync包中的RWMutex类型。
	读写锁分为两种：读锁和写锁。
	当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
	当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
	就是说遇到写锁的情况下，只有写锁解开其他的goroutine对应的读写锁才能解开
	*/
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	fmt.Println("lock")
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	fmt.Println("unlock")
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock() // 加读锁
	fmt.Println("rlock")
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	fmt.Println("unrlock")
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}
