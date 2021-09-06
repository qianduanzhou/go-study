# Data

## Allocation with `new`（new函数）

Go有两个分配原语，内置函数new和make。它们做不同的事情，适用于不同的类型

- new

它是一个分配内存的内置函数，但不同于其他语言中的同名函数，它不初始化内存，只是将其归零。

也就是说，new(T)为类型为T的新项分配零存储空间，并返回它的地址，类型为*T。

在Go术语中，它返回一个指针，指向新分配的T类型的零值。

由于new返回的内存是归零的，所以在设计数据结构时安排每种类型的零值无需进一步初始化即可使用是很有帮助的。这意味着数据结构的用户可以用new创建一个数据结构并直接开始工作。例如，bytes.Buffer。Buffer声明“Buffer的零值是一个准备使用的空缓冲区”。同样,sync.Mutex没有显式的构造函数或Init方法。而是同步的零值，sync.Mutex定义为一个未锁定的mutex。

“零值即有用”的属性是传递性的。考虑这个类型声明。

```
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

SyncedBuffer类型的值也可以在分配或声明后立即使用。在下一个代码片段中，p和v都可以正常工作，无需进一步安排。

```
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

## Constructors and composite literals（构造函数和复合字面量）

有时候，零值还不够好，需要初始化构造函数，如这个来自包os的示例所示。

```
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

上述代码有些多，我们可以使用复合字面值简化它，复合字面值是一个表达式，每次计算时都会创建一个新的实例。

```
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

注意，与C语言不同，返回局部变量的地址是完全可以的;与变量关联的存储在函数返回后继续存在。实际上，每次计算复合字面值时，取复合字面值的地址都会分配一个新实例，因此我们可以将最后两行合并起来。

```
    return &File{fd, name, nil, 0}
```

复合文本的字段是按顺序排列的，并且必须全部出现。然而，通过显式地将元素标记为字段:值对，初始化器可以以任何顺序出现，而缺失的元素作为它们各自的零值。因此我们可以说

```
    return &File{fd: fd, name: name}
```

作为一种限制情况，如果复合文本根本不包含字段，它将为该类型创建一个零值。表达式new(File)和&File{}是等价的。

还可以为arrays, slices, and maps创建复合文本，适当的字段标签是索引或map的key值。在这些例子中，不管Enone、Eio和Einval的值如何，初始化都是有效的，只要它们是不同的。

```
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```

## Allocation with `make`

内置函数make(T, args)的用途与new(T)不同，它只创建slices, maps, and channels，并返回类型为T(不是*T)的初始化(不是归零)值。

区别的原因是，这三种类型代表了对数据结构的引用，这些引用在使用之前必须初始化。

例如，一个slice是一个包含三项的描述符，它包含一个指向数据(在array中)的指针、长度和容量，在这些项被初始化之前，slice是nil。

对于slices, maps和channels，make初始化内部数据结构并准备使用值。例如

```
make([]int, 10, 100)
```

分配一个100个整型数的array，然后创建一个长度为10、容量为100的slice结构，指向数组的前10个元素。

相反，new([]int)返回一个指向新分配的归零slice结构的指针，也就是说，一个指向nil切片值的指针。

加上\*代表指针，指向该变量的内存单元。

如果变量本身是指针，加上*就代表对这个指针解引用，也就是取这个指针指向的变量的值。

这些例子说明了new和make之间的区别。

```
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```

记住，make只应用于maps, slices 和 channels，不返回指针。要获取显式指针，请使用new进行分配或显式地获取变量的地址。