# Control structures（流程控制）

go的流程控制只有for，没有do while，而且不需要括号()。

switch也更加灵活。

提出一种新的控制方式select。

## if

在go中，一个简单的if是这样的：

```
var x = 3
var y = 5
if x > 0 {
	fmt.Println(y)
}
```

强制大括号鼓励将简单的if语句写在多行上。无论如何，这样做都是一种很好的样式，特别是当主体包含一个控制语句(如return或break)时。

由于if和switch接受初始化语句，因此经常看到使用初始化语句来设置局部变量。

```
if err := check("11"); err != nil {
	fmt.Println(err)
}
```

在Go库中，您会发现，如果一个if语句没有流到下一个语句中，也就是说，主体将以break、continue、goto结束，或者return，不必要的else将被省略。

```
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

这是代码必须防范一系列错误条件的常见情况的一个例子。如果成功的控制流沿着页面运行，那么代码读起来很好，消除了出现错误的情况。由于错误情况往往以return语句结束，因此生成的代码不需要else语句。

```
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

## Redeclaration and reassignment（重定义和调动）

如下例子：

```
f, err := os.Open(name)
```

这个语句声明了两个变量， f和err，之后调用f.Stat()

```
d, err := f.Stat()
```

这样是可行的，虽然看起来声明了两次err，但是第二个err是重新分配并赋值，只有第一个是声明

## For

Go for循环类似于C循环，但不相同。没有while和do-while。有三种形式，只有一种有分号。

```
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

简短的声明使得在循环中很容易声明index变量。

```
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

在array，slice，string或者map中循环，可以使用range

```
for key, value := range oldMap {
    newMap[key] = value
}
```

可以删除一项，保留key或者value

```
 for key := range oldMap {
    fmt.Println(key)
 }

 for _, value := range oldMap {
    fmt.Println(value)
 }
```

空白标识符有许多用途。

对于字符串，range 做了更多的工作，通过解析UTF-8来分解单个Unicode代码点。错误的编码消耗一个字节并产生替代符文U+FFFD。名称(与相关的内置类型)符文是一个Unicode代码点的Go术语。详情请参阅语言规范。)循环

```
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

最后，Go没有逗号操作符，而且++和--都是语句而不是表达式。因此，如果你想在for中运行多个变量，你应该使用并行赋值(尽管这排除了++和--)。

```
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

## Switch

go的switch比C的更一般，表达式不需要是常量，甚至不需要是整数，情况从上到下计算，直到找到匹配，如果switch没有表达式，它就切换为true。因此，可以将if-else-if-else链写成switch。

可以像if那样进行逻辑判断，也可以像普通C的switch那样匹配字符或者字符串类型。

```
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

没有自动失败，但是可以用逗号分隔的列表来表示大小写。

```
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

switch不需要break来中断，只要匹配到就自动中断

```
d := 0
switch {
default:
	fmt.Println("未匹配")
case 5 > 4:
	d = 6
case 5 > 3:
	d = 7
}
```

## Type switch

可以用来发现接口变量的动态类型

type switch语法，将关键字type放在括号内

```
var t interface{} = true
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```