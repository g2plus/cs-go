package entity

import (
	"fmt"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

func (this *Student) Show() {
	fmt.Printf("%v", this)
}

func (this *Student) SetName(name string) {
	this.Name = name
}
func (this *Student) GetName() string {
	return this.Name
}
func (this *Student) SetAge(age int) {
	this.Age = age
}

func (this *Student) GetAge() int {
	return this.Age
}

func (this *Student) SetGender(gender string) {
	this.Gender = gender
}
func (this *Student) GetGender() string {
	return this.Gender
}
