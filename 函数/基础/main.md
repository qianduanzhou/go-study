# 函数定义

### golang函数特点：

- 无需声明原型。
- 支持不定 变参。
- 支持多返回值。
- 支持命名返回参数。 
- 支持匿名函数和闭包。
- 函数也是一种类型，一个函数可以赋值给变量。
- 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
- 不支持 重载 (overload) 
- 不支持 默认参数 (default parameter)。

### 函数声明：

函数声明包含一个函数名，参数列表， 返回值列表和函数体。如果函数没有返回值，则返回列表可以省略。函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。

函数可以没有参数或接受多个参数。

注意类型在变量名之后 。

当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

函数可以返回任意数量的返回值。

使用关键字 func 定义函数，左大括号依旧不能另起一行。

```go
func test(x, y int, s string) (int, string) {
    // 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
    n := x + y          
    return n, fmt.Sprintf(s, n)
}
```

函数是第一类对象，可作为参数传递。建议将复杂签名定义为函数类型，以便于阅读。

```go
package main

import "fmt"

func test(fn func() int) int {
    return fn()
}
// 定义函数类型。
type FormatFunc func(s string, x, y int) string 

func format(fn FormatFunc, s string, x, y int) string {
    return fn(s, x, y)
}

func main() {
    s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

    s2 := format(func(s string, x, y int) string {
        return fmt.Sprintf(s, x, y)
    }, "%d, %d", 10, 20)

    println(s1, s2)
}
```

# 参数

### 函数参数

函数定义时指出，函数定义时有参数，该变量可称为函数的形参。形参就像定义在函数体内的局部变量。

但当调用函数，传递过来的变量就是函数的实参，函数可以通过两种方式来传递参数：

值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

```go
func swap(x, y int) int {
       ... ...
  }
```

引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

```go
package main

import (
    "fmt"
)

/* 定义相互交换值的函数 */
func swap(x, y *int) {
    var temp int

    temp = *x /* 保存 x 的值 */
    *x = *y   /* 将 y 值赋给 x */
    *y = temp /* 将 temp 值赋给 y*/

}

func main() {
    var a, b int = 1, 2
    /*
        调用 swap() 函数
        &a 指向 a 指针，a 变量的地址
        &b 指向 b 指针，b 变量的地址
    */
    swap(&a, &b)

    fmt.Println(a, b)
}
```

在默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。

注意2：map、slice、chan、指针、interface默认以引用的方式传递。

不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）

Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。

在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可。

```go
  func myfunc(args ...int) {    //0个或多个参数
  }

  func add(a int, args…int) int {    //1个或多个参数
  }

  func add(a int, b int, args…int) int {    //2个或多个参数
  }
```

注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数.

任意类型的不定参数： 就是函数的参数和每个参数的类型都不是固定的。

用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的。

```go
func myfunc(args ...interface{}) {
  }
```

代码：

```go
package main

import (
    "fmt"
)

func test(s string, n ...int) string {
    var x int
    for _, i := range n {
        x += i
    }

    return fmt.Sprintf(s, x)
}

func main() {
    println(test("sum: %d", 1, 2, 3))
}
```

使用 slice 对象做变参时，必须展开。`（slice...）`

```go
package main

import (
    "fmt"
)

func test(s string, n ...int) string {
    var x int
    for _, i := range n {
        x += i
    }

    return fmt.Sprintf(s, x)
}

func main() {
    s := []int{1, 2, 3}
    res := test("sum: %d", s...)    // slice... 展开slice
    println(res)
}
```

# 返回值

### 函数返回值

`"_"`标识符，用来忽略函数的某个返回值

Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用。

返回值的名称应当具有一定的意义，可以作为文档使用。

没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。

直接返回语句仅应当用在像下面这样的短函数中。在长的函数中它们会影响代码的可读性。

```go
package main

import (
    "fmt"
)

func add(a, b int) (c int) {
    c = a + b
    return
}

func calc(a, b int) (sum int, avg int) {
    sum = a + b
    avg = (a + b) / 2

    return
}

func main() {
    var a, b int = 1, 2
    c := add(a, b)
    sum, avg := calc(a, b)
    fmt.Println(a, b, c, sum, avg)
}
```

Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 `"_"` 忽略。

```go
package main

func test() (int, int) {
    return 1, 2
}

func main() {
    // s := make([]int, 2)
    // s = test()   // Error: multiple-value test() in single-value context

    x, _ := test()
    println(x)
}
```

# 匿名函数

### 匿名函数

匿名函数是指不需要定义函数名的一种函数实现方式。1958年LISP首先采用匿名函数。

在Go里面，函数可以像普通变量一样被传递或使用，Go语言支持随时在代码里定义匿名函数。

匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    getSqrt := func(a float64) float64 {
        return math.Sqrt(a)
    }
    fmt.Println(getSqrt(4))
}
```

上面先定义了一个名为getSqrt 的变量，初始化该变量时和之前的变量初始化有些不同，使用了func，func是定义函数的，可是这个函数和上面说的函数最大不同就是没有函数名，也就是匿名函数。这里将一个函数当做一个变量一样的操作。

Golang匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。

```go
package main

func main() {
    // --- function variable ---
    fn := func() { println("Hello, World!") }
    fn()

    // --- function collection ---
    fns := [](func(x int) int){
        func(x int) int { return x + 1 },
        func(x int) int { return x + 2 },
    }
    println(fns[0](100))

    // --- function as field ---
    d := struct {
        fn func() string
    }{
        fn: func() string { return "Hello, World!" },
    }
    println(d.fn())

    // --- channel of function ---
    fc := make(chan func() string, 2)
    fc <- func() string { return "Hello, World!" }
    println((<-fc)())
}
```