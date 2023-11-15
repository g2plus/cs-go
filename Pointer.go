package main

import "fmt"

/**
地址引用 & *
*/
func main() {
	var a, b = 10, 20
	fmt.Println("a=", a, "b=", b)
	swap(&a, &b)
	var pa *int
	pa = &a
	fmt.Println(&a)
	fmt.Println(pa)
	fmt.Println(pa == &a)
	fmt.Println("a=", a, "b=", b)
}

func swap(i *int, j *int) {
	var temp int
	temp = *i
	*i = *j
	*j = temp
}
