package main

import "fmt"

//defer与return的执行功能顺序 实际开发不要这样写
func func1() {
	fmt.Print("A")
}
func func2() {
	fmt.Print("B")
}
func func3() {
	fmt.Print("C")
}

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return call func...")
	return 1
}

func deferAndReturnFunc() int {
	defer deferFunc() //在return 执行前执行
	return returnFunc()
}

func main() {
	//压栈 return call func...
	//defer call func...
	//1
	//CBA
	defer func1()
	defer func2()
	defer func3()
	fmt.Println(deferAndReturnFunc())
}
