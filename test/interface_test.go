package test

import (
	"fmt"
	"testing"
)

type sleep interface {
	sleep()
	modifyName(n string)
}

type person struct {
	name string
}

type richPerson struct {
	person
	money int64
}

func (p *person) sleep() {
	fmt.Printf("%s开始睡觉...\n", p.name)
}

func (r *richPerson) sleep() {
	fmt.Printf("%s开始在自己的大房子里睡觉...\n", r.name)
}

//func (p person) modifyName(n string)  {
//	p.name=n
//}

func (p *person) modifyName(n string) {
	p.name = n
}
func TestInterface(test *testing.T) {
	p := person{"小明"}
	p.sleep()
	goToSleep(&p)
	goToSleep(&p)
	p.modifyName("丽丽")
	fmt.Println(p)
	p.modifyName("李磊")
	fmt.Println(p)
	modifyName(&p, "丹尼尔")
	fmt.Println(p)

	rp := richPerson{
		person: person{name: "富有的小明"},
		money:  10000000000,
	}

	rp.person.sleep()
	rp.sleep()
}

func goToSleep(s sleep) {
	s.sleep()
}

func modifyName(s sleep, n string) {
	s.modifyName(n)
}
