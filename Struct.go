package main

import "fmt"

//定义一个结构体
type T struct {
	name string
}

/**
值传递
*/
func (t T) method1() {
	t.name = "new name1"
}

/**
地址引用
*/
func (t *T) method2() {
	t.name = "new name2"
}

type Movie struct {
	title   string
	country string
}

func main() {

	t := T{"old name"}

	fmt.Println("method1 调用前 ", t.name)
	t.method1()
	fmt.Println("method1 调用后 ", t.name)

	fmt.Println("method2 调用前 ", t.name)
	t.method2()
	fmt.Println("method2 调用后 ", t.name)

	var movie Movie
	movie.title = "不能错过的只有你"
	movie.country = "中国"

	fmt.Printf("%v\n", movie)

	movie2 := Movie{
		"中国机长",
		"中国",
	}
	fmt.Println(movie2)
}
