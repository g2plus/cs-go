package main

import "fmt"

/**
接口本质时指针
子类需要实现接口的所有方法
*/
type AnimalIF interface {
	Sleep()
	Eat()
	Print()
}

type Cat struct {
	Color string
}

func (this *Cat) SetColor(color string) {
	this.Color = color
}

func (this *Cat) Sleep() {
	fmt.Print("Cat is sleeping")
}

func (this *Cat) Eat() {
	fmt.Print("Cat is eating")
}
func (this *Cat) Print() {
	fmt.Printf("Cat's color %s\n", this.Color)
}

type Dog struct {
	Color string
}

func (this *Dog) SetColor(color string) {
	this.Color = color
}

func (this *Dog) Sleep() {
	fmt.Print("Cat is sleeping")
}

func (this *Dog) Eat() {
	fmt.Print("Cat is eating")
}
func (this *Dog) Print() {
	fmt.Printf("Dog's color %s\n", this.Color)
}

func show(animal AnimalIF) {
	animal.Print()
}

func main() {
	cat := Cat{"White"}
	show(&cat)
	dog := Dog{"Yellow"}
	show(&dog)

	var animal AnimalIF
	animal = &Cat{"White"}
	animal.Print()

	animal = &Dog{"Black"}
	animal.Print()
}
