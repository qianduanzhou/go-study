package main

import "fmt"

type Sayer interface {
	say()
}

type dog struct{}
type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

/*值接收者实现接口*/
type Mover interface {
	move()
}

func (d dog) move() {
	fmt.Println("狗会动")
}

/*引用接收者实现接口*/
type Mover2 interface {
	move2()
}

func (d *dog) move2() {
	fmt.Println("狗会动")
}

/*一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。*/
// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
	fmt.Println("------")
	/*值接收者实现接口*/
	var mover1 Mover = dog{}
	var mover2 Mover = &dog{}
	mover1.move()
	mover2.move()
	/*指针接收者实现接口*/
	// var mover3 Mover2 = dog{}
	var mover4 Mover2 = &dog{}
	// mover3.move2()
	mover4.move2()

	/*空接口*/
	fmt.Println("------")
	var emptyI interface{}
	emptyI = "123123"
	fmt.Println(emptyI)
	emptyI = 4
	fmt.Println(emptyI)
	emptyI = true
	fmt.Println(emptyI)

	var m = map[string]interface{}{
		"name": "zhou",
		"age":  10,
	}
	fmt.Println(m)
	fmt.Println("------")

	/*类型断言*/
	var i interface{} = "pprof.cn"
	v, ok := i.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
