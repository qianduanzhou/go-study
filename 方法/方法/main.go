package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

//结构体
type User struct {
	Name  string
	Email string
}

func (u User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}
func (u *User) Notify2() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

type Manager struct {
	User
	title string
}

func (m *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", m, m)
}
func (m Manager) ToString2() string {
	return fmt.Sprintf("Manager: %v", m)
}

type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
}
func (t *T) test2() {
	fmt.Println("类型 T 方法集包含全部 receiver *T 方法。")
}

type S struct {
	T
}

type Sx struct {
	*T
}

func (t T) testT() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
}
func (t *T) testT2() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 *T 方法。")
}

func main() {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()
	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
	u3.Notify2()

	//通过匿名字段，可获得和继承类似的复用能力。依据编译器查找次序，只需在外层定义同名方法，就可以实现 "override"。
	m := Manager{User{"Tom", "golang@golang.com"}, "Administrator"}
	m2 := &m
	fmt.Println(m.ToString())
	m.User.Notify()

	//方法集
	fmt.Println("-----")

	t1 := T{1}
	fmt.Printf("t1 is : %v\n", t1)
	t1.test()
	t1.test2()

	fmt.Println("-----")
	fmt.Println(m.ToString2())

	m.Notify()
	m.Notify2()
	m2.Notify()
	m2.Notify2()
	fmt.Println("-----")
	s := S{T{1}}
	s2 := &s
	s.testT()
	s.testT2()
	s2.testT()
	s2.testT2()
	fmt.Println("-----")
	sx := Sx{&T{1}}
	sx2 := &sx
	sx.testT()
	sx.testT2()
	sx2.testT()
	sx2.testT2()

	//表达式
	fmt.Println("-----")
	mValue := u1.Notify
	//method value方式会复制receiver
	u1.Name, u1.Email = "zhou", "@.com"
	u1.Notify()
	mValue()
	mValue2 := (User).Notify
	mValue2(u1)
	fmt.Println("-----")
	//系统抛出异常
	test04()
	fmt.Println("-----")
	area, err := getCircleArea2(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
	fmt.Println("-----")

	//自定义异常
	err2 := Open("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err2.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:

	}
}

// 系统抛
func test01() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	//a[10] = 11
	index := 10
	a[index] = 10
	fmt.Println(a)
}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		// 自己抛
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}

func test02() {
	getCircleArea(-5)
}

//
func test03() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

func test04() {
	test03()
	fmt.Println("test04")
}

//返回异常
func getCircleArea2(radius float32) (area float32, err error) {
	if radius < 0 {
		// 构建个异常对象
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 * radius * radius
	return
}

//自定义异常
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

//error类型需要有Error方法
func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}
