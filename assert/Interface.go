package main

import "fmt"

/*
多态
func funcName(a assert{}) string {
        return string(a)
}
*/

func funcName(a interface{}) string {
	value, ok := a.(string)
	if !ok {
		fmt.Println("It is not ok for type string")
		return "error"
	}
	fmt.Println("The value is ", value)

	return value
}

func main() {
	//      str := "123"
	//      funcName(str)
	//var a assert{}
	//var a string = "123"
	var a string = "132"
	str := funcName(a)
	fmt.Println(str)
}
