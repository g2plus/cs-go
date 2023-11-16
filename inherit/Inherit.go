package main

import "fmt"

/**
继承，在子类中是使用父类 子类无该方法调用父类的该方法，子类有调用子类该方法
*/
type Human struct {
	Name   string
	Gender string
}

func (this *Human) Eat() {
	fmt.Print("Human eat...")
}

func (this *Human) SetName(name string) {
	this.Name = name
}

func (this *Human) SetGender(gender string) {
	this.Gender = gender
}

func (this *Human) Print() {
	fmt.Printf("%v\n", this)
}

func (this *Human) Sleep() {
	fmt.Println("Human sleep")
}

type SuperMan struct {
	Human Human
	Hobby string
}

func (this *SuperMan) SetGender(gender string) {
	this.Human.Gender = gender
}

func (this *SuperMan) SetName(name string) {
	this.Human.Name = name
}

func (this *SuperMan) Eat() {
	fmt.Println("Superman eat...")
}

func (this *SuperMan) Sleep() {
	fmt.Println("never sleep...")
}

func (this *SuperMan) Print() {
	fmt.Printf("%v\n", this)
}

func (this *SuperMan) Fly() {
	fmt.Println("superman can fly")
}

func main() {

	superman := SuperMan{Human{"zhangsan", "男"}, "英雄联盟"}
	superman.Print()
	superman.Eat()
	superman.Sleep()
	superman.Fly()

	var man SuperMan
	man.Human.Name = "lisi"
	man.Hobby = "英雄联盟"
	man.Human.Gender = "女"
	man.Print()
	man.Eat()
	man.Fly()

	man.SetName("李木子")
	man.SetGender("女")
	man.Print()

}
