package main

import (
	"encoding/json"
	"fmt"
)

/*类型定义
 */
type NewInt int

/*类型别名
 */
type MyInt = int

/*结构体
 */
type person struct {
	name string
	city string
	age  int8
}

//### 结构体的匿名字段
type Test struct {
	string
	int
}

/*结构体嵌套
 */
//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

/*嵌套匿名结构体
 */
type User2 struct {
	Name    string
	Gender  string
	Address //匿名结构体
}

/*结构体继承
 */
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

/*结构体与JSON序列化、公私有，标签
 */
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func main() {
	//类型定义和类型别名的区别
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	//结构体实例化
	var p person
	p.name = "test"
	p.city = "Gz"
	p.age = 18
	fmt.Println("p", p)
	fmt.Printf("p=%#v\n", p)

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	//结构体指针
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}

	//取结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}

	//### 使用键值对初始化
	p5 := person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}
	//也可以对结构体指针进行键值对初始化，例如：
	p6 := &person{
		name: "pprof.cn",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"pprof.cn", city:"北京", age:18}

	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

	//使用值的列表初始化
	p8 := &person{
		"pprof.cn",
		"北京",
		18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}

	//调用构造函数
	p9 := newPerson("pprof.cn", "北京", 18)
	fmt.Printf("p9=%#v\n", p9) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}

	//调用方法
	p9.dream()

	//指针类型接受者
	p9.setAge()
	fmt.Printf("p9=%#v\n", p9)

	//值类型接受者
	p9.setCity()
	fmt.Printf("p9=%#v\n", p9)

	a.SayHello()

	//结构体的匿名字段
	var t = Test{
		"test",
		18,
	}
	fmt.Printf("t=%#v\n", t)

	//结构体嵌套
	user1 := User{
		Name:   "zhou",
		Gender: "1",
		Address: Address{
			Province: "gd",
			City:     "gz",
		},
	}
	fmt.Printf("user1=%#v\n", user1)

	//嵌套匿名结构体
	var user2 User2
	user2.Name = "zhou"
	user2.Gender = "Gender"
	user2.Address.Province = "gd"
	user2.City = "gz" //通过匿名结构体访问（如果有相同字段则必须通过上面的方式）
	fmt.Printf("user2=%#v\n", user2)

	//结构体继承
	dog := Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	dog.wang()
	dog.move()

	//结构体与JSON序列化
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	fmt.Printf("c:%#v\n", c)
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("c1:%#v\n", c1)
}

//模拟构造函数
func newPerson(name string, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//定义方法接受者
func (p person) dream() {
	fmt.Println("dream")
}

//指针类型的接受者
func (p *person) setAge() {
	p.age = 20
}

//值类型的接受者
func (p person) setCity() {
	p.city = "广州"
}

//任何类型都可以添加方法
func (nI NewInt) SayHello() {
	fmt.Println("hello")
}
