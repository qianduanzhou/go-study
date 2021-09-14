# 闭包、递归

### Go的闭包

```go
package main

import (
    "fmt"
)

func a() func() int {
    i := 0
    b := func() int {
        i++
        fmt.Println(i)
        return i
    }
    return b
}

func main() {
    c := a()
    c()
    c()
    c()

    a() //不会输出i
}
```

输出结果：

```
    1
    2
    3
```

闭包复制的是原对象指针，这就很容易解释延迟引用现象。

```go
package main

import "fmt"

func test() func() {
    x := 100
    fmt.Printf("x (%p) = %d\n", &x, x)

    return func() {
        fmt.Printf("x (%p) = %d\n", &x, x)
    }
}

func main() {
    f := test()
    f()
}
```

外部引用函数参数局部变量

```go
package main

import "fmt"

// 外部引用函数参数局部变量
func add(base int) func(int) int {
    return func(i int) int {
        base += i
        return base
    }
}

func main() {
    tmp1 := add(10)
    fmt.Println(tmp1(1), tmp1(2))
    // 此时tmp1和tmp2不是一个实体了
    tmp2 := add(100)
    fmt.Println(tmp2(1), tmp2(2))
}
```

返回2个闭包

```go
package main

import "fmt"

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
    // 定义2个函数，并返回
    // 相加
    add := func(i int) int {
        base += i
        return base
    }
    // 相减
    sub := func(i int) int {
        base -= i
        return base
    }
    // 返回
    return add, sub
}

func main() {
    f1, f2 := test01(10)
    // base一直是没有消
    fmt.Println(f1(1), f2(2))
    // 此时base是9
    fmt.Println(f1(3), f2(4))
}
```

### Go 语言递归函数

递归，就是在运行的过程中调用自己。 一个函数调用自己，就叫做递归函数。

构成递归需具备的条件：

1. 子问题须与原始问题为同样的事，且更为简单。
2. 不能无限制地调用本身，须有个出口，化简为非递归状况处理。

#### 数字阶乘

一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，并且0的阶乘为1。自然数n的阶乘写作n!。1808年，基斯顿·卡曼引进这个表示法。

```go
package main

import "fmt"

func factorial(i int) int {
    if i <= 1 {
        return 1
    }
    return i * factorial(i-1)
}

func main() {
    var i int = 7
    fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}
```

### 斐波那契数列(Fibonacci)

这个数列从第3项开始，每一项都等于前两项之和。

```go
package main

import "fmt"

func fibonaci(i int) int {
    if i == 0 {
        return 0
    }
    if i == 1 {
        return 1
    }
    return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
    var i int
    for i = 0; i < 10; i++ {
        fmt.Printf("%d\n", fibonaci(i))
    }
}
```