package main

import "fmt"

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
}
