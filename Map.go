package main

import (
	"fmt"
)

/**
地址引用
*/
func main() {
	//第一种声明
	funcName1()

	var test0 map[int]string
	if test0 == nil {
		fmt.Println("testo is nil")
	}

	var test1 = make(map[string]string, 10)
	test1["hello"] = "world"
	fmt.Println(test1)
	//第二种声明
	funcName2()
	test2 := make(map[string]string, 10)
	test2["hello"] = "world"
	fmt.Println(test2)
	for key, value := range test2 {
		fmt.Println("key=", key, " value=", value)
	}
	fmt.Println("***************************")
	for _, value := range test2 {
		fmt.Println("value=", value)
	}
	//第三种声明
	funcName3()
	test3 := map[string]string{
		"hello": "world",
	}
	fmt.Println(test3)

	language := funcName4()
	test4 := make(map[string]map[string]string, 10)
	//需要一级一级展开定义
	test4["springboot"] = make(map[string]string, 2)
	test4["springboot"]["rocketmq"] = "rocketmq"
	test4["springboot"]["eventbus"] = "eventbus"
	fmt.Println(test4["springboot"])

	test5 := make(map[string]map[string]map[string]string, 10)
	test5["java"] = make(map[string]map[string]string, 2)
	test5["java"]["hello"] = make(map[string]string, 3)
	test5["java"]["hello"]["test1"] = "test1"
	test5["java"]["hello"]["test2"] = "test2"
	fmt.Println(test5)

	//需要一级一级展开定义
	test4["springboot"] = make(map[string]string, 2)
	test4["springboot"]["rocketmq"] = "rocketmq"
	test4["springboot"]["eventbus"] = "eventbus"
	fmt.Println(test4["springboot"])

	//增删改查
	val, key := language["php"] //查找是否有php这个子元素
	if key {
		fmt.Printf("%v", val)
	} else {
		fmt.Printf("no")
	}

	//地址引用
	language["php"]["id"] = "3"         //修改了php子元素的id值
	language["php"]["nickname"] = "啪啪啪" //增加php元素里的nickname值
	delete(language, "php")             //删除了php子元素
	fmt.Println("________deleted___________")
	fmt.Println(language)
}

func funcName4() map[string]map[string]string {
	language := make(map[string]map[string]string)
	language["php"] = make(map[string]string, 2)
	language["php"]["id"] = "1"
	language["php"]["desc"] = "php是世界上最美的语言"
	language["golang"] = make(map[string]string, 2)
	language["golang"]["id"] = "2"
	language["golang"]["desc"] = "golang抗并发非常good"

	fmt.Println(language) //map[php:map[id:1 desc:php是世界上最美的语言] golang:map[id:2 desc:golang抗并发非常good]]
	return language
}

func funcName3() {
	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println(test3) //map[one:php two:golang three:java]
}

func funcName2() {
	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "golang"
	test2["three"] = "java"
	fmt.Println(test2) //map[one:php two:golang three:java]
}

func funcName1() {
	var test1 map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "golang"
	test1["three"] = "java"
	fmt.Println(test1) //map[two:golang three:java one:php]
	fmt.Println(len(test1))
}
