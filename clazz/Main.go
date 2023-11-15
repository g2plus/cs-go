package main

import "../clazz/entity"

func main() {
	t := entity.Student{Name: "张三", Gender: "男", Age: 26}
	t.Show()
	t.SetAge(18)
	t.SetName("李木子")
	t.SetGender("女")
	t.Show()
}
