package main

import "fmt"

func fun0(a string) {
	fmt.Println("--fun0--")
	fmt.Println(a)
}

func fun1(a string, b int) int {
	fmt.Println("--fun1--")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

//返回多个值，匿名
func fun2(a string, b int) (int, int) {
	fmt.Println("--fun2--")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	d := 200
	return c, d
}

//return multi values
func fun3(a string, b int) (c int, d int) {
	fmt.Println("--fun3--")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("c=", c)
	fmt.Println("d=", d)

	c = 100
	d = 200
	return c, d
}

func fun4(string, int) {
	fmt.Println("--fun4--")
	fmt.Println("Fuck the fucking world!")
}

func fun5(a string, b int) (c, d int) {
	fmt.Println("--fun5--")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c = 100
	d = 200
	return c, d
}

func main() {

	fun0("Hello World!")

	a := fun1("Hello world!", 10)
	fmt.Println("a=", a)

	fmt.Println("------------------------")

	c, d := fun2("Hello world!", 10)
	fmt.Println("c=", c, "d=", d)

	fmt.Println("------------------------")

	e, f := fun3("Hello world!", 10)
	fmt.Println("e=", e, "f=", f)

	fmt.Println("------------------------")

	fun4("", 0)

	fmt.Println("------------------------")

	g, h := fun5("Hello world!", 10)
	fmt.Println("g=", g, "h=", h)

}
