package main

import "fmt"

func main() {
	v := 42     // 修改这里！
	z := v / 78 //不是小数，而是整数
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("z is of type %T, Value: %v\n", z, z)
}

/* 类型推导

在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），
变量的类型由右值推导得出。

当右值声明了类型时，新变量的类型与其相同：

var i int
j := i // j 也是一个 int */
