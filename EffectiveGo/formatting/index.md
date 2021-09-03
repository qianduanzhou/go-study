# go格式化

- golang让机器去处理格式化，使用gofmt格式化

例如，不需要花时间排列结构字段上的注释。Gofmt会帮你的。

```
type T struct {
    name string // name of the object
    value int // its value
}
```

gofmt会将列对齐

```
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

- 标准包中的所有Go代码都使用gofmt进行了格式化。
- 一些格式细节仍然存在，如：

###### 1.行首缩进

一般使用tab缩进，gofmt默认会发出制表符。只有在必要的情况下才使用空格

###### 2.代码行长度

没有限制，太长可以换行

###### 3.圆括号

Go比C和Java需要更少的括号:控制结构(if, for, switch)在语法中没有括号。此外，操作符优先级层次结构更短更清晰

