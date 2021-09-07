# Data

## Allocation with `new`（new函数）

Go有两个分配原语，内置函数new和make。它们做不同的事情，适用于不同的类型

- new

它是一个分配内存的内置函数，但不同于其他语言中的同名函数，它不初始化内存，只是将其归零。

也就是说，new(T)为类型为T的新项分配零存储空间，并返回它的地址，类型为*T。（new返回的是一个指针）

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

加上&该变量的地址，指向该变量的内存单元。

加上*就代表对这个变量是指针，并且对指针解引用，也就是取这个指针指向的变量的值。

如\*p就是代表p是指针，*p是指针指向的变量的值。

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

## Arrays

go的数组和c的不一样

- 数组是值类型。将一个数组赋值给另一个数组将复制所有元素。
- 特别是，如果你将一个数组传递给一个函数，它将收到一个数组的副本，而不是指向它的指针。
- 数组的大小是数组类型的一部分。[10]int和[20]int类型是不同的。

value属性可能很有用，但也很昂贵;如果你想要类似c语言的行为和效率，你可以传递一个指向数组的指针。

```
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator
```

但即使是这种风格也不是地道的Go。使用slices会更好。

## Slices

slices对array进行包装，为数据序列提供更通用、更强大和更方便的接口。

slices保存对底层array的引用，如果将一个slice分配给另一个slice，则两个片slice指向相同的array。

### slice的存储结构

如下是对长度为5，数据类型为int的数组从前往后取3个元素作为slice的结构。

```
my_slice := make([]int, 3, 5)
fmt.Println(my_slice)  // 输出：[0 0 0]
```

**每一个slice结构都由3部分组成：容量(capacity)、长度(length)和指向底层数组某元素的指针，它们各占8字节(1个机器字长，64位机器上一个机器字长为64bit，共8字节大小，32位架构则是32bit，占用4字节)，所以任何一个slice都是24字节(3个机器字长)**。

- Pointer：表示该slice结构从底层数组的哪一个元素开始，该指针指向该元素
- Capacity：即底层数组的长度，表示这个slice目前最多能扩展到这么长
- Length：表示slice当前的长度，如果追加元素，长度不够时会扩展，最大扩展到Capacity的长度(不完全准确，后面数组自动扩展时解释)，所以Length必须不能比Capacity更大，否则会报错

可以通过len()函数获取slice的长度，通过cap()函数获取slice的Capacity。

```
my_slice := make([]int,3,5)

fmt.Println(len(my_slice))  // 3
fmt.Println(cap(my_slice))  // 5
```

对上面创建的slice来说，它的长度为3，容量为5，指针指向底层数组的index=0。

还可以直接通过print()或println()函数去输出slice，它将得到这个slice结构的属性值，也就是length、capacity和pointer：

```
my_slice := make([]int,3,5)
println(my_slice)      // [3/5]0xc42003df10
```

### 创建、初始化、访问slice

有几种创建slice数据结构的方式。

一种是使用make()：

```
// 创建一个length和capacity都等于5的slice
slice := make([]int,5)

// length=3,capacity=5的slice
slice := make([]int,3,5)
```

make()比new()函数多一些操作，new()函数只会进行内存分配并做默认的赋0初始化，而make()可以先为底层数组分配好内存，然后从这个底层数组中再额外生成一个slice并初始化。另外，make只能构建slice、map和channel这3种结构的数据对象，因为它们都指向底层数据结构，都需要先为底层数据结构分配好内存并初始化。

还可以直接赋值初始化的方式创建slice：

```
// 创建长度和容量都为4的slice，并初始化赋值
color_slice := []string{"red","blue","black","green"}

// 创建长度和容量为100的slice，并为第1个赋值为5，100个元素赋值为3
slice := []int{99:3}
```

注意区分array和slice：

```
// 创建长度为3的int数组
array := [3]int{10, 20, 30}

// 创建长度和容量都为3的slice
slice := []int{10, 20, 30}
```

由于slice底层是数组，所以可以使用索引的方式访问slice，或修改slice中元素的值：

```
// 创建长度为5、容量为5的slice
my_slice := []int{11,22,33,44,55}

// 访问slice的第2个元素
print(my_slice[1])

// 修改slice的第3个元素的值
my_slice[2] = 333

```

由于slice的底层是数组，所以访问`my_slice[1]`实际上是在访问它的底层数组的对应元素。slice能被访问的元素只有length范围内的元素，那些在length之外，但在capacity之内的元素暂时还不属于slice，只有在slice被扩展时(见下文append)，capacity中的元素才被纳入length，才能被访问。

### nil slice和空slice

当声明一个slice，但不做初始化的时候，这个slice就是一个nil slice。

```
var nil_slice []int
```

nil slice表示它的指针为nil，也就是这个slice不会指向哪个底层数组。也因此，nil slice的长度和容量都为0。

```
|--------|---------|----------|
|  nil   |   0     |     0    |
|  ptr   | Length  | Capacity |
|--------|---------|----------|
```

还可以创建空slice(Empty Slice)，空slice表示长度为0，容量为0，但却有指向的slice，只不过指向的底层数组暂时是长度为0的空数组。

```
// 使用make创建
empty_slice := make([]int,0)

// 直接创建
empty_slice := []int{}
```

Empty Slice的结构如下：

```
|--------|---------|----------|
|  ADDR  |   0     |     0    |
|  ptr   | Length  | Capacity |
|--------|---------|----------|
```

虽然nil slice和Empty slice的长度和容量都为0，输出时的结果都是`[]`，且都不存储任何数据，但它们是不同的。nil slice不会指向底层数组，而空slice会指向底层数组，只不过这个底层数组暂时是空数组。

可以使用println()来输出验证：

```
println(nil_slice) [0/0]0x0
println(empty_slice) [0/0]0xc042085f50
```

### 对slice进行切片

可以从slice中继续切片生成一个新的slice，这样能实现slice的缩减。

对slice切片的语法为：

```
SLICE[A:B]
SLICE[A:B:C]
```

其中A表示从SLICE的第几个元素开始切，**B控制切片的长度(B-A)，C控制切片的容量(C-A)，如果没有给定C，则表示切到底层数组的最尾部**。

还有几种简化形式：

```
SLICE[A:]  // 从A切到最尾部
SLICE[:B]  // 从最开头切到B(不包含B)
SLICE[:]   // 从头切到尾，等价于复制整个SLICE
```

例如：

```
my_slice := []int{11,22,33,44,55}
// 生成新的slice，从第二个元素取，切取的长度为2
new_slice := my_slice[1:3]
```

注意，截取时"左闭右开"。所以上面`new_slice`是从`my_slice`的index=1开始截取，截取到index=3为止，但不包括index=3这个元素。所以，新的slice是由`my_slice`中的第2个元素、第3个元素组成的新的数据结构，长度为2。

还可以控制切片时新slice的容量：

```
my_slice := []int{11,22,33,44,55}

// 从第二个元素取，切取的长度为2，容量也为2
new_slice := my_slice[1:3:3]
```

这时新slice的length等于capacity，底层数组的index=4、5将对new_slice永不可见，即使后面对new_slice进行append()导致底层数组扩容也仍然不可见。具体见下文。

由于多个slice共享同一个底层数组，所以当修改了某个slice中的元素时，其它包含该元素的slice也会随之改变，因为slice只是一个指向底层数组的指针(只不过这个指针不纯粹，多了两个额外的属性length和capacity)，实际上修改的是底层数组的值，而底层数组是被共享的。

当同一个底层数组有很多slice的时候，一切将变得混乱不堪，因为我们不可能记住谁在共享它，通过修改某个slice的元素时，将也会影响那些可能我们不想影响的slice。所以，需要一种特性，保证各个slice的底层数组互不影响，相关内容见下面的"扩容"。

### copy()函数

可以将一个slice拷贝到另一个slice中。

```
$ go doc builtin copy
func copy(dst, src []Type) int
```

这表示将src slice拷贝到dst slice，src比dst长，就截断，src比dst短，则只拷贝src那部分。（以dst为主）

copy的返回值是拷贝成功的元素数量，所以也就是src slice或dst slice中最小的那个长度。

```
s1 := []int{11, 22, 33}
s2 := make([]int, 5)
s3 := make([]int,2)

num := copy(s2, s1)
copy(s3,s1)

fmt.Println(num)   // 3
fmt.Println(s2)    // [11,22,33,0,0]
fmt.Println(s3)    // [11,22]
```

此外，copy还能将字符串拷贝到byte slice中，因为字符串实际上就是`[]byte`。

```
s1 := []byte("Hello")
num := copy(s1, "World")
fmt.Println(num)
fmt.Println(s1)    // 输出[87 111 114 108 100 32]
fmt.Println(string(s1))  //输出"World"
```

### append()函数

可以使用append()函数对slice进行扩展，因为它追加元素到slice中，所以一定会增加slice的长度。

但必须注意，append()的结果必须被使用。所谓被使用，可以将其输出、可以赋值给某个slice。如果将append()放在空上下文将会报错：append()已评估，但未使用。同时这也说明，append()返回一个新的slice，原始的slice会保留不变。

```
my_slice := []int{11,22,33,44,55}
new_slice := my_slice[1:3]

// append()追加一个元素2323，返回新的slice
app_slice := append(new_slice,2323)
```

上面的append()在`new_slice`的后面增加了一个元素2323，所以`app_slice[2]=2323`。但因为这些slice共享同一个底层数组，所以2323也会反映到其它slice中。

同样，由于string的本质是[]byte，所以string可以append到byte slice中：

```
s1 := []byte("Hello")
s2 := append(s1, "World"...)
fmt.Println(string(s2))   // 输出：HelloWorld
```

### 扩容

当slice的length已经等于capacity的时候，再使用append()给slice追加元素，会自动扩展底层数组的长度。

**底层数组扩展时，会生成一个新的底层数组。所以旧底层数组仍然会被旧slice引用，新slice和旧slice不再共享同一个底层数组**。

```
my_slice := []int{11,22,33,44,55}
new_slice := append(my_slice,66)

my_slice[3] = 444   // 修改旧的底层数组

fmt.Println(my_slice)   // [11 22 33 444 55]
fmt.Println(new_slice)  // [11 22 33 44 55 66]

fmt.Println(len(my_slice),":",cap(my_slice))     // 5:5
fmt.Println(len(new_slice),":",cap(new_slice))   // 6:10
```

从上面的结果上可以发现，底层数组被扩容为10，且是新的底层数组。

**当capacity需要扩容时，会按照当前capacity的2倍对数组进行扩容**。

扩容的对象是底层数组的真子集时：

```
my_slice := []int{11,22,33,44,55}

// 限定长度和容量，且让长度和容量相等
new_slice := my_slice[1:3:3]   // [22 33]

// 扩容
app_slice := append(new_slice,4444)
```

上面的`new_slice`的容量为2，并没有对应到底层数组的最结尾，所以`new_slice`是`my_slice`的一个真子集。这时对`new_slice`扩容，将生成一个新的底层数组，新的底层数组容量为4，而不是10。

实际上，当底层数组需要扩容时，会按照当前底层数组长度的2倍进行扩容，并生成新数组。如果底层数组的长度超过1000时，将按照25%的比率扩容，也就是1000个元素时，将扩展为1250个，不过这个增长比率的算法可能会随着go版本的递进而改变。

因为创建了新的底层数组，所以修改不同的slice，将不会互相影响。为了保证每次都是修改各自的底层数组，**通常会切出仅一个长度、仅一个容量的新slice，这样只要对它进行任何一次扩容，就会生成一个新的底层数组，从而让每个slice的底层数组都独立**。

```
my_slice := []int{11,22,33,44,55}

new_slice := my_slice[2:3:3]
app_slice := append(new_slice,3333)
```

事实上，这还是会出现共享的几率，因为没有扩容时，那个唯一的元素仍然是共享的，修改它肯定会影响至少1个slice，只有切出长度为0，容量为0的slice，才能完全保证独立性，但这和新创建一个slice没有区别。

总结：当slice的length已经等于capacity的时候，这时使用append()函数会新创建底层数组，从而实现隔离，如果新的slice是旧的slice的真子集，这时对其进行扩容会生成新的底层数组，capacity还是和新的slice一样，而不是翻倍。













如果函数接受slice参数，那么它对slice元素所做的更改将对调用者可见，类似于传递指向底层array的指针。

因此，Read函数可以接受slice实参，而不是指针和计数;slice内的长度设置了读取数据的上限。下面是包操作系统中File类型的Read方法的签名

```
func (f *File) Read(buf []byte) (n int, err error)
```

该方法返回读取的字节数和错误值(如果有的话)。要读入较大buffer buf的前32个字节，请slice  the buffer(这里用作动词)。

```
    n, err := f.Read(buf[0:32])
```

这样的slice是常见和有效的。实际上，先不考虑效率问题，下面的代码段还将读取缓冲区的前32个字节。

```
    var n int
    var err error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
```

### 