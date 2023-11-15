package main

/**
数组
*/
import "fmt"

/**
切片（动态数组 地址引用）
*/
func printArr(arr [4]int) {
	//不处理index
	for _, value := range arr {
		fmt.Println("value=", value)
	}
}

func printArr2(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	arr[0] = 100
}

func main() {
	//前六位值分别为1，2，3
	arr1 := [3]int{1, 2, 3}
	//固定长度10 每个数值为0
	var arr2 [1]int
	var arr3 [4]int

	printArr(arr3)
	fmt.Println("----------------------------------")

	//数组遍历方式
	for i := 0; i < len(arr2); i++ {
		fmt.Println("index=", i, " value=", arr2[i])
	}
	//数组遍历方式
	fmt.Println("---------------before-------------------")
	printArr2(arr1) //值copy
	fmt.Println("---------------after-------------------")
	//数组遍历方式
	for i := 0; i < len(arr1); i++ {
		fmt.Println("index=", i, " value=", arr1[i])
	}
}
