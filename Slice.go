package main

import "fmt"

/**
地址引用
*/
func printSlice(arr []int) {
	for index, value := range arr {
		fmt.Println("index=", index, " value=", value)
	}
	arr[0] = 100
}

func defineSlice() {
	slice1 := []int{1, 2, 3}
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(slice1), cap(slice1), slice1)
	var slice2 []int = make([]int, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(slice2), cap(slice2), slice2)
	slice3 := make([]int, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(slice3), cap(slice3), slice3)
	var slice4 []int
	if slice4 == nil {
		fmt.Println("slice4 is empty...")
	} else {
		fmt.Println("slice4 is not empty...")
	}
}

func appendAndSplitSlice() {

}

/**
浅拷贝与数据截取
*/
func mycopy() {
	slice1 := []int{1, 2, 3}
	slice2 := slice1[0:2]
	slice2[0] = 1024
	fmt.Println(slice1)
	fmt.Println(slice2)
}

/**
深度拷贝
*/
func deepCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 3)
	//深度copy重新建立一份额外数据
	copy(slice2, slice1)
	slice1[0] = 1024
	fmt.Println(slice1)
	fmt.Println(slice2)
}

/**
扩容机制
*/
func myappend() {
	var numbers = make([]int, 10, 10)
	fmt.Println("------扩容前-------")
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 10)
	fmt.Println("------扩容后-------")
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
}

func main() {
	arr := []int{1, 2, 4, 5}
	fmt.Println("---------------before-------------------")
	printSlice(arr) //值copy
	fmt.Println("---------------after-------------------")
	for index, value := range arr {
		fmt.Println("index=", index, " value=", value)
	}

	fmt.Println("------------")
	defineSlice()
	fmt.Println("-----------------")
	mycopy()
	fmt.Println("-----------------")
	deepCopy()
	fmt.Println("-----------------")
	myappend()
}
