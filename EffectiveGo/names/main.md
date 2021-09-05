# Names（命名）

go的命名和其他语言一样重要，以下是go程序中的命名约定。

## package names(包)

- 包名应该是好的:简短、简洁、能引起共鸣。
- 包的名字都是小写的、单字的（bytes，fmt）。

- 不用担心冲突，包名只是导入的默认名字，它不需要在所有源代码中都是唯一的
- 另一种约定是包名是其源目录的基名。

如包在目录src/encoding/base64中，引入的时候用encoding/base64，但名字是base64而不是encoding_base64或encodingBase64。

- 导入的包将通过包的命名来使用包的内容，所以可以不同包避免导出的方法的冲突。

- 长名字不会自动使内容更具可读性。一个有用的文档注释通常比一个超长的名称更有价值。

## Getters

- Go不提供对getter和setter的自动支持，自己提供getter和setter并没有什么错。

- 假如包中有个owner字段，那么getter方法应该为Owner,而不是GetOwner，如果需要，setter函数可以被称为SetOwner。

```
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

## Interface names（接口）

按照惯例，单一方法接口由方法名加上-er后缀或类似的修改来命名，以构造代理名词:Reader、Writer、Formatter、CloseNotifier等。

## MixedCaps

最后，Go中的约定是使用MixedCaps或MixedCaps而不是下划线来写多单词名称。