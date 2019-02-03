package test

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	age  int
}

type admin struct {
	person user
	level  string
}

func (u user) SayHi() {
	fmt.Println(u.name, " say hi!")
}

func (u user) ChangeUser1() {
	u.name = "haha1"
}

func (u *user) ChangeUser2() {
	u.name = "haha2"
}

func TestStruct(test *testing.T) {
	var xiaoming user
	xiaoming.name = "xiaoming"
	xiaoming.age = 20
	fmt.Println(xiaoming)
	fmt.Println("xiaoming.name:", xiaoming.name)

	lili := user{
		name: "lili",
		age:  20,
	}
	fmt.Println("lili:", lili)

	a := admin{
		person: user{
			name: "ss",
			age:  20,
		},
		level: "1",
	}

	fmt.Println("a:", a)

	lili.SayHi()

	//不改变
	lili.ChangeUser1()
	fmt.Println("执行ChangeUser1后;", lili)
	//改变
	lili.ChangeUser2()
	fmt.Println("执行ChangeUser2后;", lili)
}
