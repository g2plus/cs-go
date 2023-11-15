package main

import (
	_ "../init/lib1"   //匿名导入
	lib "../init/lib2" //给包起一个别名
	//. 导入 直接调用包内函数 不需要加包名.方法名Lx
)

/**
导入程序包时会执行包里面的init函数
*/
func main() {
	//lib1.Lib1Test()
	lib.Lib2Test()
}
