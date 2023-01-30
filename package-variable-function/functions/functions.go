package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(5, 89))
}

/*函数可以没有参数或接受多个参数。

在本例中，add 接受两个 int 类型的参数。

注意类型在变量名 之后。

当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。*/
