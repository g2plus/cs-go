package lib1

import "fmt"

/**
函数名的第一个字母大写表示可以被调用
*/
func Lib1Test() {
	fmt.Println("Hello world!")
}

func init() {
	fmt.Println("lib1 init")
}
