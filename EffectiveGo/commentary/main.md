# 注释

- go行的注释有c的/**/与c++的//两种注释。块注释主要以包注释的形式出现，但在表达式中很有用，或禁用大量代码。

- 程序和web服务器godoc处理Go源文件，以提取有关包内容的文档。出现在顶级声明之前的注释(中间没有换行)与声明一起被提取，作为项目的解释性文本。这些注释的性质和风格决定了godoc生成的文档的质量。

- 每个包都应该有一个包注释，在包子句之前有一个块注释。对于多文件包，包注释只需要出现在一个文件中，任何一个文件都可以。包注释应该介绍包并提供与包整体相关的信息。它将首先出现在godoc页面上，并设置接下来的详细文档。

```
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

如果包很简单，那么包注释可以很简短。

```
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

- 注释不需要额外的格式。生成的输出甚至可能不会以固定宽度的字体显示，所以不要依赖间距来对齐，godoc(比如gofmt)会处理这个问题。注释是未经解释的纯文本，因此HTML和其他类似的注释将逐字复制，不应该使用。godoc做的一项调整是用固定宽度的字体显示缩进的文本，这适合于程序片段。fmt包的包注释很好地使用了这一点。
- 根据上下文的不同，godoc甚至可能不会重新格式化注释，所以要确保注释看起来很好:使用正确的拼写、标点符号和句子结构，折叠长行，等等。
- 在包中，紧挨着顶级声明前面的任何注释都可以作为该声明的文档注释。程序中每个导出的(大写的)名字都应该有一个doc注释。
- Doc注释最好是完整的句子，它允许各种各样的自动化演示。第一句话应该是一个单句总结，以声明的名称开始。

```
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

- 如果每个doc注释都以它所描述的项目的名称开头，那么您可以使用go工具的doc子命令并通过grep运行输出。假设您不记得“Compile”这个名字，但正在查找正则表达式的解析函数，因此运行该命令

```
$ go doc -all regexp | grep -i parse
    Compile parses a regular expression and returns, if successful, a Regexp
    MustCompile is like Compile but panics if the expression cannot be parsed.
    parsed. It simplifies safe initialization of global variables holding
$
```

Go的声明语法允许对声明进行分组。一个文档注释可以引入一组相关的常量或变量。由于整个声明是呈现的，这样的评论往往是敷衍的。

```
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

分组还可以指示项之间的关系，例如一组变量是由互斥锁保护的。

```
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```