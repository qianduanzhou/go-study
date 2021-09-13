# 条件语句if

### if 语句 if 语句 由一个布尔表达式后紧跟一个或多个语句组成。

Go 编程语言中 if 语句的语法如下：

- 可省略条件表达式括号。
- 持初始化语句，可定义代码块局部变量。 
- 代码块左 括号必须在条件表达式尾部。

    if 布尔表达式 {
    /* 在布尔表达式为 true 时执行 */
    }

if 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则不执行。

```go
 x := 0

// if x > 10        // Error: missing condition in if statement
// {
// }

if n := "abc"; x > 0 {     // 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
    println(n[2])
} else if x < 0 {    // 注意 else if 和 else 左大括号位置。
    println(n[1])
} else {
    println(n[0])
}
```

**不支持三元操作符(三目运算符) "a > b ? a : b"。**

#### 实例:

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10
   /* 使用 if 语句判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}
```

和其他语言一样也支持if else if  else的形式