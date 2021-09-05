# Functions

## Multiple return values（多个返回值）

go函数可以返回多个返回值

在Go中，Write可以返回一个计数和一个错误:是的，你写了一些字节，但不是全部，因为你填满了设备。来自包操作系统的文件的Write方法签名为

```
func (file *File) Write(b []byte) (n int, err error)
```

## Named result parameters（命名的返回值）

函数返回值可以先命名并自动初始化为0，如果函数没有return该值，就使用命名的参数作为返回值。

```
func test() (i int) {
	// i = 5
	return
}

func main() {
	i := test()
	fmt.Println(i) //0
}
```

## Defer

有defer的函数执行顺序是：

1. 从上到下执行，遇到defer，如果defer后接的是函数，则执行函数内的语句，直到把函数返回值放入栈，如果不是则直接放入队列。
2. 多个defer依次按第一点放入栈。
3. 在有defer的函数执行完成前会将栈里的语句按后进先出依次执行。
4. 最后再是函数本身的返回值执行。

``` 
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() string {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
	return "end"
}

func main() {
	fmt.Println(b())
}
```

输出：

```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
end
```

defer函数按后进先出顺序执行，因此当函数返回时，这段代码将导致打印4 3 2 10

```
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```