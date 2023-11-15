package main

import "fmt"

var gA int = 1000
var gB = 1000

//const常量(readonly)
//iota标识符实现一组自增常量值来实现枚举类型
const (
	a1 = 10*iota + 1 // 0 iota=0 a1=10*0+1=2
	a2               // 1 iota=1 a2=10*1+1=11
	a3               // 2 iota=2 a3=10*2+1=21
)

//:= 只能在函数体内使用 不可以定义全局变量
//gc:=100

func main() {
	//默认值为0
	var a int = 10
	fmt.Println(a)

	//自动匹配数据类型
	var e = 10
	fmt.Println(e)
	fmt.Printf("type of e=%T\n", e)

	//常用定义变量方式 省略var
	f := "Hello word!"
	fmt.Printf("f=%s,type of f=%T\n", f, f)

	fmt.Println("gA=", gA, "gB=", gB)
	//fmt.Println("gC=", gC)

	//定义多个变量
	var x, y = 1, 2
	fmt.Println(x, y)

	//多行变量
	const (
		h uint16 = 120
		b        // 与上一行x类型、右值相同
		c = "123"
		z // 与上一行s类型、右值相同
	)

	fmt.Printf("%T, %v\n", h, h)
	fmt.Printf("%T, %v\n", z, z)

	var (
		j, n int
		l, q = 100, "abc"
	)

	fmt.Println(j, n, l, q)

	fmt.Println(a1, a2, a3)

}
